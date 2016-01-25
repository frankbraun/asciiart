// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
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
				Err: ErrExpRecLineOrCorn,
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
