// This code is generated from grammar.peg, do not edit!

package path

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var cmds = map[string]string{
	"m": "moveto",
	"l": "lineto",
	"h": "horizontal lineto",
	"v": "vertical lineto",
	"c": "curveto",
	"s": "smooth curveto",
	"q": "quadratic curveto",
	"t": "smooth quadratic curveto",
	"a": "elliptical arc",
	"z": "closepath",
}

func init() {
	var keys []string
	for k := range cmds {
		keys = append(keys, k)
	}
	for _, k := range keys {
		cmds[strings.ToUpper(k)] = cmds[k]
	}
}

func merge(first interface{}, more []interface{}) ([]interface{}, error) {
	/*
	   fmt.Println("merge")
	   jsn, err := json.MarshalIndent(first, "", "  ")
	   if err != nil {
	     return nil, err
	   }
	   fmt.Println("JSON:", string(jsn))
	   jsn, err = json.MarshalIndent(more, "", "  ")
	   if err != nil {
	     return nil, err
	   }
	   fmt.Println("JSON:", string(jsn))
	*/

	res := first.([]interface{})
	if len(more) > 0 {
		res = append(res, more...)
	}
	return res, nil
}

func commands(c string, args []interface{}) ([]interface{}, error) {
	var res []interface{}
	if len(args) == 0 {
		cmd := map[string]interface{}{
			"code":    c,
			"command": cmds[c],
		}
		res = append(res, cmd)
	} else {
		for _, arg := range args {
			cmd := map[string]interface{}{
				"code":    c,
				"command": cmds[c],
			}
			if c == strings.ToLower(c) {
				cmd["relative"] = true
			}
			m := arg.(map[string]interface{})
			for k, v := range m {
				cmd[k] = v
			}
			res = append(res, cmd)
		}
	}
	return res, nil
}

