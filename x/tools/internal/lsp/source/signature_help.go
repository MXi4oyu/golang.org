// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

	"golang.org/x/tools/go/ast/astutil"
)

type SignatureInformation struct {
	Label           string
	Parameters      []ParameterInformation
	ActiveParameter int
}

type ParameterInformation struct {
	Label string
}

func SignatureHelp(ctx context.Context, f File, pos token.Pos) (*SignatureInformation, error) {
<<<<<<< HEAD
	fAST, err := f.GetAST()
	if err != nil {
		return nil, err
	}
	pkg, err := f.GetPackage()
	if err != nil {
		return nil, err
=======
	fAST := f.GetAST(ctx)
	pkg := f.GetPackage(ctx)
	if pkg.IsIllTyped() {
		return nil, fmt.Errorf("package for %s is ill typed", f.URI())
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	}

	// Find a call expression surrounding the query position.
	var callExpr *ast.CallExpr
	path, _ := astutil.PathEnclosingInterval(fAST, pos, pos)
	if path == nil {
		return nil, fmt.Errorf("cannot find node enclosing position")
	}
	for _, node := range path {
<<<<<<< HEAD
		if c, ok := node.(*ast.CallExpr); ok {
=======
		if c, ok := node.(*ast.CallExpr); ok && pos >= c.Lparen && pos <= c.Rparen {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
			callExpr = c
			break
		}
	}
	if callExpr == nil || callExpr.Fun == nil {
		return nil, fmt.Errorf("cannot find an enclosing function")
	}

<<<<<<< HEAD
	// Get the type information for the function corresponding to the call expression.
	var obj types.Object
	switch t := callExpr.Fun.(type) {
	case *ast.Ident:
		obj = pkg.TypesInfo.ObjectOf(t)
	case *ast.SelectorExpr:
		obj = pkg.TypesInfo.ObjectOf(t.Sel)
	default:
		return nil, fmt.Errorf("the enclosing function is malformed")
	}
	if obj == nil {
		return nil, fmt.Errorf("cannot resolve %s", callExpr.Fun)
	}
	// Find the signature corresponding to the object.
	var sig *types.Signature
	switch obj.(type) {
	case *types.Var:
		if underlying, ok := obj.Type().Underlying().(*types.Signature); ok {
			sig = underlying
		}
	case *types.Func:
		sig = obj.Type().(*types.Signature)
	}
	if sig == nil {
		return nil, fmt.Errorf("no function signatures found for %s", obj.Name())
	}
	pkgStringer := qualifier(fAST, pkg.Types, pkg.TypesInfo)
	var paramInfo []ParameterInformation
	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		label := types.TypeString(param.Type(), pkgStringer)
=======
	// Get the type information for the function being called.
	sigType := pkg.GetTypesInfo().TypeOf(callExpr.Fun)
	if sigType == nil {
		return nil, fmt.Errorf("cannot get type for Fun %[1]T (%[1]v)", callExpr.Fun)
	}

	sig, _ := sigType.Underlying().(*types.Signature)
	if sig == nil {
		return nil, fmt.Errorf("cannot find signature for Fun %[1]T (%[1]v)", callExpr.Fun)
	}

	qf := qualifier(fAST, pkg.GetTypes(), pkg.GetTypesInfo())
	var paramInfo []ParameterInformation
	for i := 0; i < sig.Params().Len(); i++ {
		param := sig.Params().At(i)
		label := types.TypeString(param.Type(), qf)
		if sig.Variadic() && i == sig.Params().Len()-1 {
			label = strings.Replace(label, "[]", "...", 1)
		}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		if param.Name() != "" {
			label = fmt.Sprintf("%s %s", param.Name(), label)
		}
		paramInfo = append(paramInfo, ParameterInformation{
			Label: label,
		})
	}
<<<<<<< HEAD
	// Determine the query position relative to the number of parameters in the function.
	var activeParam int
	var start, end token.Pos
	for i, expr := range callExpr.Args {
=======

	// Determine the query position relative to the number of parameters in the function.
	var activeParam int
	var start, end token.Pos
	for _, expr := range callExpr.Args {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		if start == token.NoPos {
			start = expr.Pos()
		}
		end = expr.End()
<<<<<<< HEAD
		if i < len(callExpr.Args)-1 {
			end = callExpr.Args[i+1].Pos() - 1 // comma
		}
		if start <= pos && pos <= end {
			break
		}
		activeParam++
		start = expr.Pos() + 1 // to account for commas
	}
	// Label for function, qualified by package name.
	label := obj.Name()
	if pkg := pkgStringer(obj.Pkg()); pkg != "" {
		label = pkg + "." + label
	}
	return &SignatureInformation{
		Label:           label + formatParams(sig.Params(), sig.Variadic(), pkgStringer),
=======
		if start <= pos && pos <= end {
			break
		}

		// Don't advance the active parameter for the last parameter of a variadic function.
		if !sig.Variadic() || activeParam < sig.Params().Len()-1 {
			activeParam++
		}
		start = expr.Pos() + 1 // to account for commas
	}

	// Get the object representing the function, if available.
	// There is no object in certain cases such as calling a function returned by
	// a function (e.g. "foo()()").
	var obj types.Object
	switch t := callExpr.Fun.(type) {
	case *ast.Ident:
		obj = pkg.GetTypesInfo().ObjectOf(t)
	case *ast.SelectorExpr:
		obj = pkg.GetTypesInfo().ObjectOf(t.Sel)
	}

	var name string
	if obj != nil {
		name = obj.Name()
	} else {
		name = "func"
	}

	results, writeResultParens := formatResults(sig.Results(), qf)
	label, detail := formatFunction(name, formatParams(sig.Params(), sig.Variadic(), qf), results, writeResultParens)
	if sig.Results().Len() > 0 {
		label += " " + detail
	}
	return &SignatureInformation{
		Label:           label,
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		Parameters:      paramInfo,
		ActiveParameter: activeParam,
	}, nil
}
