// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

import (
	"errors"

	"github.com/frankbraun/kitchensink/bit"
)

// The Polyline element.
type Polyline struct {
	X          []float64 // x-axis coordinates of points on polyline
	Y          []float64 // y-axis coordinates of points on polyline
	ArrowStart bool      // arrow at the start of the polyline
	ArrowEnd   bool      // arrow at the end of the polyline
	Dotted     []bool    // polyline segment is dotted (len(Dotted) == len(X)-1)
}

// The Polygon element.
type Polygon struct {
	X      []float64 // x-axis coordinates of points on polygon
	Y      []float64 // y-axis coordinates of points on polygon
	Dotted []bool    // polygon segment is dotted (len(Dotted) == len(X))
}

// define edge directions as bit fields
const (
	nEdge  = 1 << iota // north
	neEdge             // northeast
	eEdge              // east
	seEdge             // southeast
	sEdge              // south
	swEdge             // southwest
	wEdge              // west
	nwEdge             // northwest
)

const (
	cornerState = iota
	edgeState
	endState
)

// parsePolyline parses a polyline.
//
//  \|/
//  -+-
//  /|\
//
func (p *Parser) parsePolyline(
	parent elem,
	lines [][]byte,
	l *Line,
) error {
	var pl Polyline
	// take stuff over from parsed line
	pl.X = append(pl.X, l.X1)
	pl.Y = append(pl.Y, l.Y1)
	pl.ArrowStart = l.ArrowStart

	x, y, arrowEnd, dotted, err := followLine(&pl, lines, int(l.X2), int(l.Y2),
		cornerState, 0, l.Dotted, false)
	if err != nil {
		return err
	}
	// add final segment
	pl.X = append(pl.X, float64(x))
	pl.Y = append(pl.Y, float64(y))
	pl.ArrowEnd = arrowEnd
	pl.Dotted = append(pl.Dotted, dotted)
	// scale
	pl.scale(p)
	// add polyline to parent
	parent.addElem(&pl)
	return nil
}

// parsePolygon parses a polygon.
func (p *Parser) parsePolygon(
	parent elem,
	lines [][]byte,
	x, y int,
) error {
	var pg Polygon
	pg.X = append(pg.X, float64(x))
	pg.Y = append(pg.Y, float64(y))

	startX := x
	startY := y
	outEdges := outgoingEdges(lines, x, y)
	switch bit.Count(outEdges) {
	case 0:
		return &ParseError{X: x, Y: y, Err: ErrPolyCornerNoEdge}
	case 1:
		return &ParseError{X: x, Y: y, Err: ErrPolyCornerOneEdge}
	case 2:
	default:
		return &ParseError{X: x, Y: y, Err: ErrPolyCornerTooManyEdges}
	}

	var edge byte
	if outEdges&eEdge > 0 {
		edge = wEdge
		x++
	} else if outEdges&seEdge > 0 {
		edge = nwEdge
		x++
		y++
	} else if outEdges&sEdge > 0 {
		edge = nEdge
		y++
	} else if outEdges&swEdge > 0 {
		edge = neEdge
		x--
		y++
	}

	endX, endY, _, dotted, err := followLine(&pg, lines, x, y,
		edgeState, edge, false, true)
	if err != nil {
		return err
	}
	// check final point
	if endX != startX || endY != startY {
		return &ParseError{X: x, Y: y, Err: ErrPolygonNotClosed}
	}
	pg.Dotted = append(pg.Dotted, dotted)
	// scale
	pg.scale(p)
	// add polygon to parent
	parent.addElem(&pg)
	return nil
}

type appender interface {
	appendPoint(x, y float64, dotted bool)
}

func (pl *Polyline) appendPoint(x, y float64, dotted bool) {
	pl.X = append(pl.X, x)
	pl.Y = append(pl.Y, y)
	pl.Dotted = append(pl.Dotted, dotted)
}

func (pg *Polygon) appendPoint(x, y float64, dotted bool) {
	pg.X = append(pg.X, x)
	pg.Y = append(pg.Y, y)
	pg.Dotted = append(pg.Dotted, dotted)
}

