// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
<<<<<<< HEAD
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/internal/lsp/source"
=======
	"context"
	"go/ast"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
)

// File holds all the information we know about a file.
type File struct {
<<<<<<< HEAD
	URI     source.URI
=======
	uris     []span.URI
	filename string
	basename string

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	view    *View
	active  bool
	content []byte
	ast     *ast.File
	token   *token.File
<<<<<<< HEAD
	pkg     *packages.Package
}

// SetContent sets the overlay contents for a file.
// Setting it to nil will revert it to the on disk contents, and remove it
// from the active set.
func (f *File) SetContent(content []byte) {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	f.content = content
	// the ast and token fields are invalid
	f.ast = nil
	f.token = nil
	f.pkg = nil
	// and we might need to update the overlay
	switch {
	case f.active && content == nil:
		// we were active, and want to forget the content
		f.active = false
		if filename, err := f.URI.Filename(); err == nil {
			delete(f.view.Config.Overlay, filename)
		}
		f.content = nil
	case content != nil:
		// an active overlay, update the map
		f.active = true
		if filename, err := f.URI.Filename(); err == nil {
			f.view.Config.Overlay[filename] = f.content
		}
	}
}

// Read returns the contents of the file, reading it from file system if needed.
func (f *File) Read() ([]byte, error) {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	return f.read()
}

func (f *File) GetFileSet() (*token.FileSet, error) {
	if f.view.Config == nil {
		return nil, fmt.Errorf("no config for file view")
	}
	if f.view.Config.Fset == nil {
		return nil, fmt.Errorf("no fileset for file view config")
	}
	return f.view.Config.Fset, nil
}

func (f *File) GetToken() (*token.File, error) {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	if f.token == nil {
		if err := f.view.parse(f.URI); err != nil {
			return nil, err
		}
		if f.token == nil {
			return nil, fmt.Errorf("failed to find or parse %v", f.URI)
		}
	}
	return f.token, nil
}

func (f *File) GetAST() (*ast.File, error) {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	if f.ast == nil {
		if err := f.view.parse(f.URI); err != nil {
			return nil, err
		}
	}
	return f.ast, nil
}

func (f *File) GetPackage() (*packages.Package, error) {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	if f.pkg == nil {
		if err := f.view.parse(f.URI); err != nil {
			return nil, err
		}
	}
	return f.pkg, nil
}

// read is the internal part of Read that presumes the lock is already held
func (f *File) read() ([]byte, error) {
	if f.content != nil {
		return f.content, nil
	}
	// we don't know the content yet, so read it
	filename, err := f.URI.Filename()
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	f.content = content
	return f.content, nil
=======
	pkg     *Package
	meta    *metadata
	imports []*ast.ImportSpec
}

func basename(filename string) string {
	return strings.ToLower(filepath.Base(filename))
}

func (f *File) URI() span.URI {
	return f.uris[0]
}

// View returns the view associated with the file.
func (f *File) View() source.View {
	return f.view
}

// GetContent returns the contents of the file, reading it from file system if needed.
func (f *File) GetContent(ctx context.Context) []byte {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()

	if ctx.Err() == nil {
		f.read(ctx)
	}

	return f.content
}

func (f *File) GetFileSet(ctx context.Context) *token.FileSet {
	return f.view.Config.Fset
}

func (f *File) GetToken(ctx context.Context) *token.File {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()

	if f.token == nil || len(f.view.contentChanges) > 0 {
		if _, err := f.view.parse(ctx, f); err != nil {
			return nil
		}
	}
	return f.token
}

func (f *File) GetAST(ctx context.Context) *ast.File {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()

	if f.ast == nil || len(f.view.contentChanges) > 0 {
		if _, err := f.view.parse(ctx, f); err != nil {
			return nil
		}
	}
	return f.ast
}

func (f *File) GetPackage(ctx context.Context) source.Package {
	f.view.mu.Lock()
	defer f.view.mu.Unlock()
	if f.pkg == nil || len(f.view.contentChanges) > 0 {
		if errs, err := f.view.parse(ctx, f); err != nil {
			// Create diagnostics for errors if we are able to.
			if len(errs) > 0 {
				return &Package{errors: errs}
			}
			return nil
		}
	}
	return f.pkg
}

// read is the internal part of GetContent. It assumes that the caller is
// holding the mutex of the file's view.
func (f *File) read(ctx context.Context) {
	if f.content != nil {
		if len(f.view.contentChanges) == 0 {
			return
		}

		f.view.mcache.mu.Lock()
		err := f.view.applyContentChanges(ctx)
		f.view.mcache.mu.Unlock()

		if err == nil {
			return
		}
	}
	// We might have the content saved in an overlay.
	if content, ok := f.view.Config.Overlay[f.filename]; ok {
		f.content = content
		return
	}
	// We don't know the content yet, so read it.
	content, err := ioutil.ReadFile(f.filename)
	if err != nil {
		f.view.Logger().Errorf(ctx, "unable to read file %s: %v", f.filename, err)
		return
	}
	f.content = content
}

// isPopulated returns true if all of the computed fields of the file are set.
func (f *File) isPopulated() bool {
	return f.ast != nil && f.token != nil && f.pkg != nil && f.meta != nil && f.imports != nil
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
