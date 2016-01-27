// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
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
}

// A Grid is an abstract representation of two-dimensional hierarchical ASCII
// art which various elements.
type Grid struct {
	W     float64                           // size of grid in x-dimension (width)
	H     float64                           // size of grid in y-dimension (height)
	Refs  map[string]map[string]interface{} // JSON references
	Elems []interface{}                     // list of elements on the grid
}

// The Rectangle element.
type Rectangle struct {
	X               float64       // x-axis coordinate
	Y               float64       // y-axis coordinate
	W               float64       // width of rectangle
	H               float64       // height of rectangle
	RoundUpperLeft  bool          // rounded upper-left corner
	RoundUpperRight bool          // rounded upper-right corner
	RoundLowerLeft  bool          // rounded lower-left corner
	RoundLowerRight bool          // rounded lower-right corner
	Ref             interface{}   // JSON reference of the rectangle, if def
	Elems           []interface{} // contained elements
}

// The Line element.
type Line struct {
	X1         float64                // x-axis coordinate of the start of the line
	Y1         float64                // y-axis coordinate of the start of the line
	X2         float64                // x-axis coordinate of the end of the line
	Y2         float64                // y-axis coordinate of the end of the line
	ArrowStart bool                   // arrow at the start of the line
	ArrowEnd   bool                   // arrow at the end of the line
	Dotted     bool                   // line is dotted
	Ref        map[string]interface{} // JSON reference of the line, if defined
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
	var (
		g      Grid
		maxLen float64
	)
	lines := bytes.Split([]byte(asciiArt), []byte("\n"))
	for _, line := range lines {
		if float64(len(line)) > maxLen {
			maxLen = float64(len(line))
		}
	}
	// remove empty lines at the end
	var rem int
	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) == 0 {
			rem++
		} else {
			break
		}
	}
	lines = lines[:len(lines)-rem]
	g.W = p.xScale * maxLen
	g.H = p.yScale * float64(len(lines))
	// add some extra space at the side for effects
	g.W += p.xScale
	g.H += p.yScale
	g.Elems = make([]interface{}, 0)
	if err := p.parseGrid(&g, lines); err != nil {
		return nil, err
	}
	return &g, nil
}

func (p *Parser) parseGrid(g *Grid, lines [][]byte) error {
	for y, line := range lines {
		if len(line) > 0 {
			switch line[0] {
			case '[':
				if err := g.parseReference(lines, y); err != nil {
					return err
				}
				continue
			}
		}
	innerLoop:
		for x, cell := range line {
			switch cell {
			case '#':
				if err := p.parseRectangle(g, lines, x, y, false); err != nil {
					return err
				}
				goto innerLoop
			case '.':
				if err := p.parseRectangle(g, lines, x, y, true); err != nil {
					return err
				}
				goto innerLoop
			case ' ':
				continue
			default:
				return &ParseError{X: x, Y: y, Err: ErrUnknownCharacter}
			}
		}
	}
	return nil
}

func (p *Parser) parseRectangle(
	g *Grid,
	lines [][]byte,
	startX, startY int,
	roundUpperLeft bool,
) error {
	if startX+1 == len(lines[startY]) || lines[startY][startX+1] != '-' {
		return &ParseError{X: startX + 1, Y: startY, Err: ErrExpRecLine}
	}
	var r Rectangle
	r.X = float64(startX)
	r.Y = float64(startY)
	r.RoundUpperLeft = roundUpperLeft
	// go right
	x := startX + 2
	found := false
loopRight:
	for ; x < len(lines[startY]); x++ {
		switch lines[startY][x] {
		case '-':
			continue
		case '.':
			r.RoundUpperRight = true
			fallthrough
		case '#':
			r.W = float64(x) - r.X + 1
			found = true
			break loopRight
		default:
			return &ParseError{X: x, Y: startY, Err: ErrExpRecLineOrUpCorn}
		}
	}
	if !found {
		return &ParseError{X: x, Y: startY, Err: ErrNoRecUpRightCorn}
	}
	// go down
	y := startY + 1
	found = false
loopDown:
	for ; y < len(lines); y++ {
		if len(lines[y]) < x {
			return &ParseError{X: x, Y: y, Err: ErrExpRecLineOrLowCorn}
		}
		switch lines[y][x] {
		case '|':
			continue
		case '\'':
			r.RoundLowerRight = true
			fallthrough
		case '#':
			r.H = float64(y) - r.Y + 1
			found = true
			break loopDown
		default:
			return &ParseError{X: x, Y: y, Err: ErrExpRecLineOrLowCorn}
		}
	}
	if !found {
		return &ParseError{X: x, Y: y, Err: ErrNoRecLowRightCorn}
	}
	// go left
	x--
	for ; x > startX; x-- {
		switch lines[y][x] {
		case '-':
			continue
		default:
			return &ParseError{X: x, Y: y, Err: ErrExpRecHorizontalLine}
		}
	}
	switch lines[y][x] {
	case '\'':
		r.RoundLowerLeft = true
		fallthrough
	case '#':
		break
	default:
		return &ParseError{X: 0, Y: y, Err: ErrExpRecLowCorn}
	}
	// go up
	y--
	for ; y > startY; y-- {
		switch lines[y][startX] {
		case '|':
			continue
		default:
			return &ParseError{X: x, Y: y, Err: ErrExpRecVerticalLine}
		}
	}
	// remove rectangle
	r.remove(lines)
	// scale
	r.X = r.X*p.xScale + p.xScale/2
	r.Y = r.Y*p.yScale + p.yScale/2
	r.W = r.W*p.xScale - p.xScale
	r.H = r.H*p.yScale - p.yScale
	// add rectangle to grid
	g.Elems = append(g.Elems, r)
	return nil
}

func (r *Rectangle) remove(lines [][]byte) {
	// remove upper and lower line
	for x := int(r.X); x < int(r.X)+int(r.W); x++ {
		lines[int(r.Y)][x] = ' '
		lines[int(r.Y)+int(r.H)-1][x] = ' '
	}
	// remove side lines
	for y := int(r.Y) + 1; y < int(r.Y)+int(r.H)-1; y++ {
		lines[y][int(r.X)] = ' '
		lines[y][int(r.X)+int(r.W)-1] = ' '
	}
}

var matchReference = regexp.MustCompile(`^\[(.+)\]: (.*)$`)

func (g *Grid) parseReference(lines [][]byte, startY int) error {
	m := matchReference.FindStringSubmatch(string(lines[startY]))
	if m == nil {
		return fmt.Errorf("aa2d: cannot parse reference on line %d: %s",
			startY, string(lines[startY]))
	}
	key := m[1]
	jsn := m[2]
	var ref interface{}
	if err := json.Unmarshal([]byte(jsn), &ref); err != nil {
		return err
	}
	refMap, ok := ref.(map[string]interface{})
	if !ok {
		return fmt.Errorf("aa2d: reference on line %d is not a JSON object: %s",
			startY, jsn)
	}
	if g.Refs == nil {
		g.Refs = make(map[string]map[string]interface{})
	}
	if strings.HasPrefix(key, "_") {
		if g.Refs[key] != nil {
			return fmt.Errorf("aa2d: reference on line %d defined twice: %s",
				startY, key)
		}
		g.Refs[key] = refMap
	} else {
		// TODO
		return errors.New("aa2d: non-grid references not implemented yet")
	}
	return nil
}