func followLine(
	poly appender,
	lines [][]byte,
	x, y int,
	state, edge byte,
	dotted, endEmptyCorner bool,
) (outX, outY int, arrowEnd, outDotted bool, err error) {
	for {
		switch state {
		case cornerState:
			lines[y][x] = ' ' // nom nom nom
			// process corner (+)
			outEdges := outgoingEdges(lines, x, y)
			switch bit.Count(outEdges) {
			case 0:
				if !endEmptyCorner {
					return 0, 0, false, false,
						&ParseError{X: x, Y: y, Err: ErrPolyCornerOneEdge}
				}
				state = endState
			case 1:
				// add segement
				poly.appendPoint(float64(x), float64(y), dotted)
				dotted = false
				// follow edge
				switch {
				case outEdges&nEdge > 0:
					edge = sEdge
					y--
				case outEdges&neEdge > 0:
					edge = swEdge
					x++
					y--
				case outEdges&eEdge > 0:
					edge = wEdge
					x++
				case outEdges&seEdge > 0:
					edge = nwEdge
					x++
					y++
				case outEdges&sEdge > 0:
					edge = nEdge
					y++
				case outEdges&swEdge > 0:
					edge = neEdge
					x--
					y++
				case outEdges&wEdge > 0:
					edge = eEdge
					x--
				case outEdges&nwEdge > 0:
					edge = seEdge
					x--
					y--
				}
				dotted = startsDotted(lines, x, y, edge)
				state = edgeState
			default:
				// TODO: split
				return 0, 0, false, false,
					errors.New("poly split not implemented")
			}
		case edgeState:
			next, a, d := procEdge(lines, &x, &y, edge)
			if a {
				arrowEnd = true
			}
			if d {
				dotted = true
			}
			state = next
		case endState:
			lines[y][x] = ' ' // nom nom nom
			outX = x
			outY = y
			outDotted = dotted
			return
		}
	}
}

func startsDotted(lines [][]byte, x, y int, incomingEdge byte) bool {
	cell := lines[y][x]
	if (incomingEdge == nEdge || incomingEdge == sEdge) && cell == ':' {
		return true
	}
	if (incomingEdge == eEdge || incomingEdge == wEdge) && cell == '=' {
		return true
	}
	return false
}

