// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"bytes"
	"context"
	"fmt"
	"go/token"
<<<<<<< HEAD
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/packages/packagestest"
	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
)

// TODO(rstambler): Remove this once Go 1.12 is released as we will end support
// for versions of Go <= 1.10.
var goVersion111 = true

=======
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages/packagestest"
	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/diff"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/lsp/tests"
	"golang.org/x/tools/internal/lsp/xlog"
	"golang.org/x/tools/internal/span"
)

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
func TestLSP(t *testing.T) {
	packagestest.TestAll(t, testLSP)
}

<<<<<<< HEAD
func testLSP(t *testing.T, exporter packagestest.Exporter) {
	const dir = "testdata"

	// We hardcode the expected number of test cases to ensure that all tests
	// are being executed. If a test is added, this number must be changed.
	const expectedCompletionsCount = 44
	const expectedDiagnosticsCount = 14
	const expectedFormatCount = 3
	const expectedDefinitionsCount = 16
	const expectedTypeDefinitionsCount = 2

	files := packagestest.MustCopyFileTree(dir)
	for fragment, operation := range files {
		if trimmed := strings.TrimSuffix(fragment, ".in"); trimmed != fragment {
			delete(files, fragment)
			files[trimmed] = operation
		}
	}
	modules := []packagestest.Module{
		{
			Name:  "golang.org/x/tools/internal/lsp",
			Files: files,
		},
	}
	exported := packagestest.Export(t, exporter, modules)
	defer exported.Cleanup()

	s := &server{
		view: cache.NewView(exported.Config.Dir),
	}
	// Merge the exported.Config with the view.Config.
	cfg := *exported.Config
	cfg.Fset = s.view.Config.Fset
	cfg.Mode = packages.LoadSyntax
	s.view.Config = &cfg

	// Do a first pass to collect special markers for completion.
	if err := exported.Expect(map[string]interface{}{
		"item": func(name string, r packagestest.Range, _, _ string) {
			exported.Mark(name, r)
		},
	}); err != nil {
		t.Fatal(err)
	}

	expectedDiagnostics := make(diagnostics)
	completionItems := make(completionItems)
	expectedCompletions := make(completions)
	expectedFormat := make(formats)
	expectedDefinitions := make(definitions)
	expectedTypeDefinitions := make(definitions)

	// Collect any data that needs to be used by subsequent tests.
	if err := exported.Expect(map[string]interface{}{
		"diag":     expectedDiagnostics.collect,
		"item":     completionItems.collect,
		"complete": expectedCompletions.collect,
		"format":   expectedFormat.collect,
		"godef":    expectedDefinitions.collect,
		"typdef":   expectedTypeDefinitions.collect,
	}); err != nil {
		t.Fatal(err)
	}

	t.Run("Completion", func(t *testing.T) {
		t.Helper()
		if goVersion111 { // TODO(rstambler): Remove this when we no longer support Go 1.10.
			if len(expectedCompletions) != expectedCompletionsCount {
				t.Errorf("got %v completions expected %v", len(expectedCompletions), expectedCompletionsCount)
			}
		}
		expectedCompletions.test(t, exported, s, completionItems)
	})

	t.Run("Diagnostics", func(t *testing.T) {
		t.Helper()
		diagnosticsCount := expectedDiagnostics.test(t, exported, s.view)
		if goVersion111 { // TODO(rstambler): Remove this when we no longer support Go 1.10.
			if diagnosticsCount != expectedDiagnosticsCount {
				t.Errorf("got %v diagnostics expected %v", diagnosticsCount, expectedDiagnosticsCount)
			}
		}
	})

	t.Run("Format", func(t *testing.T) {
		t.Helper()
		if goVersion111 { // TODO(rstambler): Remove this when we no longer support Go 1.10.
			if len(expectedFormat) != expectedFormatCount {
				t.Errorf("got %v formats expected %v", len(expectedFormat), expectedFormatCount)
			}
		}
		expectedFormat.test(t, s)
	})

	t.Run("Definitions", func(t *testing.T) {
		t.Helper()
		if goVersion111 { // TODO(rstambler): Remove this when we no longer support Go 1.10.
			if len(expectedDefinitions) != expectedDefinitionsCount {
				t.Errorf("got %v definitions expected %v", len(expectedDefinitions), expectedDefinitionsCount)
			}
		}
		expectedDefinitions.test(t, s, false)
	})

	t.Run("TypeDefinitions", func(t *testing.T) {
		t.Helper()
		if goVersion111 { // TODO(rstambler): Remove this when we no longer support Go 1.10.
			if len(expectedTypeDefinitions) != expectedTypeDefinitionsCount {
				t.Errorf("got %v type definitions expected %v", len(expectedTypeDefinitions), expectedTypeDefinitionsCount)
			}
		}
		expectedTypeDefinitions.test(t, s, true)
	})
}