var g = &grammar{
	rules: []*rule{
		{
			name: "svg_path",
			pos:  position{line: 81, col: 1, offset: 1594},
			expr: &actionExpr{
				pos: position{line: 81, col: 12, offset: 1605},
				run: (*parser).callonsvg_path1,
				expr: &seqExpr{
					pos: position{line: 81, col: 12, offset: 1605},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 12, offset: 1605},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 12, offset: 1605},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 17, offset: 1610},
							label: "data",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 22, offset: 1615},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 22, offset: 1615},
									name: "moveTo_drawTo_commandGroups",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 51, offset: 1644},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 51, offset: 1644},
								name: "wsp",
							},
						},
					},
				},
			},
		},
		{
			name: "moveTo_drawTo_commandGroups",
			pos:  position{line: 98, col: 1, offset: 2016},
			expr: &actionExpr{
				pos: position{line: 98, col: 31, offset: 2046},
				run: (*parser).callonmoveTo_drawTo_commandGroups1,
				expr: &seqExpr{
					pos: position{line: 98, col: 31, offset: 2046},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 98, col: 31, offset: 2046},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 37, offset: 2052},
								name: "moveTo_drawTo_commandGroup",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 64, offset: 2079},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 69, offset: 2084},
								expr: &seqExpr{
									pos: position{line: 98, col: 70, offset: 2085},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 98, col: 70, offset: 2085},
											expr: &ruleRefExpr{
												pos:  position{line: 98, col: 70, offset: 2085},
												name: "wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 98, col: 75, offset: 2090},
											name: "moveTo_drawTo_commandGroup",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "moveTo_drawTo_commandGroup",
			pos:  position{line: 102, col: 1, offset: 2168},
			expr: &actionExpr{
				pos: position{line: 102, col: 30, offset: 2197},
				run: (*parser).callonmoveTo_drawTo_commandGroup1,
				expr: &seqExpr{
					pos: position{line: 102, col: 30, offset: 2197},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 102, col: 30, offset: 2197},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 36, offset: 2203},
								name: "moveto",
							},
						},
						&labeledExpr{
							pos:   position{line: 102, col: 43, offset: 2210},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 102, col: 48, offset: 2215},
								expr: &seqExpr{
									pos: position{line: 102, col: 49, offset: 2216},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 102, col: 49, offset: 2216},
											expr: &ruleRefExpr{
												pos:  position{line: 102, col: 49, offset: 2216},
												name: "wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 102, col: 54, offset: 2221},
											name: "drawto_command",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "drawto_command",
			pos:  position{line: 106, col: 1, offset: 2287},
			expr: &choiceExpr{
				pos: position{line: 106, col: 18, offset: 2304},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 106, col: 18, offset: 2304},
						name: "closepath",
					},
					&ruleRefExpr{
						pos:  position{line: 107, col: 18, offset: 2331},
						name: "lineto",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 18, offset: 2355},
						name: "horizontal_lineto",
					},
					&ruleRefExpr{
						pos:  position{line: 109, col: 18, offset: 2390},
						name: "vertical_lineto",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 18, offset: 2423},
						name: "curveto",
					},
					&ruleRefExpr{
						pos:  position{line: 111, col: 18, offset: 2448},
						name: "smooth_curveto",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 18, offset: 2480},
						name: "quadratic_bezier_curveto",
					},
					&ruleRefExpr{
						pos:  position{line: 113, col: 18, offset: 2522},
						name: "smooth_quadratic_bezier_curveto",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 18, offset: 2571},
						name: "elliptical_arc",
					},
				},
			},
		},
		{
			name: "moveto",
			pos:  position{line: 116, col: 1, offset: 2587},
			expr: &actionExpr{
				pos: position{line: 116, col: 10, offset: 2596},
				run: (*parser).callonmoveto1,
				expr: &seqExpr{
					pos: position{line: 116, col: 10, offset: 2596},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 116, col: 10, offset: 2596},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 116, col: 13, offset: 2599},
								val:        "[Mm]",
								chars:      []rune{'M', 'm'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 116, col: 18, offset: 2604},
							expr: &ruleRefExpr{
								pos:  position{line: 116, col: 18, offset: 2604},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 116, col: 23, offset: 2609},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 116, col: 29, offset: 2615},
								name: "coordinate_pair",
							},
						},
						&labeledExpr{
							pos:   position{line: 116, col: 45, offset: 2631},
							label: "more",
							expr: &zeroOrOneExpr{
								pos: position{line: 116, col: 50, offset: 2636},
								expr: &seqExpr{
									pos: position{line: 116, col: 51, offset: 2637},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 116, col: 51, offset: 2637},
											expr: &ruleRefExpr{
												pos:  position{line: 116, col: 51, offset: 2637},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 116, col: 62, offset: 2648},
											name: "lineto_argument_sequence",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "closepath",
			pos:  position{line: 141, col: 1, offset: 3173},
			expr: &actionExpr{
				pos: position{line: 141, col: 13, offset: 3185},
				run: (*parser).callonclosepath1,
				expr: &charClassMatcher{
					pos:        position{line: 141, col: 13, offset: 3185},
					val:        "[Zz]",
					chars:      []rune{'Z', 'z'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "lineto",
			pos:  position{line: 145, col: 1, offset: 3223},
			expr: &actionExpr{
				pos: position{line: 145, col: 10, offset: 3232},
				run: (*parser).callonlineto1,
				expr: &seqExpr{
					pos: position{line: 145, col: 10, offset: 3232},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 145, col: 10, offset: 3232},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 145, col: 13, offset: 3235},
								val:        "[Ll]",
								chars:      []rune{'L', 'l'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 145, col: 18, offset: 3240},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 18, offset: 3240},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 23, offset: 3245},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 28, offset: 3250},
								name: "lineto_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "lineto_argument_sequence",
			pos:  position{line: 149, col: 1, offset: 3341},
			expr: &actionExpr{
				pos: position{line: 149, col: 28, offset: 3368},
				run: (*parser).callonlineto_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 149, col: 28, offset: 3368},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 149, col: 28, offset: 3368},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 34, offset: 3374},
								name: "coordinate_pair",
							},
						},
						&labeledExpr{
							pos:   position{line: 149, col: 50, offset: 3390},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 149, col: 55, offset: 3395},
								expr: &seqExpr{
									pos: position{line: 149, col: 56, offset: 3396},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 149, col: 56, offset: 3396},
											expr: &ruleRefExpr{
												pos:  position{line: 149, col: 56, offset: 3396},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 149, col: 67, offset: 3407},
											name: "coordinate_pair",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "horizontal_lineto",
			pos:  position{line: 153, col: 1, offset: 3488},
			expr: &actionExpr{
				pos: position{line: 153, col: 21, offset: 3508},
				run: (*parser).callonhorizontal_lineto1,
				expr: &seqExpr{
					pos: position{line: 153, col: 21, offset: 3508},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 153, col: 21, offset: 3508},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 153, col: 24, offset: 3511},
								val:        "[Hh]",
								chars:      []rune{'H', 'h'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 153, col: 29, offset: 3516},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 29, offset: 3516},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 153, col: 34, offset: 3521},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 39, offset: 3526},
								name: "coordinate_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "coordinate_sequence",
			pos:  position{line: 161, col: 1, offset: 3729},
			expr: &actionExpr{
				pos: position{line: 161, col: 23, offset: 3751},
				run: (*parser).calloncoordinate_sequence1,
				expr: &seqExpr{
					pos: position{line: 161, col: 23, offset: 3751},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 161, col: 23, offset: 3751},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 29, offset: 3757},
								name: "number",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 36, offset: 3764},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 161, col: 41, offset: 3769},
								expr: &seqExpr{
									pos: position{line: 161, col: 42, offset: 3770},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 161, col: 42, offset: 3770},
											expr: &ruleRefExpr{
												pos:  position{line: 161, col: 42, offset: 3770},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 161, col: 53, offset: 3781},
											name: "number",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "vertical_lineto",
			pos:  position{line: 165, col: 1, offset: 3839},
			expr: &actionExpr{
				pos: position{line: 165, col: 19, offset: 3857},
				run: (*parser).callonvertical_lineto1,
				expr: &seqExpr{
					pos: position{line: 165, col: 19, offset: 3857},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 19, offset: 3857},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 165, col: 22, offset: 3860},
								val:        "[Vv]",
								chars:      []rune{'V', 'v'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 165, col: 27, offset: 3865},
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 27, offset: 3865},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 32, offset: 3870},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 37, offset: 3875},
								name: "coordinate_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "curveto",
			pos:  position{line: 173, col: 1, offset: 4072},
			expr: &actionExpr{
				pos: position{line: 173, col: 11, offset: 4082},
				run: (*parser).calloncurveto1,
				expr: &seqExpr{
					pos: position{line: 173, col: 11, offset: 4082},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 173, col: 11, offset: 4082},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 173, col: 14, offset: 4085},
								val:        "[Cc]",
								chars:      []rune{'C', 'c'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 173, col: 19, offset: 4090},
							expr: &ruleRefExpr{
								pos:  position{line: 173, col: 19, offset: 4090},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 173, col: 24, offset: 4095},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 173, col: 29, offset: 4100},
								name: "curveto_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "curveto_argument_sequence",
			pos:  position{line: 177, col: 1, offset: 4184},
			expr: &actionExpr{
				pos: position{line: 177, col: 29, offset: 4212},
				run: (*parser).calloncurveto_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 177, col: 29, offset: 4212},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 177, col: 29, offset: 4212},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 177, col: 35, offset: 4218},
								name: "curveto_argument",
							},
						},
						&labeledExpr{
							pos:   position{line: 177, col: 52, offset: 4235},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 177, col: 57, offset: 4240},
								expr: &seqExpr{
									pos: position{line: 177, col: 58, offset: 4241},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 177, col: 58, offset: 4241},
											expr: &ruleRefExpr{
												pos:  position{line: 177, col: 58, offset: 4241},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 177, col: 69, offset: 4252},
											name: "curveto_argument",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "curveto_argument",
			pos:  position{line: 181, col: 1, offset: 4320},
			expr: &actionExpr{
				pos: position{line: 181, col: 20, offset: 4339},
				run: (*parser).calloncurveto_argument1,
				expr: &seqExpr{
					pos: position{line: 181, col: 20, offset: 4339},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 181, col: 20, offset: 4339},
							label: "a",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 22, offset: 4341},
								name: "coordinate_pair",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 181, col: 38, offset: 4357},
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 38, offset: 4357},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 181, col: 49, offset: 4368},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 51, offset: 4370},
								name: "coordinate_pair",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 181, col: 67, offset: 4386},
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 67, offset: 4386},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 181, col: 78, offset: 4397},
							label: "cc",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 81, offset: 4400},
								name: "coordinate_pair",
							},
						},
					},
				},
			},
		},
		{
			name: "smooth_curveto",
			pos:  position{line: 196, col: 1, offset: 4711},
			expr: &actionExpr{
				pos: position{line: 196, col: 18, offset: 4728},
				run: (*parser).callonsmooth_curveto1,
				expr: &seqExpr{
					pos: position{line: 196, col: 18, offset: 4728},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 196, col: 18, offset: 4728},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 196, col: 21, offset: 4731},
								val:        "[Ss]",
								chars:      []rune{'S', 's'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 196, col: 26, offset: 4736},
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 26, offset: 4736},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 31, offset: 4741},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 36, offset: 4746},
								name: "smooth_curveto_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "smooth_curveto_argument_sequence",
			pos:  position{line: 200, col: 1, offset: 4837},
			expr: &actionExpr{
				pos: position{line: 200, col: 36, offset: 4872},
				run: (*parser).callonsmooth_curveto_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 200, col: 36, offset: 4872},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 200, col: 36, offset: 4872},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 200, col: 42, offset: 4878},
								name: "smooth_curveto_argument",
							},
						},
						&labeledExpr{
							pos:   position{line: 200, col: 66, offset: 4902},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 200, col: 71, offset: 4907},
								expr: &seqExpr{
									pos: position{line: 200, col: 72, offset: 4908},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 200, col: 72, offset: 4908},
											expr: &ruleRefExpr{
												pos:  position{line: 200, col: 72, offset: 4908},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 200, col: 83, offset: 4919},
											name: "smooth_curveto_argument",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "smooth_curveto_argument",
			pos:  position{line: 204, col: 1, offset: 4994},
			expr: &actionExpr{
				pos: position{line: 204, col: 27, offset: 5020},
				run: (*parser).callonsmooth_curveto_argument1,
				expr: &seqExpr{
					pos: position{line: 204, col: 27, offset: 5020},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 204, col: 27, offset: 5020},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 29, offset: 5022},
								name: "coordinate_pair",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 204, col: 45, offset: 5038},
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 45, offset: 5038},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 204, col: 56, offset: 5049},
							label: "cp",
							expr: &ruleRefExpr{
								pos:  position{line: 204, col: 59, offset: 5052},
								name: "coordinate_pair",
							},
						},
					},
				},
			},
		},
		{
			name: "quadratic_bezier_curveto",
			pos:  position{line: 216, col: 1, offset: 5284},
			expr: &actionExpr{
				pos: position{line: 216, col: 28, offset: 5311},
				run: (*parser).callonquadratic_bezier_curveto1,
				expr: &seqExpr{
					pos: position{line: 216, col: 28, offset: 5311},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 216, col: 28, offset: 5311},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 216, col: 31, offset: 5314},
								val:        "[Qq]",
								chars:      []rune{'Q', 'q'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 216, col: 36, offset: 5319},
							expr: &ruleRefExpr{
								pos:  position{line: 216, col: 36, offset: 5319},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 216, col: 41, offset: 5324},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 216, col: 46, offset: 5329},
								name: "quadratic_bezier_curveto_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "quadratic_bezier_curveto_argument_sequence",
			pos:  position{line: 220, col: 1, offset: 5430},
			expr: &actionExpr{
				pos: position{line: 220, col: 46, offset: 5475},
				run: (*parser).callonquadratic_bezier_curveto_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 220, col: 46, offset: 5475},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 220, col: 46, offset: 5475},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 220, col: 52, offset: 5481},
								name: "quadratic_bezier_curveto_argument",
							},
						},
						&labeledExpr{
							pos:   position{line: 220, col: 86, offset: 5515},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 220, col: 91, offset: 5520},
								expr: &seqExpr{
									pos: position{line: 220, col: 92, offset: 5521},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 220, col: 92, offset: 5521},
											expr: &ruleRefExpr{
												pos:  position{line: 220, col: 92, offset: 5521},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 220, col: 103, offset: 5532},
											name: "quadratic_bezier_curveto_argument",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "quadratic_bezier_curveto_argument",
			pos:  position{line: 224, col: 1, offset: 5617},
			expr: &actionExpr{
				pos: position{line: 224, col: 37, offset: 5653},
				run: (*parser).callonquadratic_bezier_curveto_argument1,
				expr: &seqExpr{
					pos: position{line: 224, col: 37, offset: 5653},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 224, col: 37, offset: 5653},
							label: "a",
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 39, offset: 5655},
								name: "coordinate_pair",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 224, col: 55, offset: 5671},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 55, offset: 5671},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 224, col: 66, offset: 5682},
							label: "b",
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 68, offset: 5684},
								name: "coordinate_pair",
							},
						},
					},
				},
			},
		},
		{
			name: "smooth_quadratic_bezier_curveto",
			pos:  position{line: 236, col: 1, offset: 5912},
			expr: &actionExpr{
				pos: position{line: 236, col: 35, offset: 5946},
				run: (*parser).callonsmooth_quadratic_bezier_curveto1,
				expr: &seqExpr{
					pos: position{line: 236, col: 35, offset: 5946},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 236, col: 35, offset: 5946},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 236, col: 38, offset: 5949},
								val:        "[Tt]",
								chars:      []rune{'T', 't'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 236, col: 43, offset: 5954},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 43, offset: 5954},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 48, offset: 5959},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 53, offset: 5964},
								name: "smooth_quadratic_bezier_curveto_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "smooth_quadratic_bezier_curveto_argument_sequence",
			pos:  position{line: 240, col: 1, offset: 6072},
			expr: &actionExpr{
				pos: position{line: 240, col: 53, offset: 6124},
				run: (*parser).callonsmooth_quadratic_bezier_curveto_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 240, col: 53, offset: 6124},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 240, col: 53, offset: 6124},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 59, offset: 6130},
								name: "coordinate_pair",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 75, offset: 6146},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 240, col: 80, offset: 6151},
								expr: &seqExpr{
									pos: position{line: 240, col: 81, offset: 6152},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 240, col: 81, offset: 6152},
											expr: &ruleRefExpr{
												pos:  position{line: 240, col: 81, offset: 6152},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 240, col: 92, offset: 6163},
											name: "coordinate_pair",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "elliptical_arc",
			pos:  position{line: 244, col: 1, offset: 6230},
			expr: &actionExpr{
				pos: position{line: 244, col: 18, offset: 6247},
				run: (*parser).callonelliptical_arc1,
				expr: &seqExpr{
					pos: position{line: 244, col: 18, offset: 6247},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 244, col: 18, offset: 6247},
							label: "cc",
							expr: &charClassMatcher{
								pos:        position{line: 244, col: 21, offset: 6250},
								val:        "[Aa]",
								chars:      []rune{'A', 'a'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 244, col: 26, offset: 6255},
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 26, offset: 6255},
								name: "wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 31, offset: 6260},
							label: "args",
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 36, offset: 6265},
								name: "elliptical_arc_argument_sequence",
							},
						},
					},
				},
			},
		},
		{
			name: "elliptical_arc_argument_sequence",
			pos:  position{line: 248, col: 1, offset: 6356},
			expr: &actionExpr{
				pos: position{line: 248, col: 36, offset: 6391},
				run: (*parser).callonelliptical_arc_argument_sequence1,
				expr: &seqExpr{
					pos: position{line: 248, col: 36, offset: 6391},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 36, offset: 6391},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 42, offset: 6397},
								name: "elliptical_arc_argument",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 66, offset: 6421},
							label: "more",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 71, offset: 6426},
								expr: &seqExpr{
									pos: position{line: 248, col: 72, offset: 6427},
									exprs: []interface{}{
										&zeroOrOneExpr{
											pos: position{line: 248, col: 72, offset: 6427},
											expr: &ruleRefExpr{
												pos:  position{line: 248, col: 72, offset: 6427},
												name: "comma_wsp",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 248, col: 83, offset: 6438},
											name: "elliptical_arc_argument",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "elliptical_arc_argument",
			pos:  position{line: 252, col: 1, offset: 6513},
			expr: &actionExpr{
				pos: position{line: 252, col: 27, offset: 6539},
				run: (*parser).callonelliptical_arc_argument1,
				expr: &seqExpr{
					pos: position{line: 252, col: 27, offset: 6539},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 252, col: 27, offset: 6539},
							label: "rx",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 30, offset: 6542},
								name: "nonnegative_number",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 252, col: 49, offset: 6561},
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 49, offset: 6561},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 252, col: 60, offset: 6572},
							label: "ry",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 63, offset: 6575},
								name: "nonnegative_number",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 252, col: 82, offset: 6594},
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 82, offset: 6594},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 252, col: 93, offset: 6605},
							label: "xrot",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 98, offset: 6610},
								name: "number",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 252, col: 105, offset: 6617},
							name: "comma_wsp",
						},
						&labeledExpr{
							pos:   position{line: 252, col: 115, offset: 6627},
							label: "large",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 121, offset: 6633},
								name: "flag",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 252, col: 126, offset: 6638},
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 126, offset: 6638},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 252, col: 137, offset: 6649},
							label: "sweep",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 143, offset: 6655},
								name: "flag",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 252, col: 148, offset: 6660},
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 148, offset: 6660},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 252, col: 159, offset: 6671},
							label: "xy",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 162, offset: 6674},
								name: "coordinate_pair",
							},
						},
					},
				},
			},
		},
		{
			name: "coordinate_pair",
			pos:  position{line: 266, col: 1, offset: 6925},
			expr: &actionExpr{
				pos: position{line: 266, col: 19, offset: 6943},
				run: (*parser).calloncoordinate_pair1,
				expr: &seqExpr{
					pos: position{line: 266, col: 19, offset: 6943},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 266, col: 19, offset: 6943},
							label: "x",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 21, offset: 6945},
								name: "number",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 266, col: 28, offset: 6952},
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 28, offset: 6952},
								name: "comma_wsp",
							},
						},
						&labeledExpr{
							pos:   position{line: 266, col: 39, offset: 6963},
							label: "y",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 41, offset: 6965},
								name: "number",
							},
						},
					},
				},
			},
		},
		{
			name: "nonnegative_number",
			pos:  position{line: 274, col: 1, offset: 7052},
			expr: &choiceExpr{
				pos: position{line: 274, col: 22, offset: 7073},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 274, col: 22, offset: 7073},
						name: "floating_point_constant",
					},
					&actionExpr{
						pos: position{line: 274, col: 48, offset: 7099},
						run: (*parser).callonnonnegative_number3,
						expr: &ruleRefExpr{
							pos:  position{line: 274, col: 48, offset: 7099},
							name: "digit_sequence",
						},
					},
				},
			},
		},
		{
			name: "number",
			pos:  position{line: 278, col: 1, offset: 7169},
			expr: &choiceExpr{
				pos: position{line: 278, col: 10, offset: 7178},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 278, col: 10, offset: 7178},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 278, col: 10, offset: 7178},
								expr: &ruleRefExpr{
									pos:  position{line: 278, col: 10, offset: 7178},
									name: "sign",
								},
							},
							&ruleRefExpr{
								pos:  position{line: 278, col: 16, offset: 7184},
								name: "floating_point_constant",
							},
						},
					},
					&actionExpr{
						pos: position{line: 278, col: 42, offset: 7210},
						run: (*parser).callonnumber6,
						expr: &seqExpr{
							pos: position{line: 278, col: 42, offset: 7210},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 278, col: 42, offset: 7210},
									expr: &ruleRefExpr{
										pos:  position{line: 278, col: 42, offset: 7210},
										name: "sign",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 278, col: 48, offset: 7216},
									name: "digit_sequence",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "flag",
			pos:  position{line: 282, col: 1, offset: 7286},
			expr: &actionExpr{
				pos: position{line: 282, col: 8, offset: 7293},
				run: (*parser).callonflag1,
				expr: &charClassMatcher{
					pos:        position{line: 282, col: 8, offset: 7293},
					val:        "[01]",
					chars:      []rune{'0', '1'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "comma_wsp",
			pos:  position{line: 289, col: 1, offset: 7377},
			expr: &choiceExpr{
				pos: position{line: 289, col: 13, offset: 7389},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 289, col: 14, offset: 7390},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 289, col: 14, offset: 7390},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 14, offset: 7390},
									name: "wsp",
								},
							},
							&zeroOrOneExpr{
								pos: position{line: 289, col: 19, offset: 7395},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 19, offset: 7395},
									name: "comma",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 289, col: 26, offset: 7402},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 26, offset: 7402},
									name: "wsp",
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 289, col: 35, offset: 7411},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 289, col: 35, offset: 7411},
								name: "comma",
							},
							&zeroOrMoreExpr{
								pos: position{line: 289, col: 41, offset: 7417},
								expr: &ruleRefExpr{
									pos:  position{line: 289, col: 41, offset: 7417},
									name: "wsp",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "comma",
			pos:  position{line: 291, col: 1, offset: 7424},
			expr: &litMatcher{
				pos:        position{line: 291, col: 9, offset: 7432},
				val:        ",",
				ignoreCase: false,
			},
		},
		{
			name: "floating_point_constant",
			pos:  position{line: 293, col: 1, offset: 7437},
			expr: &choiceExpr{
				pos: position{line: 293, col: 27, offset: 7463},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 293, col: 27, offset: 7463},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 293, col: 27, offset: 7463},
								name: "fractional_constant",
							},
							&zeroOrOneExpr{
								pos: position{line: 293, col: 47, offset: 7483},
								expr: &ruleRefExpr{
									pos:  position{line: 293, col: 47, offset: 7483},
									name: "exponent",
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 293, col: 59, offset: 7495},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 293, col: 59, offset: 7495},
								name: "digit_sequence",
							},
							&ruleRefExpr{
								pos:  position{line: 293, col: 74, offset: 7510},
								name: "exponent",
							},
						},
					},
				},
			},
		},
		{
			name: "fractional_constant",
			pos:  position{line: 295, col: 1, offset: 7520},
			expr: &choiceExpr{
				pos: position{line: 295, col: 23, offset: 7542},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 295, col: 23, offset: 7542},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 295, col: 23, offset: 7542},
								expr: &ruleRefExpr{
									pos:  position{line: 295, col: 23, offset: 7542},
									name: "digit_sequence",
								},
							},
							&litMatcher{
								pos:        position{line: 295, col: 39, offset: 7558},
								val:        ".",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 295, col: 43, offset: 7562},
								name: "digit_sequence",
							},
						},
					},
					&seqExpr{
						pos: position{line: 295, col: 60, offset: 7579},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 295, col: 60, offset: 7579},
								name: "digit_sequence",
							},
							&litMatcher{
								pos:        position{line: 295, col: 75, offset: 7594},
								val:        ".",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "exponent",
			pos:  position{line: 297, col: 1, offset: 7599},
			expr: &seqExpr{
				pos: position{line: 297, col: 12, offset: 7610},
				exprs: []interface{}{
					&charClassMatcher{
						pos:        position{line: 297, col: 12, offset: 7610},
						val:        "[eE]",
						chars:      []rune{'e', 'E'},
						ignoreCase: false,
						inverted:   false,
					},
					&zeroOrOneExpr{
						pos: position{line: 297, col: 17, offset: 7615},
						expr: &ruleRefExpr{
							pos:  position{line: 297, col: 17, offset: 7615},
							name: "sign",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 297, col: 23, offset: 7621},
						name: "digit_sequence",
					},
				},
			},
		},
		{
			name: "sign",
			pos:  position{line: 299, col: 1, offset: 7637},
			expr: &charClassMatcher{
				pos:        position{line: 299, col: 8, offset: 7644},
				val:        "[+-]",
				chars:      []rune{'+', '-'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "digit_sequence",
			pos:  position{line: 301, col: 1, offset: 7650},
			expr: &oneOrMoreExpr{
				pos: position{line: 301, col: 18, offset: 7667},
				expr: &charClassMatcher{
					pos:        position{line: 301, col: 18, offset: 7667},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "wsp",
			pos:  position{line: 303, col: 1, offset: 7675},
			expr: &charClassMatcher{
				pos:        position{line: 303, col: 7, offset: 7681},
				val:        "[ \\t\\n\\r]",
				chars:      []rune{' ', '\t', '\n', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
	},
}

func (c *current) onsvg_path1(data interface{}) (interface{}, error) {
	// according to the specification the first moveto is never relative
	if data == nil {
		return nil, nil
	}
	d := data.([]interface{})
	if len(d) > 0 {
		m := d[0].(map[string]interface{})
		r, ok := m["relative"]
		if ok && r.(bool) {
			delete(m, "relative")
		}
		m["code"] = "M"
	}
	return d, nil
}

func (p *parser) callonsvg_path1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsvg_path1(stack["data"])
}

func (c *current) onmoveTo_drawTo_commandGroups1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonmoveTo_drawTo_commandGroups1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onmoveTo_drawTo_commandGroups1(stack["first"], stack["more"])
}

func (c *current) onmoveTo_drawTo_commandGroup1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonmoveTo_drawTo_commandGroup1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onmoveTo_drawTo_commandGroup1(stack["first"], stack["more"])
}

func (c *current) onmoveto1(cc, first, more interface{}) (interface{}, error) {
	firstCmd, err := commands(string(cc.([]byte)), []interface{}{first})
	if err != nil {
		return nil, err
	}
	if more == nil {
		return firstCmd, nil
	}
	moreCmds, err := commands(string(cc.([]byte)), more.([]interface{}))
	if err != nil {
		return nil, err
	}
	var code string
	if cc == "M" {
		code = "L"
	} else {
		code = "l"
	}
	for _, cmd := range moreCmds {
		c := cmd.(map[string]interface{})
		c["code"] = code
	}
	return merge(firstCmd, moreCmds)
}

func (p *parser) callonmoveto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onmoveto1(stack["cc"], stack["first"], stack["more"])
}

func (c *current) onclosepath1() (interface{}, error) {
	return commands("Z", nil)
}

func (p *parser) callonclosepath1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onclosepath1()
}

