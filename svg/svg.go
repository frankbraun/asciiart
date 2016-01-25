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

	"github.com/ajstarks/svgo"
	"github.com/frankbraun/asciiart/aa2d"
)

// Generate generates a SVG from g and writes it to w.
func Generate(w io.Writer, g *aa2d.Grid) error {
	var buf bytes.Buffer
	s := svg.New(&buf) // don't write to w in case of error
	s.Start(g.W, g.H)
	for _, elem := range g.Elems {
		switch t := elem.(type) {
		case aa2d.Rectangle:
			if err := drawRectangle(s, &t); err != nil {
				return err
			}
		case aa2d.Line:
			if err := drawLine(s, &t); err != nil {
				return err
			}
		case aa2d.Polyline:
			if err := drawPolyline(s, &t); err != nil {
				return err
			}
		case aa2d.Polygon:
			if err := drawPolygon(s, &t); err != nil {
				return err
			}
		case aa2d.Textline:
			if err := drawTextline(s, &t); err != nil {
				return err
			}
		default:
			return errors.New("svg: unknown type")
		}
	}
	s.End()
	io.Copy(w, &buf)
	return nil
}

func drawRectangle(s *svg.SVG, r *aa2d.Rectangle) error {
	return nil
}

func drawLine(s *svg.SVG, l *aa2d.Line) error {
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
