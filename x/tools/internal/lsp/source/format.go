// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package source provides core features for use by Go editors and tools.
package source

import (
	"bytes"
	"context"
	"fmt"
<<<<<<< HEAD
	"go/ast"
	"go/format"

	"golang.org/x/tools/go/ast/astutil"
)

// Format formats a document with a given range.
func Format(ctx context.Context, f File, rng Range) ([]TextEdit, error) {
	fAST, err := f.GetAST()
	if err != nil {
		return nil, err
	}
=======
	"go/format"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/imports"
	"golang.org/x/tools/internal/lsp/diff"
	"golang.org/x/tools/internal/span"
)

// Format formats a file with a given range.
func Format(ctx context.Context, f File, rng span.Range) ([]TextEdit, error) {
	pkg := f.GetPackage(ctx)
	if hasParseErrors(pkg.GetErrors()) {
		return nil, fmt.Errorf("%s has parse errors, not formatting", f.URI())
	}
	fAST := f.GetAST(ctx)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	path, exact := astutil.PathEnclosingInterval(fAST, rng.Start, rng.End)
	if !exact || len(path) == 0 {
		return nil, fmt.Errorf("no exact AST node matching the specified range")
	}
	node := path[0]
<<<<<<< HEAD
	// format.Node can fail when the AST contains a bad expression or
	// statement. For now, we preemptively check for one.
	// TODO(rstambler): This should really return an error from format.Node.
	var isBad bool
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.BadDecl, *ast.BadExpr, *ast.BadStmt:
			isBad = true
			return false
		default:
			return true
		}
	})
	if isBad {
		return nil, fmt.Errorf("unable to format file due to a badly formatted AST")
	}
=======
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	// format.Node changes slightly from one release to another, so the version
	// of Go used to build the LSP server will determine how it formats code.
	// This should be acceptable for all users, who likely be prompted to rebuild
	// the LSP server on each Go release.
<<<<<<< HEAD
	fset, err := f.GetFileSet()
	if err != nil {
		return nil, err
	}
=======
	fset := f.GetFileSet(ctx)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	buf := &bytes.Buffer{}
	if err := format.Node(buf, fset, node); err != nil {
		return nil, err
	}
<<<<<<< HEAD
	// TODO(rstambler): Compute text edits instead of replacing whole file.
	return []TextEdit{
		{
			Range:   rng,
			NewText: buf.String(),
		},
	}, nil
=======
	return computeTextEdits(ctx, f, buf.String()), nil
}

func hasParseErrors(errors []packages.Error) bool {
	for _, err := range errors {
		if err.Kind == packages.ParseError {
			return true
		}
	}
	return false
}

// Imports formats a file using the goimports tool.
func Imports(ctx context.Context, f File, rng span.Range) ([]TextEdit, error) {
	formatted, err := imports.Process(f.GetToken(ctx).Name(), f.GetContent(ctx), nil)
	if err != nil {
		return nil, err
	}
	return computeTextEdits(ctx, f, string(formatted)), nil
}

func computeTextEdits(ctx context.Context, file File, formatted string) (edits []TextEdit) {
	u := diff.SplitLines(string(file.GetContent(ctx)))
	f := diff.SplitLines(formatted)
	return DiffToEdits(file.URI(), diff.Operations(u, f))
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
