package printf_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/printf"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	printf.Analyzer.Flags.Set("funcs", "Warn,Warnf")
<<<<<<< HEAD
	analysistest.Run(t, testdata, printf.Analyzer, "a", "b")
=======
	analysistest.Run(t, testdata, printf.Analyzer, "a", "b", "nofmt")
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
