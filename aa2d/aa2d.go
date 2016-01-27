// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"bytes"
	"encoding/json"
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
	Ref             interface{}   // JSON reference of the rectangle, if defined
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

const (
	lineChars   = `-|/\:=<>^v`
	leftArrows  = `<^`
	rightArrows = `>v`
	dottedChars = `:=`
)

func (p *Parser) parseContent(e elem, lines [][]byte, xo, y, w, h int) error {
	for ; y < h; y++ {
		line := lines[y]
		if len(line) > 0 {
			switch line[0] {
			case '[':
				return fmt.Errorf("aa2d: reference on line %d defined in the "+
					"middle of document: %s", y, line)
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

func (p *Parser) parseRectangle(
	parent elem,
	lines [][]byte,
	startX, startY int,
	roundUpperLeft bool,
) (*Rectangle, error) {
	if startX+1 == len(lines[startY]) || lines[startY][startX+1] != '-' {
		return nil, &ParseError{X: startX + 1, Y: startY, Err: ErrExpRecLine}
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
			return nil, &ParseError{X: x, Y: startY, Err: ErrExpRecLineOrUpCorn}
		}
	}
	if !found {
		return nil, &ParseError{X: x, Y: startY, Err: ErrNoRecUpRightCorn}
	}
	// go down
	y := startY + 1
	found = false
loopDown:
	for ; y < len(lines); y++ {
		if len(lines[y]) < x {
			return nil, &ParseError{X: x, Y: y, Err: ErrExpRecLineOrLowCorn}
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
			return nil, &ParseError{X: x, Y: y, Err: ErrExpRecLineOrLowCorn}
		}
	}
	if !found {
		return nil, &ParseError{X: x, Y: y, Err: ErrNoRecLowRightCorn}
	}
	// go left
	x--
	for ; x > startX; x-- {
		switch lines[y][x] {
		case '-':
			continue
		default:
			return nil, &ParseError{X: x, Y: y, Err: ErrExpRecHorizontalLine}
		}
	}
	switch lines[y][x] {
	case '\'':
		r.RoundLowerLeft = true
		fallthrough
	case '#':
		break
	default:
		return nil, &ParseError{X: 0, Y: y, Err: ErrExpRecLowCorn}
	}
	// go up
	y--
	for ; y > startY; y-- {
		switch lines[y][startX] {
		case '|':
			continue
		default:
			return nil, &ParseError{X: x, Y: y, Err: ErrExpRecVerticalLine}
		}
	}
	// remove rectangle
	r.remove(lines)
	// add rectangle to parent
	parent.addElem(&r)
	return &r, nil
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

func (r *Rectangle) scale(p *Parser) {
	r.X = r.X*p.xScale + p.xScale/2
	r.Y = r.Y*p.yScale + p.yScale/2
	r.W = r.W*p.xScale - p.xScale
	r.H = r.H*p.yScale - p.yScale
}

// parseLine tries to parse a line starting at position (starX, startY).
// Since the parsing runs from top to bottom and from left to right at this
// stage we only have to consider 4 possible directions (starting from x):
//
//   x-
//  /|\
//
func (p *Parser) parseLine(
	parent elem,
	lines [][]byte,
	startX, startY int,
) error {
	var l Line
	cell := string(lines[startY][startX])

	switch {
	case strings.ContainsAny(cell, leftArrows):
		l.ArrowStart = true
	case strings.ContainsAny(cell, rightArrows):
		return &ParseError{X: startX, Y: startY, Err: ErrRightArrow}
	}

	// try to go to right
	//if ylines[y]

	return nil
}

var matchReference = regexp.MustCompile(`^\[(.+)\]: (.*)$`)

func (p *Parser) parseReference(g *Grid, lines [][]byte, startY int) error {
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
	if strings.HasPrefix(key, "_") {
		if g.Refs == nil {
			g.Refs = make(map[string]map[string]interface{})
		}
		if g.Refs[key] != nil {
			return fmt.Errorf("aa2d: reference on line %d defined twice: %s",
				startY, key)
		}
		g.Refs[key] = refMap
	} else {
		if p.refs[key] != nil {
			return fmt.Errorf("aa2d: reference on line %d defined twice: %s",
				startY, key)
		}
		p.refs[key] = refMap
	}
	return nil
}
