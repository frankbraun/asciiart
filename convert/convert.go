// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package convert implements converters for two-dimensional hierarchical
// ASCII art.
package convert

import (
	"io"
	"io/ioutil"

	"github.com/frankbraun/asciiart/aa2d"
	"github.com/frankbraun/asciiart/svg"
)

// ASCIIArt2SVG converts two-dimensional ASCII art read from r to a SVG
// written to w.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func ASCIIArt2SVG(w io.Writer, r io.Reader, xScale, yScale int) error {
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
	if svg.Generate(w, grid); err != nil {
		return err
	}
	return nil
}
