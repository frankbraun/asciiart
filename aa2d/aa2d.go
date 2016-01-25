// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"bytes"
	"fmt"
)

const (
	// XScale is the default scale for the x-axis of a grid.
	XScale = 9
	// YScale is the default scale for the y-axis of a grid.
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

// SetScale sets the grid scale for parser p.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func (p *Parser) SetScale(xScale, yScale int) error {
	if xScale <= 0 {
		return ErrWrongXScale
	}
	if yScale <= 0 {
		return ErrWrongYScale
	}
	p.xScale = xScale
	p.yScale = yScale
	return nil
}

// Parse parses asciiArt with parser p and returns a grid.
// If there is an error, it will be of type *ParseError.
func (p *Parser) Parse(asciiArt string) (*Grid, error) {
	var (
		g      Grid
		maxLen int
	)
	lines := bytes.Split([]byte(asciiArt), []byte("\n"))
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	g.W = maxLen
	g.H = len(lines)
	if err := g.parse(lines); err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *Grid) parse(lines [][]byte) error {
outerLoop:
	for y, line := range lines {
		for x, cell := range line {
			switch cell {
			case '#':
				if err := g.parseRectangle(lines, x, y, false); err != nil {
					return err
				}
				break outerLoop
			case '.':
				if err := g.parseRectangle(lines, x, y, true); err != nil {
					return err
				}
				break outerLoop
			}
		}
	}
	return nil
}

func (g *Grid) parseRectangle(
	lines [][]byte,
	startX, startY int,
	roundUpperLeft bool,
) error {
	if startX+1 == len(lines[startY]) || lines[startY][startX+1] != '-' {
		return &ParseError{X: startX + 1, Y: startY, Err: ErrExpRecLine}
	}
	for x := startX + 2; x < len(lines[startY]); x++ {
		switch lines[startY][x] {
		case '-':
			continue
		case '#':
			fmt.Println("go down")
			break
		case '.':
			fmt.Println("go down")
			break
		default:
			return &ParseError{X: x, Y: startY, Err: ErrExpRecLineOrCorn}
		}
	}
	return nil
}