func (c *current) onlineto1(cc, args interface{}) (interface{}, error) {
	return commands(string(cc.([]byte)), args.([]interface{}))
}

func (p *parser) callonlineto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onlineto1(stack["cc"], stack["args"])
}

func (c *current) onlineto_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge([]interface{}{first}, []interface{}{more})
}

func (p *parser) callonlineto_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onlineto_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) onhorizontal_lineto1(cc, args interface{}) (interface{}, error) {
	var mArgs []interface{}
	for _, x := range args.([]interface{}) {
		mArgs = append(mArgs, map[string]interface{}{"x": x})
	}
	return commands(string(cc.([]byte)), mArgs)
}

func (p *parser) callonhorizontal_lineto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onhorizontal_lineto1(stack["cc"], stack["args"])
}

func (c *current) oncoordinate_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) calloncoordinate_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.oncoordinate_sequence1(stack["first"], stack["more"])
}

func (c *current) onvertical_lineto1(cc, args interface{}) (interface{}, error) {
	var mArgs []interface{}
	for _, y := range args.([]interface{}) {
		mArgs = append(mArgs, map[string]interface{}{"y": y})
	}
	return commands(cc.(string), mArgs)
}

func (p *parser) callonvertical_lineto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onvertical_lineto1(stack["cc"], stack["args"])
}

