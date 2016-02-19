// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

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

// A Parser for hierarchical ASCII art.
type Parser struct {
	xScale  float64               // scaling factor in x-dimension
	yScale  float64               // scaling factor in y-dimension
	refs    map[string]*reference // JSON references for non-grid elements
	refUsed map[string]bool       // records if given JSON reference was used at least once
}

type reference struct {
	Y   int                    // position of reference on y-axis
	ref map[string]interface{} // the actual reference
}

// A Grid is an abstract representation of hierarchical ASCII art which
// various elements.
type Grid struct {
	W      float64       // size of grid in x-dimension (width)
	H      float64       // size of grid in y-dimension (height)
	XScale float64       // original scaling factor in x-dimension (default 9)
	YScale float64       // original scaling factor in y-dimension (default 16)
	Elems  []interface{} // list of elements on the grid
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

// NewParser returns a new parser for hierarchical ASCII art.
func NewParser() *Parser {
	return &Parser{
		xScale: XScale,
		yScale: YScale,
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
	// init/reset maps
	p.refs = make(map[string]*reference)
	p.refUsed = make(map[string]bool)
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
	g.XScale = p.xScale
	g.YScale = p.yScale
	g.Elems = make([]interface{}, 0)
	if err := p.parseContent(&g, lines, 0, 0, maxLen, len(lines)); err != nil {
		return nil, err
	}
	// make sure all defined reference have been used
	for key := range p.refs {
		if !p.refUsed[key] {
			return nil, &ParseError{X: 0, Y: p.refs[key].Y, Err: ErrRefUnused}
		}
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

func (p *Parser) parseContent(e elem, lines [][]byte, x0, y0, w, h int) error {
	for y := y0; y < y0+h; y++ {
		line := lines[y]
		if len(line) > 0 {
			switch line[0] {
			case '[':
				return &ParseError{X: 0, Y: y, Err: ErrRefMiddle}
			}
		}
	innerLoop:
		for x := x0; x < x0+w && x < len(line); x++ {
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
			case cell == '+':
				err := p.parsePolygon(e, lines, x, y)
				if err != nil {
					return err
				}
			case cell == ' ' || cell == '\t' || cell == '\r':
				continue
			default:
				err := p.parseTextline(e, lines, x, y)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
