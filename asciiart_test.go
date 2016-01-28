// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

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
?
`,
			res: &Grid{
				W:      1,
				H:      2,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Textline{
						X:    0.5,
						Y:    1.5,
						Text: "?",
					},
				},
			},
		},
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
				W:      3,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 2,
						H: 2,
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
				W:      7,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 6,
						H: 3,
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
				W:      7,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X:               0.5,
						Y:               1.5,
						W:               6,
						H:               3,
						RoundUpperLeft:  true,
						RoundUpperRight: true,
						RoundLowerLeft:  true,
						RoundLowerRight: true,
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

[_SVG]: {"blur": false}
`,
			res: &Grid{
				W:      7,
				H:      5,
				XScale: 1,
				YScale: 1,
				Refs: map[string]map[string]interface{}{
					"_SVG": {
						"blur": false,
					},
				},
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 6,
						H: 3,
					},
				},
			},
		},
		{
			aa: `
#-----#
|#---#|
||   ||
|#---#|
#-----#
`,
			res: &Grid{
				W:      7,
				H:      6,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 6,
						H: 4,
						Elems: []interface{}{
							&Rectangle{
								X: 1.5,
								Y: 2.5,
								W: 4,
								H: 2,
							},
						},
					},
				},
			},
		},
		{
			aa: `--`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1: 0.5,
						Y1: 0.5,
						X2: 1.5,
						Y2: 0.5,
					},
				},
			},
		},
		{
			aa: `==`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:     0.5,
						Y1:     0.5,
						X2:     1.5,
						Y2:     0.5,
						Dotted: true,
					},
				},
			},
		},
		{
			aa: `=-`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:     0.5,
						Y1:     0.5,
						X2:     1.5,
						Y2:     0.5,
						Dotted: true,
					},
				},
			},
		},
		{
			aa: `<-`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         0.5,
						X2:         1.5,
						Y2:         0.5,
						ArrowStart: true,
					},
				},
			},
		},
		{
			aa: `<=`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         0.5,
						X2:         1.5,
						Y2:         0.5,
						ArrowStart: true,
						Dotted:     true,
					},
				},
			},
		},
		{
			aa: `->`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:       0.5,
						Y1:       0.5,
						X2:       1.5,
						Y2:       0.5,
						ArrowEnd: true,
					},
				},
			},
		},
		{
			aa: `<=>`,
			res: &Grid{
				W:      3,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         0.5,
						X2:         2.5,
						Y2:         0.5,
						ArrowStart: true,
						ArrowEnd:   true,
						Dotted:     true,
					},
				},
			},
		},
		{
			aa: `
|
|
`,
			res: &Grid{
				W:      1,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1: 0.5,
						Y1: 1.5,
						X2: 0.5,
						Y2: 2.5,
					},
				},
			},
		},
		{
			aa: `
:
:
`,
			res: &Grid{
				W:      1,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:     0.5,
						Y1:     1.5,
						X2:     0.5,
						Y2:     2.5,
						Dotted: true,
					},
				},
			},
		},
		{
			aa: `
^
|
`,
			res: &Grid{
				W:      1,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         1.5,
						X2:         0.5,
						Y2:         2.5,
						ArrowStart: true,
					},
				},
			},
		},
		{
			aa: `
^
:
`,
			res: &Grid{
				W:      1,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         1.5,
						X2:         0.5,
						Y2:         2.5,
						ArrowStart: true,
						Dotted:     true,
					},
				},
			},
		},
		{
			aa: `
 /
/
`,
			res: &Grid{
				W:      2,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1: 1.5,
						Y1: 1.5,
						X2: 0.5,
						Y2: 2.5,
					},
				},
			},
		},
		{
			aa: `
 ^
/
`,
			res: &Grid{
				W:      2,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         1.5,
						Y1:         1.5,
						X2:         0.5,
						Y2:         2.5,
						ArrowStart: true,
					},
				},
			},
		},
		{
			aa: `
\
 \
`,
			res: &Grid{
				W:      2,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1: 0.5,
						Y1: 1.5,
						X2: 1.5,
						Y2: 2.5,
					},
				},
			},
		},
		{
			aa: `
^
 \
`,
			res: &Grid{
				W:      2,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1:         0.5,
						Y1:         1.5,
						X2:         1.5,
						Y2:         2.5,
						ArrowStart: true,
					},
				},
			},
		},
		{
			aa: `>`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrRightArrow,
			},
		},
		{
			aa: `-<`,
			res: &ParseError{
				X:   1,
				Y:   0,
				Err: ErrLineLeftArrow,
			},
		},
		{
			aa: `-`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrLineTooShort,
			},
		},
		{
			aa: `[]`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrRefParseError,
			},
		},
		{
			aa: `[REF]: false`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrRefJSONObj,
			},
		},
		{
			aa: `
[REF]: {"foo": "bar"}
[REF]: {"foo": "baz"}
`,
			res: &ParseError{
				X:   0,
				Y:   2,
				Err: ErrRefTwice,
			},
		},
		{
			aa: `
[_REF]: {"foo": "bar"}
[_REF]: {"foo": "baz"}
`,
			res: &ParseError{
				X:   0,
				Y:   2,
				Err: ErrRefTwice,
			},
		},
		{
			aa: `[REF]: {`, // ' '
			res: &ParseError{
				X:   7,
				Y:   0,
				Err: ErrRefJSONUnmarshal,
			},
		},
		{
			aa: `[REF]:	{`, // '\t'
			res: &ParseError{
				X:   7,
				Y:   0,
				Err: ErrRefJSONUnmarshal,
			},
		},
		{
			aa: `
#-----#
|[    |
|     |
#-----#
`,
			res: &ParseError{
				X:   6,
				Y:   2,
				Err: ErrRefMissingBracket,
			},
		},
		{
			aa: `
#-----#
|[]   |
|     |
#-----#
`,
			res: &ParseError{
				X:   2,
				Y:   2,
				Err: ErrRefKeyEmpty,
			},
		},
		{
			aa: `
#-----#
|[UN] |
|     |
#-----#
`,
			res: &ParseError{
				X:   2,
				Y:   2,
				Err: ErrRefKeyUndefined,
			},
		},
		{
			aa: `
#-----#
|[REF]|
|     |
#-----#

[REF]: { "foo": "bar" }
`,
			res: &Grid{
				W:      7,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 6,
						H: 3,
						Ref: map[string]interface{}{
							"foo": "bar",
						},
					},
				},
			},
		},
		/* TODO: fix test
		   		{
		   			aa: `
		   [REF]: {}
		   --
		   `,
		   			res: &ParseError{
		   				X:   7,
		   				Y:   0,
		   				Err: ErrRefMiddle,
		   			},
		   		},
		*/
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
	assert.Equal(t, "asciiart: expected rectangle line (-) at (1,2)", err.Error())
}
