// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package convert implements converters for two-dimensional hierarchical
// ASCII art.
package convert

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/frankbraun/asciiart/aa2d"
	"github.com/frankbraun/asciiart/svg"
)

// ASCIIArt2SVG converts two-dimensional ASCII art read from r to a SVG
// written to w.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func ASCIIArt2SVG(w io.Writer, r io.Reader, xScale, yScale float64) error {
	aa, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	p := aa2d.NewParser()
	if err := p.SetScale(xScale, yScale); err != nil {
		return err
	}
	grid, err := p.Parse(string(aa))
	if err != nil {
		return err
	}
	if err := svg.Generate(w, grid); err != nil {
		return err
	}
	return nil
}

// ASCIIArt2Txt converts two-dimensional ASCII art read from r to text
// written to w.
func ASCIIArt2Txt(w io.Writer, r io.Reader) error {
	aa, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = generateTxt(w, string(aa))
	if err != nil {
		return err
	}
	return nil
}

func generateTxt(w io.Writer, aa string) error {
	p := aa2d.NewParser()
	p.SetScale(1, 1)
	grid, err := p.Parse(aa)
	if err != nil {
		return err
	}
	fmtFlt := func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }
	var traverse func(w io.Writer, elems []interface{}, indent string)
	traverse = func(w io.Writer, elems []interface{}, indent string) {
		for _, elem := range elems {
			switch t := elem.(type) {
			case *aa2d.Rectangle:
				fmt.Fprintln(w, indent, "rect", t.X, t.Y, t.W, t.H)
				traverse(w, t.Elems, indent+"  ") // recursion
			case *aa2d.Line:
				fmt.Fprintln(w, indent, "line", t.X1, t.Y1, t.X2, t.Y2)
			case *aa2d.Polyline:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Fprintln(w, indent, "polyline", strings.Join(p, " "))
			case *aa2d.Polygon:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Fprintln(w, indent, "polygon", strings.Join(p, " "))
				traverse(w, t.Elems, indent+"  ") // recursion
			case *aa2d.Textline:
				fmt.Fprintln(w, indent, "textline", t.X, t.Y, t.Text)
			}
		}
	}
	fmt.Fprintln(w, "grid", grid.W, grid.H)
	traverse(w, grid.Elems, " ")
	return nil
}
