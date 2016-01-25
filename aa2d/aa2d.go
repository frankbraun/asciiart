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

type Parser struct {
	xScale int
	yScale int
}

type Grid struct {
	Elems []interface{} // list of elements on the grid
	W     int           // size of grid in x-dimension (width)
	H     int           // size of grid in y-dimension (height)
}

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

type Polyline struct {
	X      []int       // x-axis coordinates of points on polyline
	Y      []int       // y-axis coordinates of points on polyline
	Dotted []bool      // polyline segment is dotted
	Ref    interface{} // JSON reference of the polyline, if defined
}

type Polygon struct{}
type Textline struct{}

func NewParser() *Parser {
	return &Parser{
		xScale: XScale,
		yScale: YScale,
	}
}

// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func (p *Parser) SetScale(xScale, yScale int) {
	p.xScale = xScale
	p.yScale = yScale
}

func (p *Parser) Parse(asciiArt string) (*Grid, error) {
	_ = bytes.Split([]byte(asciiArt), []byte("\n"))
	return nil, nil
}
