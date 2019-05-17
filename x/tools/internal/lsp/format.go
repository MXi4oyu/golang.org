<<<<<<< HEAD
=======
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
package lsp

import (
	"context"
<<<<<<< HEAD
	"go/token"

	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
)

// formatRange formats a document with a given range.
func formatRange(ctx context.Context, v *cache.View, uri protocol.DocumentURI, rng *protocol.Range) ([]protocol.TextEdit, error) {
	f := v.GetFile(source.URI(uri))
	tok, err := f.GetToken()
	if err != nil {
		return nil, err
	}
	var r source.Range
	if rng == nil {
		r.Start = tok.Pos(0)
		r.End = tok.Pos(tok.Size())
	} else {
		r = fromProtocolRange(tok, *rng)
	}
	content, err := f.Read()
	if err != nil {
		return nil, err
	}
	edits, err := source.Format(ctx, f, r)
	if err != nil {
		return nil, err
	}
	return toProtocolEdits(tok, content, edits), nil
}

func toProtocolEdits(tok *token.File, content []byte, edits []source.TextEdit) []protocol.TextEdit {
	if edits == nil {
		return nil
	}
	// When a file ends with an empty line, the newline character is counted
	// as part of the previous line. This causes the formatter to insert
	// another unnecessary newline on each formatting. We handle this case by
	// checking if the file already ends with a newline character.
	hasExtraNewline := content[len(content)-1] == '\n'
	result := make([]protocol.TextEdit, len(edits))
	for i, edit := range edits {
		rng := toProtocolRange(tok, edit.Range)
		// If the edit ends at the end of the file, add the extra line.
		if hasExtraNewline && tok.Offset(edit.Range.End) == len(content) {
			rng.End.Line++
			rng.End.Character = 0
=======
	"fmt"

	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/span"
)

func (s *Server) formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	uri := span.NewURI(params.TextDocument.URI)
	view := s.findView(ctx, uri)
	spn := span.New(uri, span.Point{}, span.Point{})
	return formatRange(ctx, view, spn)
}

func (s *Server) rangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	uri := span.NewURI(params.TextDocument.URI)
	view := s.findView(ctx, uri)
	_, m, err := newColumnMap(ctx, view, uri)
	if err != nil {
		return nil, err
	}
	spn, err := m.RangeSpan(params.Range)
	if err != nil {
		return nil, err
	}
	return formatRange(ctx, view, spn)
}

// formatRange formats a document with a given range.
func formatRange(ctx context.Context, v source.View, s span.Span) ([]protocol.TextEdit, error) {
	f, m, err := newColumnMap(ctx, v, s.URI())
	if err != nil {
		return nil, err
	}
	rng, err := s.Range(m.Converter)
	if err != nil {
		return nil, err
	}
	if rng.Start == rng.End {
		// If we have a single point, assume we want the whole file.
		tok := f.GetToken(ctx)
		if tok == nil {
			return nil, fmt.Errorf("no file information for %s", f.URI())
		}
		rng.End = tok.Pos(tok.Size())
	}
	edits, err := source.Format(ctx, f, rng)
	if err != nil {
		return nil, err
	}
	return ToProtocolEdits(m, edits)
}

func ToProtocolEdits(m *protocol.ColumnMapper, edits []source.TextEdit) ([]protocol.TextEdit, error) {
	if edits == nil {
		return nil, nil
	}
	result := make([]protocol.TextEdit, len(edits))
	for i, edit := range edits {
		rng, err := m.Range(edit.Span)
		if err != nil {
			return nil, err
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
		}
		result[i] = protocol.TextEdit{
			Range:   rng,
			NewText: edit.NewText,
		}
	}
<<<<<<< HEAD
	return result
=======
	return result, nil
}

func FromProtocolEdits(m *protocol.ColumnMapper, edits []protocol.TextEdit) ([]source.TextEdit, error) {
	if edits == nil {
		return nil, nil
	}
	result := make([]source.TextEdit, len(edits))
	for i, edit := range edits {
		spn, err := m.RangeSpan(edit.Range)
		if err != nil {
			return nil, err
		}
		result[i] = source.TextEdit{
			Span:    spn,
			NewText: edit.NewText,
		}
	}
	return result, nil
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
