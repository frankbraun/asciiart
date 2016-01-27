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
	grid, err := p.Parse(aa)
	if err != nil {
		return err
	}
	var traverse func(w io.Writer, elems []interface{}, indent string)
	traverse = func(w io.Writer, elems []interface{}, indent string) {
		for _, elem := range elems {
			switch t := elem.(type) {
			case *aa2d.Rectangle:
				fmt.Fprintf(w, "%s[rect x=%.2f y=%.2f w=%.2f h=%.2f]\n",
					indent, t.X, t.Y, t.W, t.H)
				traverse(w, t.Elems, indent+"  ")
			case *aa2d.Line:
				fmt.Fprintf(w, "%s[line x1=%.2f y1=%.2f x2=%.2f y2=%.2f]\n",
					indent, t.X1, t.Y1, t.X2, t.Y2)
			case *aa2d.Polyline:
				fmt.Fprintf(w, "%s[polyline", indent)
				for i := 0; i < len(t.X); i++ {
					fmt.Fprintf(w, " x%d=%.2f y%d=%.2f",
						i+1, t.X[i], i+1, t.Y[i])
				}
				fmt.Fprintf(w, "]\n")
			case *aa2d.Polygon:
				fmt.Fprintf(w, "%s[polygon", indent)
				for i := 0; i < len(t.X); i++ {
					fmt.Fprintf(w, " x%d=%.2f y%d=%.2f",
						i+1, t.X[i], i+1, t.Y[i])
				}
				traverse(w, t.Elems, indent+"  ")
			case *aa2d.Textline:
				fmt.Fprintf(w, "%s[text x=%.2f y=%.2f t=%q]",
					indent, t.X, t.Y, t.Text)
			}
		}
	}
	traverse(w, grid.Elems, "")
	return nil
}
