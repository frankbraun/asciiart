package asciiart

import (
	"fmt"
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
|     |
|     |
#-----#

[REF]: { "foo": "bar" }
`,
			res: &ParseError{
				X:   0,
				Y:   6,
				Err: ErrRefUnused,
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
		{
			aa: `
#----#
|foo |
#----#
`,
			res: &Grid{
				W:      6,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Rectangle{
						X: 0.5,
						Y: 1.5,
						W: 5,
						H: 2,
						Elems: []interface{}{
							&Textline{
								X:    1.5,
								Y:    2.5,
								Text: "foo",
							},
						},
					},
				},
			},
		},
		{
			aa: `
foo#---#
   |   |
   #---#
`,
			res: &Grid{
				W:      8,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Textline{
						X:    0.5,
						Y:    1.5,
						Text: "foo",
					},
					&Rectangle{
						X: 3.5,
						Y: 1.5,
						W: 4,
						H: 2,
					},
				},
			},
		},
		{
			aa: `--- `,
			res: &Grid{
				W:      4,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Line{
						X1: 0.5,
						Y1: 0.5,
						X2: 3.5,
						Y2: 0.5,
					},
				},
			},
		},
		{
			aa: `-+`,
			res: &ParseError{
				X:   1,
				Y:   0,
				Err: ErrPolyCornerOneEdge,
			},
		},
		{
			aa: `
---+
   |
    
`,
			res: &Grid{
				W:      4,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5},
						Y:      []float64{1.5, 1.5, 2.5},
						Dotted: []bool{false, false},
					},
				},
			},
		},
		{
			aa: `
---+
   |

`,
			res: &Grid{
				W:      4,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5},
						Y:      []float64{1.5, 1.5, 2.5},
						Dotted: []bool{false, false},
					},
				},
			},
		},
		{
			aa: `
---+
   :
`,
			res: &Grid{
				W:      4,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5},
						Y:      []float64{1.5, 1.5, 2.5},
						Dotted: []bool{false, true},
					},
				},
			},
		},
		{
			aa: `
---+
   |
   |
   +=
`,
			res: &Grid{
				W:      5,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5, 4.5},
						Y:      []float64{1.5, 1.5, 4.5, 4.5},
						Dotted: []bool{false, false, true},
					},
				},
			},
		},
		{
			aa: `
---+
   |
   ?
`,
			res: &Grid{
				W:      4,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5},
						Y:      []float64{1.5, 1.5, 2.5},
						Dotted: []bool{false, false},
					},
					&Textline{
						X:    3.5,
						Y:    3.5,
						Text: "?",
					},
				},
			},
		},
		{
			aa: `
---+
   :
   :
   v
`,
			res: &Grid{
				W:      4,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 3.5},
						Y:        []float64{1.5, 1.5, 4.5},
						ArrowEnd: true,
						Dotted:   []bool{false, true},
					},
				},
			},
		},
		{
			aa: `
===+
  /
 /
v
`,
			res: &Grid{
				W:      4,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 0.5},
						Y:        []float64{1.5, 1.5, 4.5},
						ArrowEnd: true,
						Dotted:   []bool{true, false},
					},
				},
			},
		},
		{
			aa: `
===+
  /
 /
+==>
`,
			res: &Grid{
				W:      4,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 0.5, 3.5},
						Y:        []float64{1.5, 1.5, 4.5, 4.5},
						ArrowEnd: true,
						Dotted:   []bool{true, false, true},
					},
				},
			},
		},
		{
			aa: `
---+
  /
 /
`,
			res: &Grid{
				W:      4,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 1.5},
						Y:      []float64{1.5, 1.5, 3.5},
						Dotted: []bool{false, false},
					},
				},
			},
		},
		{
			aa: `
---+
   |
<--+
`,
			res: &Grid{
				W:      4,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 3.5, 0.5},
						Y:        []float64{1.5, 1.5, 3.5, 3.5},
						ArrowEnd: true,
						Dotted:   []bool{false, false, false},
					},
				},
			},
		},
		{
			aa: `
---+
   |
=--+
`,
			res: &Grid{
				W:      4,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 3.5, 3.5, 0.5},
						Y:      []float64{1.5, 1.5, 3.5, 3.5},
						Dotted: []bool{false, false, true},
					},
				},
			},
		},
		{
			aa: `
-+
  \
`,
			res: &Grid{
				W:      3,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 1.5, 2.5},
						Y:      []float64{1.5, 1.5, 2.5},
						Dotted: []bool{false, false},
					},
				},
			},
		},
		{
			aa: `
-+
  \
   +=
`,
			res: &Grid{
				W:      5,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 1.5, 3.5, 4.5},
						Y:      []float64{1.5, 1.5, 3.5, 3.5},
						Dotted: []bool{false, false, true},
					},
				},
			},
		},

		{
			aa: `
<--+
   |
+--+
| ^
|/
+
`,
			res: &Grid{
				W:      4,
				H:      7,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:          []float64{0.5, 3.5, 3.5, 0.5, 0.5, 2.5},
						Y:          []float64{1.5, 1.5, 3.5, 3.5, 6.5, 4.5},
						ArrowStart: true,
						ArrowEnd:   true,
						Dotted:     []bool{false, false, false, false, false},
					},
				},
			},
		},
		{
			aa: `
---+ +->
   | :
+--+ :
|    :
|    :
+----+
`,
			res: &Grid{
				W:      8,
				H:      7,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 3.5, 0.5, 0.5, 5.5, 5.5, 7.5},
						Y:        []float64{1.5, 1.5, 3.5, 3.5, 6.5, 6.5, 1.5, 1.5},
						ArrowEnd: true,
						Dotted:   []bool{false, false, false, false, false, true, false},
					},
				},
			},
		},
		{
			aa: `
---+
   |
+==+
|
|  ^
|   \
+----+
`,
			res: &Grid{
				W:      6,
				H:      8,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 3.5, 0.5, 0.5, 5.5, 3.5},
						Y:        []float64{1.5, 1.5, 3.5, 3.5, 7.5, 7.5, 5.5},
						ArrowEnd: true,
						Dotted:   []bool{false, false, true, false, false, false},
					},
				},
			},
		},
		{
			aa: `<--+ /
   |/
   +`,
			res: &Grid{
				W:      6,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:          []float64{0.5, 3.5, 3.5, 5.5},
						Y:          []float64{0.5, 0.5, 2.5, 0.5},
						ArrowStart: true,
						Dotted:     []bool{false, false, false},
					},
				},
			},
		},
		{
			aa: `-+ |
 | |
 +-+`,
			res: &Grid{
				W:      4,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 1.5, 1.5, 3.5, 3.5},
						Y:      []float64{0.5, 0.5, 2.5, 2.5, 0.5},
						Dotted: []bool{false, false, false, false},
					},
				},
			},
		},

		{
			aa: `-+ \
 |  \
 +---+`,
			res: &Grid{
				W:      6,
				H:      3,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:      []float64{0.5, 1.5, 1.5, 5.5, 3.5},
						Y:      []float64{0.5, 0.5, 2.5, 2.5, 0.5},
						Dotted: []bool{false, false, false, false},
					},
				},
			},
		},

		{
			aa: `
---+
   :
+--+
 \
  \
   v
`,
			res: &Grid{
				W:      4,
				H:      7,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 3.5, 3.5, 0.5, 3.5},
						Y:        []float64{1.5, 1.5, 3.5, 3.5, 6.5},
						ArrowEnd: true,
						Dotted:   []bool{false, true, false, false},
					},
				},
			},
		},
		{
			aa: `
----------+
          :
+--------=+
:
:       ^ 
:       |
:       |
:       +
:      /
:     /
:    + 
:     \ 
:      \ 
+-------+  
`,
			res: &Grid{
				W:      11,
				H:      15,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polyline{
						X:        []float64{0.5, 10.5, 10.5, 0.5, 0.5, 8.5, 5.5, 8.5, 8.5},
						Y:        []float64{1.5, 1.5, 3.5, 3.5, 14.5, 14.5, 11.5, 8.5, 5.5},
						ArrowEnd: true,
						Dotted:   []bool{false, true, true, true, false, false, false, false},
					},
				},
			},
		},
		{
			aa: `
+--+
|  |
+--+
		`,
			res: &Grid{
				W:      4,
				H:      5,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polygon{
						X:      []float64{0.5, 3.5, 3.5, 0.5},
						Y:      []float64{1.5, 1.5, 3.5, 3.5},
						Dotted: []bool{false, false, false, false},
					},
				},
			},
		},
		{
			aa: `
+
|\ 
+-+
`,
			res: &Grid{
				W:      3,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polygon{
						X:      []float64{0.5, 2.5, 0.5},
						Y:      []float64{1.5, 3.5, 3.5},
						Dotted: []bool{false, false, false},
					},
				},
			},
		},
		{
			aa: `
  +
 /|
+-+
`,
			res: &Grid{
				W:      3,
				H:      4,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Polygon{
						X:      []float64{2.5, 2.5, 0.5},
						Y:      []float64{1.5, 3.5, 3.5},
						Dotted: []bool{false, false, false},
					},
				},
			},
		},
		{
			aa: `<T`,
			res: &Grid{
				W:      2,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Textline{
						X:    0.5,
						Y:    0.5,
						Text: "<T",
					},
				},
			},
		},
		{
			aa: `<`,
			res: &Grid{
				W:      1,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Textline{
						X:    0.5,
						Y:    0.5,
						Text: "<",
					},
				},
			},
		},
		{
			aa: `^`,
			res: &Grid{
				W:      1,
				H:      1,
				XScale: 1,
				YScale: 1,
				Elems: []interface{}{
					&Textline{
						X:    0.5,
						Y:    0.5,
						Text: "^",
					},
				},
			},
		},
		{
			aa: `
+--+
|  |
   |
+--+
`,
			res: &ParseError{
				X:   0,
				Y:   1,
				Err: ErrPolygonNotClosed,
			},
		},
		{
			aa: `+`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrPolyCornerNoEdge,
			},
		},
		{
			aa: `
+-
|\
`,
			res: &ParseError{
				X:   0,
				Y:   1,
				Err: ErrPolyCornerTooManyEdges,
			},
		},
		{
			aa: `+-`,
			res: &ParseError{
				X:   0,
				Y:   0,
				Err: ErrPolyCornerOneEdge,
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
			fmt.Println(err)
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
