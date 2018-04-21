package asciiart

// The Textline element.
type Textline struct {
	X    float64 // x-axis coordinate of the start of the text
	Y    float64 // y-axis coordinate of the start of the text
	Text string  // the actual text string
}

func (p *Parser) parseTextline(
	parent elem,
	lines [][]byte,
	startX, startY int,
) error {
	var t Textline
	t.X = float64(startX)
	t.Y = float64(startY)
	var space bool
	x := startX
forLoop:
	for ; x < len(lines[startY]); x++ {
		cell := lines[startY][x]
		switch cell {
		case '#', '|':
			space = false
			break forLoop
		case ' ', '\t', '\r':
			if space {
				break forLoop
			}
			space = true
		default:
			space = false
		}
	}
	if space {
		x--
	}
	t.Text = string(lines[startY][startX:x])
	for i := startX; i < x; i++ {
		lines[startY][i] = ' ' // nom nom nom
	}
	// scale
	t.X = t.X*p.xScale + p.xScale/2
	t.Y = t.Y*p.yScale + p.yScale/2
	// add line to parent
	parent.addElem(&t)
	return nil
}
