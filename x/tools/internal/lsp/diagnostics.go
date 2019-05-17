// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"
<<<<<<< HEAD
	"sort"
=======
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a

	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
<<<<<<< HEAD
)

func (s *server) CacheAndDiagnose(ctx context.Context, uri protocol.DocumentURI, text string) {
	f := s.view.GetFile(source.URI(uri))
	f.SetContent([]byte(text))

	go func() {
		reports, err := source.Diagnostics(ctx, f)
		if err != nil {
			return // handle error?
		}
		for filename, diagnostics := range reports {
			s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
				URI:         protocol.DocumentURI(source.ToURI(filename)),
				Diagnostics: toProtocolDiagnostics(s.view, diagnostics),
			})
		}
	}()
}

func toProtocolDiagnostics(v *cache.View, diagnostics []source.Diagnostic) []protocol.Diagnostic {
	reports := []protocol.Diagnostic{}
	for _, diag := range diagnostics {
		f := v.GetFile(source.ToURI(diag.Filename))
		tok, err := f.GetToken()
		if err != nil {
			continue // handle error?
		}
		pos := fromTokenPosition(tok, diag.Position)
		if !pos.IsValid() {
			continue // handle error?
		}
		reports = append(reports, protocol.Diagnostic{
			Message: diag.Message,
			Range: toProtocolRange(tok, source.Range{
				Start: pos,
				End:   pos,
			}),
			Severity: protocol.SeverityError, // all diagnostics have error severity for now
			Source:   "LSP",
		})
	}
	return reports
}

func sorted(d []protocol.Diagnostic) {
	sort.Slice(d, func(i int, j int) bool {
		if d[i].Range.Start.Line == d[j].Range.Start.Line {
			if d[i].Range.Start.Character == d[j].Range.End.Character {
				return d[i].Message < d[j].Message
			}
			return d[i].Range.Start.Character < d[j].Range.End.Character
		}
		return d[i].Range.Start.Line < d[j].Range.Start.Line
	})
=======
	"golang.org/x/tools/internal/span"
)

func (s *Server) cacheAndDiagnose(ctx context.Context, uri span.URI, content string) error {
	view := s.findView(ctx, uri)
	if err := view.SetContent(ctx, uri, []byte(content)); err != nil {
		return err
	}

	go func() {
		ctx := view.BackgroundContext()
		if ctx.Err() != nil {
			s.log.Errorf(ctx, "canceling diagnostics for %s: %v", uri, ctx.Err())
			return
		}
		reports, err := source.Diagnostics(ctx, view, uri)
		if err != nil {
			s.log.Errorf(ctx, "failed to compute diagnostics for %s: %v", uri, err)
			return
		}

		s.undeliveredMu.Lock()
		defer s.undeliveredMu.Unlock()

		for uri, diagnostics := range reports {
			if err := s.publishDiagnostics(ctx, view, uri, diagnostics); err != nil {
				if s.undelivered == nil {
					s.undelivered = make(map[span.URI][]source.Diagnostic)
				}
				s.undelivered[uri] = diagnostics
				continue
			}
			// In case we had old, undelivered diagnostics.
			delete(s.undelivered, uri)
		}
		// Anytime we compute diagnostics, make sure to also send along any
		// undelivered ones (only for remaining URIs).
		for uri, diagnostics := range s.undelivered {
			s.publishDiagnostics(ctx, view, uri, diagnostics)

			// If we fail to deliver the same diagnostics twice, just give up.
			delete(s.undelivered, uri)
		}
	}()
	return nil
}

func (s *Server) publishDiagnostics(ctx context.Context, view *cache.View, uri span.URI, diagnostics []source.Diagnostic) error {
	protocolDiagnostics, err := toProtocolDiagnostics(ctx, view, diagnostics)
	if err != nil {
		return err
	}
	s.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		Diagnostics: protocolDiagnostics,
		URI:         protocol.NewURI(uri),
	})
	return nil
}

func toProtocolDiagnostics(ctx context.Context, v source.View, diagnostics []source.Diagnostic) ([]protocol.Diagnostic, error) {
	reports := []protocol.Diagnostic{}
	for _, diag := range diagnostics {
		_, m, err := newColumnMap(ctx, v, diag.Span.URI())
		if err != nil {
			return nil, err
		}
		var severity protocol.DiagnosticSeverity
		switch diag.Severity {
		case source.SeverityError:
			severity = protocol.SeverityError
		case source.SeverityWarning:
			severity = protocol.SeverityWarning
		}
		rng, err := m.Range(diag.Span)
		if err != nil {
			return nil, err
		}
		reports = append(reports, protocol.Diagnostic{
			Message:  diag.Message,
			Range:    rng,
			Severity: severity,
			Source:   diag.Source,
		})
	}
	return reports, nil
>>>>>>> bd25a1f6d07d2d464980e6a8576c1ed59bb3950a
}