type diagnostics map[string][]protocol.Diagnostic
type completionItems map[token.Pos]*protocol.CompletionItem
type completions map[token.Position][]token.Pos
type formats map[string]string
type definitions map[protocol.Location]protocol.Location

func (d diagnostics) test(t *testing.T, exported *packagestest.Exported, v *cache.View) int {
	count := 0
	for filename, want := range d {
		f := v.GetFile(source.ToURI(filename))
		sourceDiagnostics, err := source.Diagnostics(context.Background(), f)
		if err != nil {
			t.Fatal(err)
		}
		got := toProtocolDiagnostics(v, sourceDiagnostics[filename])
		sorted(got)
		if equal := reflect.DeepEqual(want, got); !equal {
			t.Error(diffD(filename, want, got))
		}
		count += len(want)
	}
	return count
}

func (d diagnostics) collect(pos token.Position, msg string) {
	if _, ok := d[pos.Filename]; !ok {
		d[pos.Filename] = []protocol.Diagnostic{}
	}
	// If a file has an empty diagnostics, mark that and return. This allows us
	// to avoid testing diagnostics in files that may have a lot of them.
	if msg == "" {
		return
	}
	line := float64(pos.Line - 1)
	col := float64(pos.Column - 1)
	want := protocol.Diagnostic{
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      line,
				Character: col,
			},
			End: protocol.Position{
				Line:      line,
				Character: col,
			},
		},
		Severity: protocol.SeverityError,
		Source:   "LSP",
		Message:  msg,
	}
	d[pos.Filename] = append(d[pos.Filename], want)
}

