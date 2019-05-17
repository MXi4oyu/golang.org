// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gopackages command is a diagnostic tool that demonstrates
// how to use golang.org/x/tools/go/packages to load, parse,
// type-check, and print one or more Go packages.
// Its precise output is unspecified and may change.
package main

import (
<<<<<<< HEAD
=======
	"context"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	"encoding/json"
	"flag"
	"fmt"
	"go/types"
<<<<<<< HEAD
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
=======
	"os"
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"
<<<<<<< HEAD
)

// flags
var (
	depsFlag  = flag.Bool("deps", false, "show dependencies too")
	testFlag  = flag.Bool("test", false, "include any tests implied by the patterns")
	mode      = flag.String("mode", "imports", "mode (one of files, imports, types, syntax, allsyntax)")
	private   = flag.Bool("private", false, "show non-exported declarations too")
	printJSON = flag.Bool("json", false, "print package in JSON form")

	cpuprofile = flag.String("cpuprofile", "", "write CPU profile to this file")
	memprofile = flag.String("memprofile", "", "write memory profile to this file")
	traceFlag  = flag.String("trace", "", "write trace log to this file")

	buildFlags stringListValue
)

func init() {
	flag.Var(&buildFlags, "buildflag", "pass argument to underlying build system (may be repeated)")
}

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: gopackages [-deps] [-cgo] [-mode=...] [-private] package...

The gopackages command loads, parses, type-checks,
and prints one or more Go packages.

Packages are specified using the notation of "go list",
or other underlying build system.

Flags:`)
	flag.PrintDefaults()
}

func main() {
	log.SetPrefix("gopackages: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		usage()
		os.Exit(1)
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		// NB: profile won't be written in case of error.
		defer pprof.StopCPUProfile()
	}

	if *traceFlag != "" {
		f, err := os.Create(*traceFlag)
		if err != nil {
			log.Fatal(err)
		}
		if err := trace.Start(f); err != nil {
			log.Fatal(err)
		}
		// NB: trace log won't be written in case of error.
		defer func() {
			trace.Stop()
			log.Printf("To view the trace, run:\n$ go tool trace view %s", *traceFlag)
		}()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		// NB: memprofile won't be written in case of error.
		defer func() {
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatalf("Writing memory profile: %v", err)
			}
			f.Close()
		}()
=======
	"golang.org/x/tools/internal/tool"
)

func main() {
	tool.Main(context.Background(), &application{Mode: "imports"}, os.Args[1:])
}

type application struct {
	// Embed the basic profiling flags supported by the tool package
	tool.Profile

	Deps       bool            `flag:"deps" help:"show dependencies too"`
	Test       bool            `flag:"test" help:"include any tests implied by the patterns"`
	Mode       string          `flag:"mode" help:"mode (one of files, imports, types, syntax, allsyntax)"`
	Private    bool            `flag:"private" help:"show non-exported declarations too"`
	PrintJSON  bool            `flag:"json" help:"print package in JSON form"`
	BuildFlags stringListValue `flag:"buildflag" help:"pass argument to underlying build system (may be repeated)"`
}

// Name implements tool.Application returning the binary name.
func (app *application) Name() string { return "gopackages" }

// Usage implements tool.Application returning empty extra argument usage.
func (app *application) Usage() string { return "package..." }

// ShortHelp implements tool.Application returning the main binary help.
func (app *application) ShortHelp() string {
	return "gopackages loads, parses, type-checks, and prints one or more Go packages."
}

// DetailedHelp implements tool.Application returning the main binary help.
func (app *application) DetailedHelp(f *flag.FlagSet) {
	fmt.Fprint(f.Output(), `
Packages are specified using the notation of "go list",
or other underlying build system.

