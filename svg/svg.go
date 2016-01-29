// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package svg generates SVGs from an abstract representations of
// two-dimensional ASCII art.
package svg

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/ajstarks/svgo/float"
	"github.com/frankbraun/asciiart"
)

// Generate generates a SVG from grid g and writes it to w.
func Generate(w io.Writer, g *asciiart.Grid) error {
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
	textStyle, err := textStyle(g.YScale)
	if err != nil {
		return err
	}
	s := svg.New(&buf)                  // generate complete SVG before writing
	s.Start(g.W+g.XScale, g.H+g.YScale) // add extra space to sides for effects
	s.Def()
	if err := setFilter(s); err != nil {
		return err
	}
	if err := setMarker(s); err != nil {
		return err
	}
	s.DefEnd()
	// recursively draw elements
	err = drawElems(s, g.Elems, rectStyle, lineStyle, textStyle,
		g.XScale, g.YScale)
	if err != nil {
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
	rectStyle, lineStyle, textStyle []string,
	xScale, yScale float64,
) error {
	for _, elem := range elems {
		switch t := elem.(type) {
		case *asciiart.Rectangle:
			if err := drawRectangle(s, t, rectStyle); err != nil {
				return err
			}
			// recursion
			err := drawElems(s, t.Elems, rectStyle, lineStyle, textStyle,
				xScale, yScale)
			if err != nil {
				return err
			}
		case *asciiart.Line:
			if err := drawLine(s, t, lineStyle); err != nil {
				return err
			}
		case *asciiart.Polyline:
			if err := drawPolyline(s, t); err != nil {
				return err
			}
		case *asciiart.Polygon:
			if err := drawPolygon(s, t); err != nil {
				return err
			}
		case *asciiart.Textline:
			if err := drawTextline(s, t, textStyle, xScale, yScale); err != nil {
				return err
			}
		default:
			return errors.New("svg: unknown type")
		}
	}
	return nil
}

func drawRectangle(s *svg.SVG, r *asciiart.Rectangle, style []string) error {
	if r.RoundUpperLeft || r.RoundUpperRight || r.RoundLowerLeft || r.RoundLowerRight {
		// we got rounded corners -> construct rectangle manually as path, also see:
		// https://github.com/dhobsd/asciitosvg/blob/05f2ac06918247a79561b026a6a8011a64a98317/ASCIIToSVG.php#L968-L988
		points := []struct {
			x       float64
			y       float64
			rounded bool
		}{
			{x: r.X, y: r.Y, rounded: r.RoundUpperLeft},
			{x: r.X + r.W, y: r.Y, rounded: r.RoundUpperRight},
			{x: r.X + r.W, y: r.Y + r.H, rounded: r.RoundLowerRight},
			{x: r.X, y: r.Y + r.H, rounded: r.RoundLowerLeft},
		}
		var d string
		point := points[0]
		if point.rounded {
			cX := point.x
			cY := point.y
			sX := point.x
			sY := point.y + 10
			eX := point.x + 10
			eY := point.y
			d += fmt.Sprintf("M %f %f Q %f %f %f %f ", sX, sY, cX, cY, eX, eY)
		} else {
			d += fmt.Sprintf("M %f %f ", point.x, point.y)
		}
		point = points[1]
		if point.rounded {
			cX := point.x
			cY := point.y
			sX := point.x - 10
			sY := point.y
			eX := point.x
			eY := point.y + 10
			d += fmt.Sprintf("L %f %f Q %f %f %f %f ", sX, sY, cX, cY, eX, eY)
		} else {
			d += fmt.Sprintf("L %f %f ", point.x, point.y)
		}
		point = points[2]
		if point.rounded {
			cX := point.x
			cY := point.y
			sX := point.x
			sY := point.y - 10
			eX := point.x - 10
			eY := point.y
			d += fmt.Sprintf("L %f %f Q %f %f %f %f ", sX, sY, cX, cY, eX, eY)
		} else {
			d += fmt.Sprintf("L %f %f ", point.x, point.y)
		}
		point = points[3]
		if point.rounded {
			cX := point.x
			cY := point.y
			sX := point.x + 10
			sY := point.y
			eX := point.x
			eY := point.y - 10
			d += fmt.Sprintf("L %f %f Q %f %f %f %f ", sX, sY, cX, cY, eX, eY)
		} else {
			d += fmt.Sprintf("L %f %f ", point.x, point.y)
		}

		s.Path(d+"Z", style...)
	} else {
		// draw rect element
		s.Rect(r.X, r.Y, r.W, r.H, style...)
	}
	return nil
}

func drawLine(s *svg.SVG, l *asciiart.Line, style []string) error {
	totalStyle := make([]string, len(style))
	copy(totalStyle, style)
	if l.ArrowStart {
		totalStyle = append(totalStyle, `marker-start="url(#iPointer)"`)
	}
	if l.ArrowEnd {
		totalStyle = append(totalStyle, `marker-end="url(#Pointer)"`)
	}
	if l.Dotted {
		totalStyle = append(totalStyle, `stroke-dasharray="5 5"`)
	}
	s.Line(l.X1, l.Y1, l.X2, l.Y2, totalStyle...)
	return nil
}

func drawPolyline(s *svg.SVG, p *asciiart.Polyline) error {
	return nil
}

func drawPolygon(s *svg.SVG, p *asciiart.Polygon) error {
	return nil
}

func drawTextline(
	s *svg.SVG,
	t *asciiart.Textline,
	style []string,
	xScale, yScale float64,
) error {
	// use "magic numbers" inspired by:
	// https://github.com/dhobsd/asciitosvg/blob/05f2ac06918247a79561b026a6a8011a64a98317/ASCIIToSVG.php#L1757-L1758
	s.Text(t.X-0.6*xScale, t.Y+0.3*yScale, t.Text, style...)
	return nil
}