func (c completions) test(t *testing.T, exported *packagestest.Exported, s *server, items completionItems) {
	for src, itemList := range c {
		var want []protocol.CompletionItem
		for _, pos := range itemList {
			want = append(want, *items[pos])
		}
		list, err := s.Completion(context.Background(), &protocol.CompletionParams{
			TextDocumentPositionParams: protocol.TextDocumentPositionParams{
				TextDocument: protocol.TextDocumentIdentifier{
					URI: protocol.DocumentURI(source.ToURI(src.Filename)),
				},
				Position: protocol.Position{
					Line:      float64(src.Line - 1),
					Character: float64(src.Column - 1),
				},
			},
		})
		var got []protocol.CompletionItem
		for _, item := range list.Items {
			// Skip all types with no details (builtin types).
			if item.Detail == "" && item.Kind == float64(protocol.TypeParameterCompletion) {
				continue
			}
			// Skip remaining builtin types.
			trimmed := item.Label
			if i := strings.Index(trimmed, "("); i >= 0 {
				trimmed = trimmed[:i]
			}
			switch trimmed {
			case "append", "cap", "close", "complex", "copy", "delete",
				"error", "false", "imag", "iota", "len", "make", "new",
				"nil", "panic", "print", "println", "real", "recover", "true":
=======
type runner struct {
	server *Server
	data   *tests.Data
}

func testLSP(t *testing.T, exporter packagestest.Exporter) {
	ctx := context.Background()

	data := tests.Load(t, exporter, "testdata")
	defer data.Exported.Cleanup()

	log := xlog.New(xlog.StdSink{})
	r := &runner{
		server: &Server{
			views:       []*cache.View{cache.NewView(ctx, log, "lsp_test", span.FileURI(data.Config.Dir), &data.Config)},
			undelivered: make(map[span.URI][]source.Diagnostic),
			log:         log,
		},
		data: data,
	}
	tests.Run(t, r, data)
}

func (r *runner) Diagnostics(t *testing.T, data tests.Diagnostics) {
	v := r.server.views[0]
	for uri, want := range data {
		results, err := source.Diagnostics(context.Background(), v, uri)
		if err != nil {
			t.Fatal(err)
		}
		got := results[uri]
		if diff := diffDiagnostics(uri, want, got); diff != "" {
			t.Error(diff)
		}
	}
}

func sortDiagnostics(d []source.Diagnostic) {
	sort.Slice(d, func(i int, j int) bool {
		if r := span.Compare(d[i].Span, d[j].Span); r != 0 {
			return r < 0
		}
		return d[i].Message < d[j].Message
	})
}

// diffDiagnostics prints the diff between expected and actual diagnostics test
// results.
func diffDiagnostics(uri span.URI, want, got []source.Diagnostic) string {
	sortDiagnostics(want)
	sortDiagnostics(got)
	if len(got) != len(want) {
		return summarizeDiagnostics(-1, want, got, "different lengths got %v want %v", len(got), len(want))
	}
	for i, w := range want {
		g := got[i]
		if w.Message != g.Message {
			return summarizeDiagnostics(i, want, got, "incorrect Message got %v want %v", g.Message, w.Message)
		}
		if span.ComparePoint(w.Start(), g.Start()) != 0 {
			return summarizeDiagnostics(i, want, got, "incorrect Start got %v want %v", g.Start(), w.Start())
		}
		// Special case for diagnostics on parse errors.
		if strings.Contains(string(uri), "noparse") {
			if span.ComparePoint(g.Start(), g.End()) != 0 || span.ComparePoint(w.Start(), g.End()) != 0 {
				return summarizeDiagnostics(i, want, got, "incorrect End got %v want %v", g.End(), w.Start())
			}
		} else if !g.IsPoint() { // Accept any 'want' range if the diagnostic returns a zero-length range.
			if span.ComparePoint(w.End(), g.End()) != 0 {
				return summarizeDiagnostics(i, want, got, "incorrect End got %v want %v", g.End(), w.End())
			}
		}
		if w.Severity != g.Severity {
			return summarizeDiagnostics(i, want, got, "incorrect Severity got %v want %v", g.Severity, w.Severity)
		}
		if w.Source != g.Source {
			return summarizeDiagnostics(i, want, got, "incorrect Source got %v want %v", g.Source, w.Source)
		}
	}
	return ""
}

func summarizeDiagnostics(i int, want []source.Diagnostic, got []source.Diagnostic, reason string, args ...interface{}) string {
	msg := &bytes.Buffer{}
	fmt.Fprint(msg, "diagnostics failed")
	if i >= 0 {
		fmt.Fprintf(msg, " at %d", i)
	}
	fmt.Fprint(msg, " because of ")
	fmt.Fprintf(msg, reason, args...)
	fmt.Fprint(msg, ":\nexpected:\n")
	for _, d := range want {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	fmt.Fprintf(msg, "got:\n")
	for _, d := range got {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	return msg.String()
}

func (r *runner) Completion(t *testing.T, data tests.Completions, snippets tests.CompletionSnippets, items tests.CompletionItems) {
	for src, itemList := range data {
		var want []source.CompletionItem
		for _, pos := range itemList {
			want = append(want, *items[pos])
		}

		list := r.runCompletion(t, src)

		wantBuiltins := strings.Contains(string(src.URI()), "builtins")
		var got []protocol.CompletionItem
		for _, item := range list.Items {
			if !wantBuiltins && isBuiltin(item) {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
				continue
			}
			got = append(got, item)
		}
<<<<<<< HEAD
		if err != nil {
			t.Fatalf("completion failed for %s:%v:%v: %v", filepath.Base(src.Filename), src.Line, src.Column, err)
		}
		if diff := diffC(src, want, got); diff != "" {
			t.Errorf(diff)
=======
		if diff := diffCompletionItems(t, src, want, got); diff != "" {
			t.Errorf("%s: %s", src, diff)
		}
	}
	// Make sure we don't crash completing the first position in file set.
	firstPos, err := span.NewRange(r.data.Exported.ExpectFileSet, 1, 2).Span()
	if err != nil {
		t.Fatal(err)
	}
	_ = r.runCompletion(t, firstPos)

	r.checkCompletionSnippets(t, snippets, items)
}

func (r *runner) checkCompletionSnippets(t *testing.T, data tests.CompletionSnippets, items tests.CompletionItems) {
	origPlaceHolders := r.server.usePlaceholders
	origTextFormat := r.server.insertTextFormat
	defer func() {
		r.server.usePlaceholders = origPlaceHolders
		r.server.insertTextFormat = origTextFormat
	}()

	r.server.insertTextFormat = protocol.SnippetTextFormat
	for _, usePlaceholders := range []bool{true, false} {
		r.server.usePlaceholders = usePlaceholders

		for src, want := range data {
			list := r.runCompletion(t, src)

			wantCompletion := items[want.CompletionItem]
			var gotItem *protocol.CompletionItem
			for _, item := range list.Items {
				if item.Label == wantCompletion.Label {
					gotItem = &item
					break
				}
			}

			if gotItem == nil {
				t.Fatalf("%s: couldn't find completion matching %q", src.URI(), wantCompletion.Label)
			}

			var expected string
			if usePlaceholders {
				expected = want.PlaceholderSnippet
			} else {
				expected = want.PlainSnippet
			}

			if expected != gotItem.TextEdit.NewText {
				t.Errorf("%s: expected snippet %q, got %q", src, expected, gotItem.TextEdit.NewText)
			}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		}
	}
}

<<<<<<< HEAD
func (c completions) collect(src token.Position, expected []token.Pos) {
	c[src] = expected
}

func (i completionItems) collect(pos token.Pos, label, detail, kind string) {
	var k protocol.CompletionItemKind
	switch kind {
	case "struct":
		k = protocol.StructCompletion
	case "func":
		k = protocol.FunctionCompletion
	case "var":
		k = protocol.VariableCompletion
	case "type":
		k = protocol.TypeParameterCompletion
	case "field":
		k = protocol.FieldCompletion
	case "interface":
		k = protocol.InterfaceCompletion
	case "const":
		k = protocol.ConstantCompletion
	case "method":
		k = protocol.MethodCompletion
	case "package":
		k = protocol.ModuleCompletion
	}
	i[pos] = &protocol.CompletionItem{
		Label:  label,
		Detail: detail,
		Kind:   float64(k),
	}
}

func (f formats) test(t *testing.T, s *server) {
	for filename, gofmted := range f {
		edits, err := s.Formatting(context.Background(), &protocol.DocumentFormattingParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.DocumentURI(source.ToURI(filename)),
			},
		})
		if err != nil || len(edits) == 0 {
=======
func (r *runner) runCompletion(t *testing.T, src span.Span) *protocol.CompletionList {
	t.Helper()
	list, err := r.server.Completion(context.Background(), &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.NewURI(src.URI()),
			},
			Position: protocol.Position{
				Line:      float64(src.Start().Line() - 1),
				Character: float64(src.Start().Column() - 1),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return list
}

func isBuiltin(item protocol.CompletionItem) bool {
	// If a type has no detail, it is a builtin type.
	if item.Detail == "" && item.Kind == protocol.TypeParameterCompletion {
		return true
	}
	// Remaining builtin constants, variables, interfaces, and functions.
	trimmed := item.Label
	if i := strings.Index(trimmed, "("); i >= 0 {
		trimmed = trimmed[:i]
	}
	switch trimmed {
	case "append", "cap", "close", "complex", "copy", "delete",
		"error", "false", "imag", "iota", "len", "make", "new",
		"nil", "panic", "print", "println", "real", "recover", "true":
		return true
	}
	return false
}

// diffCompletionItems prints the diff between expected and actual completion
// test results.
func diffCompletionItems(t *testing.T, spn span.Span, want []source.CompletionItem, got []protocol.CompletionItem) string {
	if len(got) != len(want) {
		return summarizeCompletionItems(-1, want, got, "different lengths got %v want %v", len(got), len(want))
	}
	for i, w := range want {
		g := got[i]
		if w.Label != g.Label {
			return summarizeCompletionItems(i, want, got, "incorrect Label got %v want %v", g.Label, w.Label)
		}
		if w.Detail != g.Detail {
			return summarizeCompletionItems(i, want, got, "incorrect Detail got %v want %v", g.Detail, w.Detail)
		}
		if wkind := toProtocolCompletionItemKind(w.Kind); wkind != g.Kind {
			return summarizeCompletionItems(i, want, got, "incorrect Kind got %v want %v", g.Kind, wkind)
		}
	}
	return ""
}

func summarizeCompletionItems(i int, want []source.CompletionItem, got []protocol.CompletionItem, reason string, args ...interface{}) string {
	msg := &bytes.Buffer{}
	fmt.Fprint(msg, "completion failed")
	if i >= 0 {
		fmt.Fprintf(msg, " at %d", i)
	}
	fmt.Fprint(msg, " because of ")
	fmt.Fprintf(msg, reason, args...)
	fmt.Fprint(msg, ":\nexpected:\n")
	for _, d := range want {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	fmt.Fprintf(msg, "got:\n")
	for _, d := range got {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	return msg.String()
}

func (r *runner) Format(t *testing.T, data tests.Formats) {
	ctx := context.Background()
	for _, spn := range data {
		uri := spn.URI()
		filename, err := uri.Filename()
		if err != nil {
			t.Fatal(err)
		}
		gofmted := string(r.data.Golden("gofmt", filename, func(golden string) error {
			cmd := exec.Command("gofmt", filename)
			stdout, err := os.Create(golden)
			if err != nil {
				return err
			}
			defer stdout.Close()
			cmd.Stdout = stdout
			cmd.Run() // ignore error, sometimes we have intentionally ungofmt-able files
			return nil
		}))

		edits, err := r.server.Formatting(context.Background(), &protocol.DocumentFormattingParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.NewURI(uri),
			},
		})
		if err != nil {
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
			if gofmted != "" {
				t.Error(err)
			}
			continue
		}
<<<<<<< HEAD
		edit := edits[0]
		if edit.NewText != gofmted {
			t.Errorf("formatting failed: (got: %s), (expected: %s)", edit.NewText, gofmted)
=======
		_, m, err := newColumnMap(ctx, r.server.findView(ctx, uri), uri)
		if err != nil {
			t.Error(err)
		}
		sedits, err := FromProtocolEdits(m, edits)
		if err != nil {
			t.Error(err)
		}
		ops := source.EditsToDiff(sedits)
		got := strings.Join(diff.ApplyEdits(diff.SplitLines(string(m.Content)), ops), "")
		if gofmted != got {
			t.Errorf("format failed for %s, expected:\n%v\ngot:\n%v", filename, gofmted, got)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		}
	}
}

<<<<<<< HEAD
func (f formats) collect(pos token.Position) {
	cmd := exec.Command("gofmt", pos.Filename)
	stdout := bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Run() // ignore error, sometimes we have intentionally ungofmt-able files
	f[pos.Filename] = stdout.String()
}

func (d definitions) test(t *testing.T, s *server, typ bool) {
	for src, target := range d {
		params := &protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: src.URI,
			},
			Position: src.Range.Start,
		}
		var locs []protocol.Location
		var err error
		if typ {
			locs, err = s.TypeDefinition(context.Background(), params)
		} else {
			locs, err = s.Definition(context.Background(), params)
		}
		if err != nil {
			t.Fatal(err)
=======
func (r *runner) Definition(t *testing.T, data tests.Definitions) {
	for _, d := range data {
		sm := r.mapper(d.Src.URI())
		loc, err := sm.Location(d.Src)
		if err != nil {
			t.Fatalf("failed for %v: %v", d.Src, err)
		}
		params := &protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: loc.URI},
			Position:     loc.Range.Start,
		}
		var locs []protocol.Location
		var hover *protocol.Hover
		if d.IsType {
			locs, err = r.server.TypeDefinition(context.Background(), params)
		} else {
			locs, err = r.server.Definition(context.Background(), params)
			if err != nil {
				t.Fatalf("failed for %v: %v", d.Src, err)
			}
			hover, err = r.server.Hover(context.Background(), params)
		}
		if err != nil {
			t.Fatalf("failed for %v: %v", d.Src, err)
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		}
		if len(locs) != 1 {
			t.Errorf("got %d locations for definition, expected 1", len(locs))
		}
<<<<<<< HEAD
		if locs[0] != target {
			t.Errorf("for %v got %v want %v", src, locs[0], target)
=======
		locURI := span.NewURI(locs[0].URI)
		lm := r.mapper(locURI)
		if def, err := lm.Span(locs[0]); err != nil {
			t.Fatalf("failed for %v: %v", locs[0], err)
		} else if def != d.Def {
			t.Errorf("for %v got %v want %v", d.Src, def, d.Def)
		}
		if hover != nil {
			tag := fmt.Sprintf("hover-%d-%d", d.Def.Start().Line(), d.Def.Start().Column())
			filename, err := d.Def.URI().Filename()
			if err != nil {
				t.Fatalf("failed for %v: %v", d.Def, err)
			}
			expectHover := string(r.data.Golden(tag, filename, func(golden string) error {
				return ioutil.WriteFile(golden, []byte(hover.Contents.Value), 0666)
			}))
			if hover.Contents.Value != expectHover {
				t.Errorf("for %v got %q want %q", d.Src, hover.Contents.Value, expectHover)
			}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		}
	}
}

<<<<<<< HEAD
func (d definitions) collect(fset *token.FileSet, src, target packagestest.Range) {
	sRange := source.Range{Start: src.Start, End: src.End}
	sLoc := toProtocolLocation(fset, sRange)
	tRange := source.Range{Start: target.Start, End: target.End}
	tLoc := toProtocolLocation(fset, tRange)
	d[sLoc] = tLoc
}

// diffD prints the diff between expected and actual diagnostics test results.
func diffD(filename string, want, got []protocol.Diagnostic) string {
	msg := &bytes.Buffer{}
	fmt.Fprintf(msg, "diagnostics failed for %s:\nexpected:\n", filename)
	for _, d := range want {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	fmt.Fprintf(msg, "got:\n")
	for _, d := range got {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	return msg.String()
}

// diffC prints the diff between expected and actual completion test results.
func diffC(pos token.Position, want, got []protocol.CompletionItem) string {
	if len(got) != len(want) {
		goto Failed
	}
	for i, w := range want {
		g := got[i]
		if w.Label != g.Label {
			goto Failed
		}
		if w.Detail != g.Detail {
			goto Failed
		}
		if w.Kind != g.Kind {
			goto Failed
		}
	}
	return ""
Failed:
	msg := &bytes.Buffer{}
	fmt.Fprintf(msg, "completion failed for %s:%v:%v:\nexpected:\n", filepath.Base(pos.Filename), pos.Line, pos.Column)
	for _, d := range want {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	fmt.Fprintf(msg, "got:\n")
	for _, d := range got {
		fmt.Fprintf(msg, "  %v\n", d)
	}
	return msg.String()
}
=======
func (r *runner) Highlight(t *testing.T, data tests.Highlights) {
	for name, locations := range data {
		m := r.mapper(locations[0].URI())
		loc, err := m.Location(locations[0])
		if err != nil {
			t.Fatalf("failed for %v: %v", locations[0], err)
		}
		params := &protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{URI: loc.URI},
			Position:     loc.Range.Start,
		}
		highlights, err := r.server.DocumentHighlight(context.Background(), params)
		if err != nil {
			t.Fatal(err)
		}
		if len(highlights) != len(locations) {
			t.Fatalf("got %d highlights for %s, expected %d", len(highlights), name, len(locations))
		}
		for i := range highlights {
			if h, err := m.RangeSpan(highlights[i].Range); err != nil {
				t.Fatalf("failed for %v: %v", highlights[i], err)
			} else if h != locations[i] {
				t.Errorf("want %v, got %v\n", locations[i], h)
			}
		}
	}
}

func (r *runner) Symbol(t *testing.T, data tests.Symbols) {
	for uri, expectedSymbols := range data {
		params := &protocol.DocumentSymbolParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: string(uri),
			},
		}
		symbols, err := r.server.DocumentSymbol(context.Background(), params)
		if err != nil {
			t.Fatal(err)
		}

		if len(symbols) != len(expectedSymbols) {
			t.Errorf("want %d top-level symbols in %v, got %d", len(expectedSymbols), uri, len(symbols))
			continue
		}
		if diff := r.diffSymbols(uri, expectedSymbols, symbols); diff != "" {
			t.Error(diff)
		}
	}
}

func (r *runner) diffSymbols(uri span.URI, want []source.Symbol, got []protocol.DocumentSymbol) string {
	sort.Slice(want, func(i, j int) bool { return want[i].Name < want[j].Name })
	sort.Slice(got, func(i, j int) bool { return got[i].Name < got[j].Name })
	m := r.mapper(uri)
	if len(got) != len(want) {
		return summarizeSymbols(-1, want, got, "different lengths got %v want %v", len(got), len(want))
	}
	for i, w := range want {
		g := got[i]
		if w.Name != g.Name {
			return summarizeSymbols(i, want, got, "incorrect name got %v want %v", g.Name, w.Name)
		}
		if wkind := toProtocolSymbolKind(w.Kind); wkind != g.Kind {
			return summarizeSymbols(i, want, got, "incorrect kind got %v want %v", g.Kind, wkind)
		}
		spn, err := m.RangeSpan(g.SelectionRange)
		if err != nil {
			return summarizeSymbols(i, want, got, "%v", err)
		}
		if w.SelectionSpan != spn {
			return summarizeSymbols(i, want, got, "incorrect span got %v want %v", spn, w.SelectionSpan)
		}
		if msg := r.diffSymbols(uri, w.Children, g.Children); msg != "" {
			return fmt.Sprintf("children of %s: %s", w.Name, msg)
		}
	}
	return ""
}

func summarizeSymbols(i int, want []source.Symbol, got []protocol.DocumentSymbol, reason string, args ...interface{}) string {
	msg := &bytes.Buffer{}
	fmt.Fprint(msg, "document symbols failed")
	if i >= 0 {
		fmt.Fprintf(msg, " at %d", i)
	}
	fmt.Fprint(msg, " because of ")
	fmt.Fprintf(msg, reason, args...)
	fmt.Fprint(msg, ":\nexpected:\n")
	for _, s := range want {
		fmt.Fprintf(msg, "  %v %v %v\n", s.Name, s.Kind, s.SelectionSpan)
	}
	fmt.Fprintf(msg, "got:\n")
	for _, s := range got {
		fmt.Fprintf(msg, "  %v %v %v\n", s.Name, s.Kind, s.SelectionRange)
	}
	return msg.String()
}

func (r *runner) Signature(t *testing.T, data tests.Signatures) {
	for spn, expectedSignatures := range data {
		m := r.mapper(spn.URI())
		loc, err := m.Location(spn)
		if err != nil {
			t.Fatalf("failed for %v: %v", loc, err)
		}
		gotSignatures, err := r.server.SignatureHelp(context.Background(), &protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.NewURI(spn.URI()),
			},
			Position: loc.Range.Start,
		})
		if err != nil {
			t.Fatal(err)
		}

		if diff := diffSignatures(spn, expectedSignatures, gotSignatures); diff != "" {
			t.Error(diff)
		}
	}
}

func diffSignatures(spn span.Span, want source.SignatureInformation, got *protocol.SignatureHelp) string {
	decorate := func(f string, args ...interface{}) string {
		return fmt.Sprintf("Invalid signature at %s: %s", spn, fmt.Sprintf(f, args...))
	}

	if len(got.Signatures) != 1 {
		return decorate("wanted 1 signature, got %d", len(got.Signatures))
	}

	if got.ActiveSignature != 0 {
		return decorate("wanted active signature of 0, got %f", got.ActiveSignature)
	}

	if want.ActiveParameter != int(got.ActiveParameter) {
		return decorate("wanted active parameter of %d, got %f", want.ActiveParameter, got.ActiveParameter)
	}

	gotSig := got.Signatures[int(got.ActiveSignature)]

	if want.Label != gotSig.Label {
		return decorate("wanted label %q, got %q", want.Label, gotSig.Label)
	}

	var paramParts []string
	for _, p := range gotSig.Parameters {
		paramParts = append(paramParts, p.Label)
	}
	paramsStr := strings.Join(paramParts, ", ")
	if !strings.Contains(gotSig.Label, paramsStr) {
		return decorate("expected signature %q to contain params %q", gotSig.Label, paramsStr)
	}

	return ""
}

func (r *runner) Link(t *testing.T, data tests.Links) {
	for uri, wantLinks := range data {
		m := r.mapper(uri)
		gotLinks, err := r.server.DocumentLink(context.Background(), &protocol.DocumentLinkParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: protocol.NewURI(uri),
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		links := make(map[span.Span]string, len(wantLinks))
		for _, link := range wantLinks {
			links[link.Src] = link.Target
		}
		for _, link := range gotLinks {
			spn, err := m.RangeSpan(link.Range)
			if err != nil {
				t.Fatal(err)
			}
			if target, ok := links[spn]; ok {
				delete(links, spn)
				if target != link.Target {
					t.Errorf("for %v want %v, got %v\n", spn, link.Target, target)
				}
			} else {
				t.Errorf("unexpected link %v:%v\n", spn, link.Target)
			}
		}
		for spn, target := range links {
			t.Errorf("missing link %v:%v\n", spn, target)
		}
	}
}

func (r *runner) mapper(uri span.URI) *protocol.ColumnMapper {
	fname, err := uri.Filename()
	if err != nil {
		return nil
	}
	fset := r.data.Exported.ExpectFileSet
	var f *token.File
	fset.Iterate(func(check *token.File) bool {
		if check.Name() == fname {
			f = check
			return false
		}
		return true
	})
	if f == nil {
		return nil
	}
	content, err := r.data.Exported.FileContents(f.Name())
	if err != nil {
		return nil
	}
	return protocol.NewColumnMapper(uri, fset, f, content)
}

func TestBytesOffset(t *testing.T) {
	tests := []struct {
		text string
		pos  protocol.Position
		want int
	}{
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 0}, want: 0},
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 1}, want: 1},
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 2}, want: 1},
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 3}, want: 5},
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 4}, want: 6},
		{text: `a𐐀b`, pos: protocol.Position{Line: 0, Character: 5}, want: -1},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 0, Character: 3}, want: 3},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 0, Character: 4}, want: -1},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 1, Character: 0}, want: 4},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 1, Character: 3}, want: 7},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 1, Character: 4}, want: -1},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 2, Character: 0}, want: 8},
		{text: "aaa\nbbb\n", pos: protocol.Position{Line: 2, Character: 1}, want: -1},
		{text: "aaa\nbbb\n\n", pos: protocol.Position{Line: 2, Character: 0}, want: 8},
	}

	for i, test := range tests {
		fname := fmt.Sprintf("test %d", i)
		fset := token.NewFileSet()
		f := fset.AddFile(fname, -1, len(test.text))
		f.SetLinesForContent([]byte(test.text))
		mapper := protocol.NewColumnMapper(span.FileURI(fname), fset, f, []byte(test.text))
		got, err := mapper.Point(test.pos)
		if err != nil && test.want != -1 {
			t.Errorf("unexpected error: %v", err)
		}
		if err == nil && got.Offset() != test.want {
			t.Errorf("want %d for %q(Line:%d,Character:%d), but got %d", test.want, test.text, int(test.pos.Line), int(test.pos.Character), got.Offset())
		}
	}
}
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
