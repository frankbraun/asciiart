// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	//"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type vector struct {
	aa  string      // ASCII art
	res interface{} // result: *Grid or *ParseError
}

func testVectors() []vector {
	return []vector{
		{
			aa: `
#?
`,
			res: &ParseError{
				X:   1,
				Y:   1,
				Err: ErrExpRecLine,
			},
		},
		{
			aa: `
#-
`,
			res: &ParseError{
				X:   2,
				Y:   1,
				Err: ErrNoRecUpRightCorn,
			},
		},
		{
			aa: `
#-#
`,
			res: &ParseError{
				X:   2,
				Y:   2,
				Err: ErrExpRecLineOrLowCorn,
			},
		},
		{
			aa: `
#-#
  ?
`,
			res: &ParseError{
				X:   2,
				Y:   2,
				Err: ErrExpRecLineOrLowCorn,
			},
		},
		{
			aa: `
#-#
  |`,
			res: &ParseError{
				X:   2,
				Y:   3,
				Err: ErrNoRecLowRightCorn,
			},
		},
		{
			aa: `
#-#
  |
 ?#
`,
			res: &ParseError{
				X:   1,
				Y:   3,
				Err: ErrExpRecHorizontalLine,
			},
		},
		{
			aa: `
#-#
  |
?-#
`,
			res: &ParseError{
				X:   0,
				Y:   3,
				Err: ErrExpRecLowCorn,
			},
		},
		{
			aa: `
#-#
? |
#-#
`,
			res: &ParseError{
				X:   0,
				Y:   2,
				Err: ErrExpRecVerticalLine,
			},
		},
		{
			aa: `
.?
`,
			res: &ParseError{
				X:   1,
				Y:   1,
				Err: ErrExpRecLine,
			},
		},
		{
			aa: `
#-?
`,
			res: &ParseError{
				X:   2,
				Y:   1,
				Err: ErrExpRecLineOrUpCorn,
			},
		},
		// smallest possible rectangle
		{
			aa: `
#-#
| |
#-#
`,
			res: &Grid{
				W: 3,
				H: 5,
				Elems: []interface{}{
					Rectangle{
						X: 0,
						Y: 1,
						W: 3,
						H: 3,
					},
				},
			},
		},
		{
			aa: `
#-----#
|     |
|     |
#-----#
`,
			res: &Grid{
				W: 7,
				H: 6,
				Elems: []interface{}{
					Rectangle{
						X: 0,
						Y: 1,
						W: 7,
						H: 4,
					},
				},
			},
		},
		{
			aa: `
.-----.
|     |
|     |
'-----'
`,
			res: &Grid{
				W: 7,
				H: 6,
				Elems: []interface{}{
					Rectangle{
						X:               0,
						Y:               1,
						W:               7,
						H:               4,
						RoundUpperLeft:  true,
						RoundUpperRight: true,
						RoundLowerLeft:  true,
						RoundLowerRight: true,
					},
				},
			},
		},
	}
}

func TestParser(t *testing.T) {
	p := NewParser()
	p.SetScale(1, 1)
	for _, vector := range testVectors() {
		g, err := p.Parse(vector.aa)
		if err != nil {
			//fmt.Println(err)
			assert.Equal(t, vector.res, err)
		} else {
			assert.Equal(t, vector.res, g)
		}
	}
}

func TestSetScale(t *testing.T) {
	p := NewParser()
	err := p.SetScale(0, 1)
	assert.Equal(t, ErrWrongXScale, err)
	err = p.SetScale(1, 0)
	assert.Equal(t, ErrWrongYScale, err)
}

func TestParseError(t *testing.T) {
	err := &ParseError{X: 1, Y: 2, Err: ErrExpRecLine}
	assert.Equal(t, "aa2d: expected rectangle line (-) at (1,2)", err.Error())
}
