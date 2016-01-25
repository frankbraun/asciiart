// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Package aa2d parses two-dimensional ASCII art into an abstract
// representation.
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
	Rectangles []*Rectangle
	Lines      []*Line
	Polylines  []*Polyline
	Polygons   []*Polygon
	Textlines  []*Textline
}

type Rectangle struct{}
type Line struct{}
type Polyline struct{}
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
