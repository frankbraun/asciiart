// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"errors"
	"strings"
)

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

const (
	lineChars   = `-|/\:=<>^v`
	leftArrows  = `<^`
	rightArrows = `>v`
	dottedChars = `:=`
)

// parseLine tries to parse a line starting at position (starX, startY).
// Since the parsing runs from top to bottom and from left to right at this
// stage we only have to consider 4 possible directions (starting from x):
//
//   x-
//  /|\
//
// Possible start situations:
//
//  --  <-  -+   |  :  ^  ^  |  :   /   ^   /  \   ^   \
//               |  :  |  :  +  +  /   /   +    \   \   +
//  ==  <=  =+
//
func (p *Parser) parseLine(
	parent elem,
	lines [][]byte,
	startX, startY int,
) error {
	var l Line
	l.X1 = float64(startX)
	l.Y1 = float64(startY)
	// process start
	cell := lines[startY][startX]
	switch {
	case strings.ContainsAny(string(cell), leftArrows):
		l.ArrowStart = true
	case strings.ContainsAny(string(cell), rightArrows):
		return &ParseError{X: startX, Y: startY, Err: ErrRightArrow}
	case strings.ContainsAny(string(cell), dottedChars):
		l.Dotted = true
	}
	// follow line
	x := startX
	y := startY
	cell = lines[y][y]
	// for arrows we need a head start
	switch cell {
	case '<':
		x++
	case '^':
		y++
		if y < len(lines) {
			if x < len(lines[y]) && (lines[y][x] == '|' || lines[y][x] == ':') {
				break
			} else if x+1 < len(lines[y]) && lines[y][x+1] == '\\' {
				x++
				break
			} else if x > 0 && x-1 < len(lines[y]) && lines[y][x-1] == '/' {
				x--
				break
			}
		}
	}
forLoop:
	for x >= 0 && y < len(lines) && x < len(lines[y]) {
		cell := lines[y][x]
		lines[y][x] = ' ' // nom nom nom
		// move
		switch cell {
		case '-':
			x++
		case '=':
			l.Dotted = true
			x++
		case '|':
			y++
		case ':':
			l.Dotted = true
			y++
		case '/':
			x--
			y++
		case '\\':
			x++
			y++
		case '>', 'v':
			l.ArrowEnd = true
			break forLoop
		case '+':
			// TODO: switch to polyline parsing
			return errors.New("aa2g: '+' not implemented yet")
		case '<', '^':
			return &ParseError{X: x, Y: y, Err: ErrLineLeftArrow}
		}
	}
	l.X2 = float64(x)
	l.Y2 = float64(y)
	// check minimum length
	if x == startX && y == startY {
		return &ParseError{X: x, Y: y, Err: ErrLineTooShort}
	}
	return nil
}