func (c *current) oncurveto1(cc, args interface{}) (interface{}, error) {
	return commands(cc.(string), args.([]interface{}))
}

func (p *parser) calloncurveto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.oncurveto1(stack["cc"], stack["args"])
}

func (c *current) oncurveto_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) calloncurveto_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.oncurveto_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) oncurveto_argument1(a, b, cc interface{}) (interface{}, error) {
	aMap := a.(map[string]interface{})
	bMap := b.(map[string]interface{})
	ccMap := cc.(map[string]interface{})
	m := map[string]interface{}{
		"x1": aMap["x"],
		"y1": aMap["y"],
		"x2": bMap["x"],
		"y2": bMap["y"],
		"x":  ccMap["x"],
		"y":  ccMap["y"],
	}
	return m, nil
}

func (p *parser) calloncurveto_argument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.oncurveto_argument1(stack["a"], stack["b"], stack["cc"])
}

func (c *current) onsmooth_curveto1(cc, args interface{}) (interface{}, error) {
	return commands(cc.(string), args.([]interface{}))
}

func (p *parser) callonsmooth_curveto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsmooth_curveto1(stack["cc"], stack["args"])
}

func (c *current) onsmooth_curveto_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonsmooth_curveto_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsmooth_curveto_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) onsmooth_curveto_argument1(b, cp interface{}) (interface{}, error) {
	bMap := b.(map[string]interface{})
	cpMap := cp.(map[string]interface{})
	m := map[string]interface{}{
		"x2": bMap["x"],
		"y2": bMap["y"],
		"x":  cpMap["x"],
		"y":  cpMap["y"],
	}
	return m, nil
}

