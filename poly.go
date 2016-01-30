// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

import (
	"errors"
	"fmt"

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

// parsePoly parses a polyline or polygon.
//
//  \|/
//  -+-
//  /|\
//
func (p *Parser) parsePoly(
	parent elem,
	lines [][]byte,
	l *Line,
) error {
	fmt.Println("parsePoly")
	var pl Polyline
	// take stuff over from parsed line
	pl.X = append(pl.X, l.X1)
	pl.Y = append(pl.Y, l.Y1)
	pl.ArrowStart = l.ArrowStart

	x := int(l.X2)
	y := int(l.Y2)
	dotted := l.Dotted
	state := cornerState
	var edge byte
	for {
		switch state {
		case cornerState:
			fmt.Println("cornerState", x, y)
			lines[y][x] = ' ' // nom nom nom
			// add segement
			pl.X = append(pl.X, float64(x))
			pl.Y = append(pl.Y, float64(y))
			pl.Dotted = append(pl.Dotted, dotted)
			dotted = false
			// process corner (+)
			outEdges := outgoingEdges(lines, x, y)
			switch bit.Count(outEdges) {
			case 0:
				return &ParseError{X: x, Y: y, Err: ErrPolyCornerOneEdge}
			case 1:
				fmt.Println("follow edge", x, y)
				// follow edge
				switch {
				case outEdges&nEdge > 1:
					edge = sEdge
					y--
				case outEdges&neEdge > 1:
					edge = swEdge
					x++
					y--
				case outEdges&eEdge > 1:
					edge = wEdge
					x++
				case outEdges&seEdge > 1:
					edge = nwEdge
					x++
					y++
				case outEdges&sEdge > 1:
					edge = nEdge
					y++
				case outEdges&swEdge > 1:
					edge = neEdge
					x--
					y++
				case outEdges&wEdge > 1:
					edge = eEdge
					x--
				case outEdges&nwEdge > 1:
					edge = seEdge
					x--
					y--
				}
				fmt.Println("call procEdge", x, y)
				next, d := pl.procEdge(lines, &x, &y, edge)
				if d {
					dotted = true
				}
				state = next
			default:
				// TODO: split
				return errors.New("poly split not implemented")
			}
		case edgeState:
			next, d := pl.procEdge(lines, &x, &y, edge)
			if d {
				dotted = true
			}
			state = next
		case endState:
			lines[y][x] = ' ' // nom nom nom
			// add final segment
			pl.X = append(pl.X, float64(x))
			pl.Y = append(pl.Y, float64(y))
			pl.Dotted = append(pl.Dotted, dotted)
			// scale
			pl.scale(p)
			// add polyline to parent
			parent.addElem(&pl)
			return nil
		}
	}

	return nil
}

// procEdge processes an edge at position x,y and returns the next state.
// Three possible outcomes:
//   1. we can continue the line:
//      - nextState = edgeState
//      - x and y changed accordingly
//      - cell eaten
//   2. we hit a corner:
//      - nextState = cornerState
//      - x and y unchanged
//      - cell unmodified
//   3. we reach the corner of the grid or cannot continue the line:
//      - nextState = endState,
//      - x and y unchanged
//      - cell unmodified
func (pl *Polyline) procEdge(
	lines [][]byte,
	x, y *int,
	incomingEdge byte,
) (nextState int, dotted bool) {
	fmt.Println("procEdge", *x, *y)
	cell := lines[*y][*x]

	switch incomingEdge {

	case nEdge:
		switch cell {
		case ':':
			dotted = true
			fallthrough
		case '|':
			if len(lines) > *y+1 && len(lines[*y+1]) > *x {
				lines[*y][*x] = ' ' // nom nom nom
				*y++
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case neEdge:
		switch cell {
		case '/':
			if *x > 0 && len(lines) > *y+1 && len(lines[*y+1]) > *x-1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x--
				*y++
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case eEdge:
		switch cell {
		case '=':
			dotted = true
			fallthrough
		case '-':
			if *x > 0 && len(lines[*y]) > *x-1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x--
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case seEdge:
		switch cell {
		case '\\':
			if *x > 0 && *y > 0 && len(lines[*y-1]) > *x-1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x--
				*y--
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case sEdge:
		switch cell {
		case ':':
			dotted = true
			fallthrough
		case '|':
			if *y > 0 && len(lines[*y-1]) > *x {
				lines[*y][*x] = ' ' // nom nom nom
				*y--
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case swEdge:
		switch cell {
		case '/':
			if *y > 0 && len(lines[*y-1]) > *x+1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x++
				*y--
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case wEdge:
		switch cell {
		case '=':
			dotted = true
			fallthrough
		case '-':
			if len(lines[*y]) > *x+1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x++
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}

	case nwEdge:
		switch cell {
		case '\\':
			if len(lines) > *y+1 && len(lines[*y+1]) > *x+1 {
				lines[*y][*x] = ' ' // nom nom nom
				*x++
				*y++
				nextState = edgeState
			} else {
				nextState = endState
			}
		case '+':
			nextState = cornerState
		default:
			nextState = endState
		}
	}

	return
}

func outgoingEdges(lines [][]byte, x, y int) byte {
	fmt.Println("outgoingEdges")
	var edges byte
	if y > 0 && len(lines[y-1]) > x && (lines[y-1][x] == '|' || lines[y-1][x] == ':') {
		fmt.Println("north")
		edges |= nEdge
	}
	if y > 0 && len(lines[y-1]) > x+1 && lines[y-1][x+1] == '/' {
		fmt.Println("northeast")
		edges |= neEdge
	}
	if len(lines[y]) > x+1 && (lines[y][x+1] == '-' || lines[y][x+1] == '=') {
		fmt.Println("east")
		edges |= eEdge
	}
	if len(lines) > y+1 && len(lines[y+1]) > x+1 && lines[y+1][x+1] == '\\' {
		fmt.Println("southeast")
		edges |= swEdge
	}
	if len(lines) > y+1 && len(lines[y+1]) > x && (lines[y+1][x] == '|' || lines[y+1][x] == ':') {
		fmt.Println("south")
		edges |= sEdge
	}
	if x > 0 && len(lines) > y+1 && len(lines[y+1]) > x-1 && lines[y+1][x-1] == '/' {
		fmt.Println("southwest")
		edges |= swEdge
	}
	if x > 0 && len(lines[y]) > x-1 && (lines[y][x-1] == '-' || lines[y][x-1] == '=') {
		fmt.Println("west")
		edges |= wEdge
	}
	if x > 0 && y > 0 && len(lines[y-1]) > x-1 && lines[y-1][x-1] == '\\' {
		fmt.Println("northwest")
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