Flags:
`)
	f.PrintDefaults()
}

// Run takes the args after flag processing and performs the specified query.
func (app *application) Run(ctx context.Context, args ...string) error {
	if len(args) == 0 {
		return tool.CommandLineErrorf("not enough arguments")
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	}

	// Load, parse, and type-check the packages named on the command line.
	cfg := &packages.Config{
		Mode:       packages.LoadSyntax,
<<<<<<< HEAD
		Tests:      *testFlag,
		BuildFlags: buildFlags,
	}

	// -mode flag
	switch strings.ToLower(*mode) {
=======
		Tests:      app.Test,
		BuildFlags: app.BuildFlags,
	}

	// -mode flag
	switch strings.ToLower(app.Mode) {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
	case "files":
		cfg.Mode = packages.LoadFiles
	case "imports":
		cfg.Mode = packages.LoadImports
	case "types":
		cfg.Mode = packages.LoadTypes
	case "syntax":
		cfg.Mode = packages.LoadSyntax
	case "allsyntax":
		cfg.Mode = packages.LoadAllSyntax
	default:
<<<<<<< HEAD
		log.Fatalf("invalid mode: %s", *mode)
	}

	lpkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	// -deps: print dependencies too.
	if *depsFlag {
=======
		return tool.CommandLineErrorf("invalid mode: %s", app.Mode)
	}

	lpkgs, err := packages.Load(cfg, args...)
	if err != nil {
		return err
	}

	// -deps: print dependencies too.
	if app.Deps {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		// We can't use packages.All because
		// we need an ordered traversal.
		var all []*packages.Package // postorder
		seen := make(map[*packages.Package]bool)
		var visit func(*packages.Package)
		visit = func(lpkg *packages.Package) {
			if !seen[lpkg] {
				seen[lpkg] = true

				// visit imports
				var importPaths []string
				for path := range lpkg.Imports {
					importPaths = append(importPaths, path)
				}
				sort.Strings(importPaths) // for determinism
				for _, path := range importPaths {
					visit(lpkg.Imports[path])
				}

				all = append(all, lpkg)
			}
		}
		for _, lpkg := range lpkgs {
			visit(lpkg)
		}
		lpkgs = all
	}

	for _, lpkg := range lpkgs {
<<<<<<< HEAD
		print(lpkg)
	}
}

func print(lpkg *packages.Package) {
	if *printJSON {
=======
		app.print(lpkg)
	}
	return nil
}

func (app *application) print(lpkg *packages.Package) {
	if app.PrintJSON {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		data, _ := json.MarshalIndent(lpkg, "", "\t")
		os.Stdout.Write(data)
		return
	}
	// title
	var kind string
	// TODO(matloob): If IsTest is added back print "test command" or
	// "test package" for packages with IsTest == true.
	if lpkg.Name == "main" {
		kind += "command"
	} else {
		kind += "package"
	}
	fmt.Printf("Go %s %q:\n", kind, lpkg.ID) // unique ID
	fmt.Printf("\tpackage %s\n", lpkg.Name)

	// characterize type info
	if lpkg.Types == nil {
		fmt.Printf("\thas no exported type info\n")
	} else if !lpkg.Types.Complete() {
		fmt.Printf("\thas incomplete exported type info\n")
	} else if len(lpkg.Syntax) == 0 {
		fmt.Printf("\thas complete exported type info\n")
	} else {
		fmt.Printf("\thas complete exported type info and typed ASTs\n")
	}
	if lpkg.Types != nil && lpkg.IllTyped && len(lpkg.Errors) == 0 {
		fmt.Printf("\thas an error among its dependencies\n")
	}

	// source files
	for _, src := range lpkg.GoFiles {
		fmt.Printf("\tfile %s\n", src)
	}

	// imports
	var lines []string
	for importPath, imp := range lpkg.Imports {
		var line string
		if imp.ID == importPath {
			line = fmt.Sprintf("\timport %q", importPath)
		} else {
			line = fmt.Sprintf("\timport %q => %q", importPath, imp.ID)
		}
		lines = append(lines, line)
	}
	sort.Strings(lines)
	for _, line := range lines {
		fmt.Println(line)
	}

	// errors
	for _, err := range lpkg.Errors {
		fmt.Printf("\t%s\n", err)
	}

	// package members (TypeCheck or WholeProgram mode)
	if lpkg.Types != nil {
		qual := types.RelativeTo(lpkg.Types)
		scope := lpkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
<<<<<<< HEAD
			if !obj.Exported() && !*private {
=======
			if !obj.Exported() && !app.Private {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
				continue // skip unexported names
			}

			fmt.Printf("\t%s\n", types.ObjectString(obj, qual))
			if _, ok := obj.(*types.TypeName); ok {
				for _, meth := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
<<<<<<< HEAD
					if !meth.Obj().Exported() && !*private {
=======
					if !meth.Obj().Exported() && !app.Private {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
						continue // skip unexported names
					}
					fmt.Printf("\t%s\n", types.SelectionString(meth, qual))
				}
			}
		}
	}

	fmt.Println()
}

// stringListValue is a flag.Value that accumulates strings.
// e.g. --flag=one --flag=two would produce []string{"one", "two"}.
type stringListValue []string

func newStringListValue(val []string, p *[]string) *stringListValue {
	*p = val
	return (*stringListValue)(p)
}

func (ss *stringListValue) Get() interface{} { return []string(*ss) }

func (ss *stringListValue) String() string { return fmt.Sprintf("%q", *ss) }

func (ss *stringListValue) Set(s string) error { *ss = append(*ss, s); return nil }