func (p *parser) callonsmooth_curveto_argument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsmooth_curveto_argument1(stack["b"], stack["cp"])
}

func (c *current) onquadratic_bezier_curveto1(cc, args interface{}) (interface{}, error) {
	return commands(cc.(string), args.([]interface{}))
}

func (p *parser) callonquadratic_bezier_curveto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onquadratic_bezier_curveto1(stack["cc"], stack["args"])
}

func (c *current) onquadratic_bezier_curveto_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonquadratic_bezier_curveto_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onquadratic_bezier_curveto_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) onquadratic_bezier_curveto_argument1(a, b interface{}) (interface{}, error) {
	aMap := a.(map[string]interface{})
	bMap := b.(map[string]interface{})
	m := map[string]interface{}{
		"x1": aMap["x"],
		"y1": aMap["y"],
		"x":  bMap["x"],
		"y":  bMap["y"],
	}
	return m, nil
}

func (p *parser) callonquadratic_bezier_curveto_argument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onquadratic_bezier_curveto_argument1(stack["a"], stack["b"])
}

func (c *current) onsmooth_quadratic_bezier_curveto1(cc, args interface{}) (interface{}, error) {
	return commands(cc.(string), args.([]interface{}))
}

func (p *parser) callonsmooth_quadratic_bezier_curveto1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsmooth_quadratic_bezier_curveto1(stack["cc"], stack["args"])
}