// procEdge processes an edge at position x,y and returns the next state.
// Three possible outcomes:
//   1. we can continue the line:
//      - nextState = edgeState
//      - x and y changed accordingly
//   2. we hit a corner:
//      - nextState = cornerState
//      - x and y changed accordingly
//   3. we reach the corner of the grid or cannot continue the line:
//      - nextState = endState,
//      - x and y unchanged
func procEdge(
	lines [][]byte,
	x, y *int,
	incomingEdge byte,
) (nextState byte, arrowEnd, dotted bool) {
	lines[*y][*x] = ' ' // nom nom nom
	switch incomingEdge {
	case nEdge:
		if len(lines) > *y+1 && len(lines[*y+1]) > *x {
			cell := lines[*y+1][*x]
			switch cell {
			case ':':
				dotted = true
				fallthrough
			case '|':
				*y++
				nextState = edgeState
			case '+':
				*y++
				nextState = cornerState
			case 'v':
				arrowEnd = true
				*y++
				nextState = cornerState
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case neEdge:
		if *x > 0 && len(lines) > *y+1 && len(lines[*y+1]) > *x-1 {
			cell := lines[*y+1][*x-1]
			switch cell {
			case '/':
				*x--
				*y++
				nextState = edgeState
			case '+':
				*x--
				*y++
				nextState = cornerState
			case 'v':
				arrowEnd = true
				*x--
				*y++
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case eEdge:
		if *x > 0 && len(lines[*y]) > *x-1 {
			cell := lines[*y][*x-1]
			switch cell {
			case '=':
				dotted = true
				fallthrough
			case '-':
				*x--
				nextState = edgeState
			case '+':
				*x--
				nextState = cornerState
			case '<':
				arrowEnd = true
				*x--
				fallthrough
			default:
				nextState = endState

			}
		} else {
			nextState = endState
		}

	case seEdge:
		if *x > 0 && *y > 0 && len(lines[*y-1]) > *x-1 {
			cell := lines[*y-1][*x-1]
			switch cell {
			case '\\':
				*x--
				*y--
				nextState = edgeState
			case '+':
				*x--
				*y--
				nextState = cornerState
			case '^':
				arrowEnd = true
				*x--
				*y--
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case sEdge:
		if *y > 0 && len(lines[*y-1]) > *x {
			cell := lines[*y-1][*x]
			switch cell {
			case ':':
				dotted = true
				fallthrough
			case '|':
				*y--
				nextState = edgeState
			case '+':
				*y--
				nextState = cornerState
			case '^':
				arrowEnd = true
				*y--
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case swEdge:
		if *y > 0 && len(lines[*y-1]) > *x+1 {
			cell := lines[*y-1][*x+1]
			switch cell {
			case '/':
				*x++
				*y--
				nextState = edgeState
			case '+':
				*x++
				*y--
				nextState = cornerState
			case '^':
				arrowEnd = true
				*x++
				*y--
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case wEdge:
		if len(lines[*y]) > *x+1 {
			cell := lines[*y][*x+1]
			switch cell {
			case '=':
				dotted = true
				fallthrough
			case '-':
				*x++
				nextState = edgeState
			case '+':
				*x++
				nextState = cornerState
			case '>':
				arrowEnd = true
				*x++
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}

	case nwEdge:
		if len(lines) > *y+1 && len(lines[*y+1]) > *x+1 {
			cell := lines[*y+1][*x+1]
			switch cell {
			case '\\':
				*x++
				*y++
				nextState = edgeState
			case '+':
				*x++
				*y++
				nextState = cornerState
			case 'v':
				arrowEnd = true
				*x++
				*y++
				fallthrough
			default:
				nextState = endState
			}
		} else {
			nextState = endState
		}
	}
	return
}

func outgoingEdges(lines [][]byte, x, y int) byte {
	var edges byte
	if y > 0 && len(lines[y-1]) > x && (lines[y-1][x] == '|' || lines[y-1][x] == ':') {
		edges |= nEdge
	}
	if y > 0 && len(lines[y-1]) > x+1 && lines[y-1][x+1] == '/' {
		edges |= neEdge
	}
	if len(lines[y]) > x+1 && (lines[y][x+1] == '-' || lines[y][x+1] == '=') {
		edges |= eEdge
	}
	if len(lines) > y+1 && len(lines[y+1]) > x+1 && lines[y+1][x+1] == '\\' {
		edges |= seEdge
	}
	if len(lines) > y+1 && len(lines[y+1]) > x && (lines[y+1][x] == '|' || lines[y+1][x] == ':') {
		edges |= sEdge
	}
	if x > 0 && len(lines) > y+1 && len(lines[y+1]) > x-1 && lines[y+1][x-1] == '/' {
		edges |= swEdge
	}
	if x > 0 && len(lines[y]) > x-1 && (lines[y][x-1] == '-' || lines[y][x-1] == '=') {
		edges |= wEdge
	}
	if x > 0 && y > 0 && len(lines[y-1]) > x-1 && lines[y-1][x-1] == '\\' {
		edges |= nwEdge
	}
	return edges
}

func (pl *Polyline) scale(p *Parser) {
	for i := 0; i < len(pl.X); i++ {
		pl.X[i] = pl.X[i]*p.xScale + p.xScale/2
		pl.Y[i] = pl.Y[i]*p.yScale + p.yScale/2
	}
}

func (pg *Polygon) scale(p *Parser) {
	for i := 0; i < len(pg.X); i++ {
		pg.X[i] = pg.X[i]*p.xScale + p.xScale/2
		pg.Y[i] = pg.Y[i]*p.yScale + p.yScale/2
	}
}
