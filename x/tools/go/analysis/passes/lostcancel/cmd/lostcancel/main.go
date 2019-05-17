<<<<<<< HEAD
// The nilness command applies the golang.org/x/tools/go/analysis/passes/lostcancel
=======
// The lostcancel command applies the golang.org/x/tools/go/analysis/passes/lostcancel
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
// analysis to the specified packages of Go source code.
package main

import (
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(lostcancel.Analyzer) }