func (c *current) onsmooth_quadratic_bezier_curveto_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonsmooth_quadratic_bezier_curveto_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onsmooth_quadratic_bezier_curveto_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) onelliptical_arc1(cc, args interface{}) (interface{}, error) {
	return commands(cc.(string), args.([]interface{}))
}

func (p *parser) callonelliptical_arc1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onelliptical_arc1(stack["cc"], stack["args"])
}

func (c *current) onelliptical_arc_argument_sequence1(first, more interface{}) (interface{}, error) {
	return merge(first, more.([]interface{}))
}

func (p *parser) callonelliptical_arc_argument_sequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onelliptical_arc_argument_sequence1(stack["first"], stack["more"])
}

func (c *current) onelliptical_arc_argument1(rx, ry, xrot, large, sweep, xy interface{}) (interface{}, error) {
	xyMap := xy.(map[string]interface{})
	m := map[string]interface{}{
		"rx":            rx,
		"ry":            ry,
		"xAxisRotation": xrot,
		"largeArc":      large,
		"sweep":         sweep,
		"x":             xyMap["x"],
		"y":             xyMap["y"],
	}
	return m, nil
}

func (p *parser) callonelliptical_arc_argument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onelliptical_arc_argument1(stack["rx"], stack["ry"], stack["xrot"], stack["large"], stack["sweep"], stack["xy"])
}

func (c *current) oncoordinate_pair1(x, y interface{}) (interface{}, error) {
	m := map[string]interface{}{
		"x": x,
		"y": y,
	}
	return m, nil
}

func (p *parser) calloncoordinate_pair1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.oncoordinate_pair1(stack["x"], stack["y"])
}

func (c *current) onnonnegative_number3() (interface{}, error) {
	return strconv.ParseFloat(string(c.text), 64)
}

func (p *parser) callonnonnegative_number3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnonnegative_number3()
}

func (c *current) onnumber6() (interface{}, error) {
	return strconv.ParseFloat(string(c.text), 64)
}

func (p *parser) callonnumber6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onnumber6()
}

func (c *current) onflag1() (interface{}, error) {
	if string(c.text) == "1" {
		return true, nil
	}
	return false, nil
}

func (p *parser) callonflag1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onflag1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n > 0 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
