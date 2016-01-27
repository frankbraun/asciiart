// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"bytes"
	"strings"
)

const (
	// XScale is the default scale for the x-axis of a grid.
	XScale = 9
	// YScale is the default scale for the y-axis of a grid.
	YScale = 16
)

// A Parser for two-dimensional hierarchical ASCII art.
type Parser struct {
	xScale float64
	yScale float64
	refs   map[string]map[string]interface{}
}

// A Grid is an abstract representation of two-dimensional hierarchical ASCII
// art which various elements.
type Grid struct {
	W     float64                           // size of grid in x-dimension (width)
	H     float64                           // size of grid in y-dimension (height)
	Refs  map[string]map[string]interface{} // JSON references
	Elems []interface{}                     // list of elements on the grid
}

// The Polyline element.
type Polyline struct {
	X      []float64              // x-axis coordinates of points on polyline
	Y      []float64              // y-axis coordinates of points on polyline
	Dotted []bool                 // polyline segment is dotted (len(Dotted) == len(X)-1)
	Ref    map[string]interface{} // JSON reference of the polyline, if defined
}

// The Polygon element.
type Polygon struct {
	X      []float64              // x-axis coordinates of points on polygon
	Y      []float64              // y-axis coordinates of points on polygon
	Dotted []bool                 // polygon segment is dotted (len(Dotted) == len(X))
	Ref    map[string]interface{} // JSON reference of the polygon, if defined
	Elems  []interface{}          // contained elements
}

// The Textline element.
type Textline struct {
	X    float64 // x-axis coordinate of the start of the text
	Y    float64 // y-axis coordinate of the start of the text
	Text string  // the actual text string
}

type elem interface {
	addElem(interface{})
}

func (g *Grid) addElem(e interface{}) {
	g.Elems = append(g.Elems, e)
}

func (r *Rectangle) addElem(e interface{}) {
	r.Elems = append(r.Elems, e)
}

// NewParser returns a new parser for two-dimensional hierarchical ASCII art.
func NewParser() *Parser {
	return &Parser{
		xScale: XScale,
		yScale: YScale,
		refs:   make(map[string]map[string]interface{}),
	}
}

// SetScale sets the grid scale for parser p.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func (p *Parser) SetScale(xScale, yScale float64) error {
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
	var g Grid
	lines := bytes.Split([]byte(asciiArt), []byte("\n"))
	lines = removeEmptyTrailingLines(lines)
	for y, line := range lines {
		if len(line) > 0 && line[0] == '[' {
			if err := p.parseReference(&g, lines, y); err != nil {
				return nil, err
			}
			lines[y] = nil // remove reference line
		}
	}
	lines = removeEmptyTrailingLines(lines)
	var maxLen int
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	g.W = p.xScale * float64(maxLen)
	g.H = p.yScale * float64(len(lines))
	// add some extra space at the side for effects
	g.W += p.xScale
	g.H += p.yScale
	g.Elems = make([]interface{}, 0)
	if err := p.parseContent(&g, lines, 0, 0, maxLen, len(lines)); err != nil {
		return nil, err
	}
	return &g, nil
}

func removeEmptyTrailingLines(lines [][]byte) [][]byte {
	// remove empty lines at the end
	var rem int
	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) == 0 {
			rem++
		} else {
			break
		}
	}
	return lines[:len(lines)-rem]
}

func (p *Parser) parseContent(e elem, lines [][]byte, xo, y, w, h int) error {
	for ; y < h; y++ {
		line := lines[y]
		if len(line) > 0 {
			switch line[0] {
			case '[':
				return &ParseError{X: 0, Y: y, Err: ErrRefMiddle}
			}
		}
	innerLoop:
		for x := xo; x < w && x < len(line); x++ {
			cell := line[x]
			var roundUpperLeft bool
			switch {
			case cell == '.':
				roundUpperLeft = true
				fallthrough
			case cell == '#':
				r, err := p.parseRectangle(e, lines, x, y, roundUpperLeft)
				if err != nil {
					return err
				}
				// recursion
				err = p.parseContent(r, lines,
					int(r.X)+1, int(r.Y)+1, int(r.W)-2, int(r.H)-2)
				if err != nil {
					return err
				}
				// scale rectangle afterwards
				r.scale(p)
				goto innerLoop
			case strings.ContainsAny(string(cell), lineChars):
				err := p.parseLine(e, lines, x, y)
				if err != nil {
					return err
				}
			case cell == ' ':
				continue
			default:
				return &ParseError{X: x, Y: y, Err: ErrUnknownCharacter}
			}
		}
	}
	return nil
}
