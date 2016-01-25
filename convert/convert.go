// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package convert implements converters for two-dimensional ASCII art.
package convert

import (
	"io"
)

// ASCIIArt2SVG converts two-dimensional ASCII art read from r to a SVG
// written to w.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func ASCIIArt2SVG(w io.Writer, r io.Reader, xScale, yScale int) error {
	return nil
}
