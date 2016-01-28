// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

// The Rectangle element.
type Rectangle struct {
	X               float64                // x-axis coordinate
	Y               float64                // y-axis coordinate
	W               float64                // width of rectangle
	H               float64                // height of rectangle
	RoundUpperLeft  bool                   // rounded upper-left corner
	RoundUpperRight bool                   // rounded upper-right corner
	RoundLowerLeft  bool                   // rounded lower-left corner
	RoundLowerRight bool                   // rounded lower-right corner
	Ref             map[string]interface{} // JSON reference of the rectangle, if defined
	Elems           []interface{}          // contained elements
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
