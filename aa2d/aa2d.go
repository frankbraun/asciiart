// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"bytes"
)

const (
	// XScale is the default scale for the x-axis.
	XScale = 9
	// YScale is the default scale for the y-axis.
	YScale = 16
)

// A Parser for two-dimensional hierarchical ASCII art.
type Parser struct {
	xScale int
	yScale int
}

// A Grid is an abstract representation of two-dimensional hierarchical ASCII
// art which various elements.
type Grid struct {
	Elems []interface{} // list of elements on the grid
	W     int           // size of grid in x-dimension (width)
	H     int           // size of grid in y-dimension (height)
}

// The Rectangle element.
type Rectangle struct {
	Elems           []interface{} // contained elements
	X               int           // x-axis coordinate
	Y               int           // y-axis coordinate
	W               int           // width of rectangle
	H               int           // height of rectangle
	RoundUpperLeft  bool          // rounded upper-left corner
	RoundUpperRight bool          // rounded upper-right corner
	RoundLowerLeft  bool          // rounded lower-left corner
	RoundLowerRight bool          // rounded lower-right corner
	Ref             interface{}   // JSON reference of the rectangle, if defined
}

// The Line element.
type Line struct {
	X1         int         // x-axis coordinate of the start of the line
	Y1         int         // y-axis coordinate of the start of the line
	X2         int         // x-axis coordinate of the end of the line
	Y2         int         // y-axis coordinate of the end of the line
	ArrowStart bool        // arrow at the start of the line
	ArrowEnd   bool        // arrow at the end of the line
	Dotted     bool        // line is dotted
	Ref        interface{} // JSON reference of the line, if defined
}

// The Polyline element.
type Polyline struct {
	X      []int       // x-axis coordinates of points on polyline
	Y      []int       // y-axis coordinates of points on polyline
	Dotted []bool      // polyline segment is dotted (len(Dotted) == len(X)-1)
	Ref    interface{} // JSON reference of the polyline, if defined
}

// The Polygon element.
type Polygon struct {
	Elems  []interface{} // contained elements
	X      []int         // x-axis coordinates of points on polygon
	Y      []int         // y-axis coordinates of points on polygon
	Dotted []bool        // polygon segment is dotted (len(Dotted) == len(X))
	Ref    interface{}   // JSON reference of the polygon, if defined
}

// The Textline element.
type Textline struct {
	X    int    // x-axis coordinate of the start of the text
	Y    int    // y-axis coordinate of the start of the text
	Text string // the actual text string
}

// NewParser returns a new parser for two-dimensional hierarchical ASCII art.
func NewParser() *Parser {
	return &Parser{
		xScale: XScale,
		yScale: YScale,
	}
}

// SetScale sets the scale for parser p.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func (p *Parser) SetScale(xScale, yScale int) {
	p.xScale = xScale
	p.yScale = yScale
}

// Parse parses asciiArt with parser p and returns a grid.
func (p *Parser) Parse(asciiArt string) (*Grid, error) {
	_ = bytes.Split([]byte(asciiArt), []byte("\n"))
	return nil, nil
}
