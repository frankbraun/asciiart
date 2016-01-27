// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package svg generates SVGs from an abstract representations of
// two-dimensional ASCII art.
package svg

import (
	"bytes"
	"errors"
	"io"

	"github.com/ajstarks/svgo/float"
	"github.com/frankbraun/asciiart/aa2d"
)

// Generate generates a SVG from grid g and writes it to w.
func Generate(w io.Writer, g *aa2d.Grid) error {
	var buf bytes.Buffer
	blur, ok := g.Refs["_SVG"]["blur"].(bool)
	if !ok {
		blur = true
	}
	rectStyle, err := rectStyle(blur)
	if err != nil {
		return err
	}
	lineStyle, err := lineStyle()
	if err != nil {
		return err
	}
	s := svg.New(&buf) // generate SVG completely before we write it to w
	s.Start(g.W, g.H)
	s.Def()
	if err := setFilter(s); err != nil {
		return err
	}
	s.DefEnd()
	// recursively draw elements
	if err := drawElems(s, g.Elems, rectStyle, lineStyle); err != nil {
		return err
	}
	s.End()
	if _, err := io.Copy(w, &buf); err != nil {
		return err
	}
	return nil
}

func drawElems(
	s *svg.SVG,
	elems []interface{},
	rectStyle, lineStyle []string,
) error {
	for _, elem := range elems {
		switch t := elem.(type) {
		case *aa2d.Rectangle:
			if err := drawRectangle(s, t, rectStyle); err != nil {
				return err
			}
			// recursion
			if err := drawElems(s, t.Elems, rectStyle, lineStyle); err != nil {
				return err
			}
		case *aa2d.Line:
			if err := drawLine(s, t, lineStyle); err != nil {
				return err
			}
		case *aa2d.Polyline:
			if err := drawPolyline(s, t); err != nil {
				return err
			}
		case *aa2d.Polygon:
			if err := drawPolygon(s, t); err != nil {
				return err
			}
			// recursion
			if err := drawElems(s, t.Elems, rectStyle, lineStyle); err != nil {
				return err
			}
		case *aa2d.Textline:
			if err := drawTextline(s, t); err != nil {
				return err
			}
		default:
			return errors.New("svg: unknown type")
		}
	}
	return nil
}

func drawRectangle(s *svg.SVG, r *aa2d.Rectangle, style []string) error {
	s.Rect(r.X, r.Y, r.W, r.H, style...)
	return nil
}

func drawLine(s *svg.SVG, l *aa2d.Line, style []string) error {
	// TODO: draw arrows, if necessary
	s.Line(l.X1, l.Y1, l.X2, l.Y2, style...)
	return nil
}

func drawPolyline(s *svg.SVG, pl *aa2d.Polyline) error {
	return nil
}

func drawPolygon(s *svg.SVG, pg *aa2d.Polygon) error {
	return nil
}

func drawTextline(s *svg.SVG, pg *aa2d.Textline) error {
	return nil
}
