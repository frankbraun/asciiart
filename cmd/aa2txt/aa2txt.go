// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
//
// This code is generated, do not edit!

// aa2svg converts ASCII diagrams to ASCII trees.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/frankbraun/asciiart"
)

func main() {
	asciiArt, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	toTxt(string(asciiArt))
}

func toTxt(asciiArt string) {
	p := asciiart.NewParser()
	p.SetScale(1, 1)
	grid, err := p.Parse(asciiArt)
	if err != nil {
		log.Fatal(err)
	}
	var traverse func(elems []interface{}, indent string)
	fmtJsn := func(j map[string]interface{}) string { b, _ := json.Marshal(j); return string(b) }
	fmtFlt := func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }
	fmt.Println("grid", grid.W, grid.H)
	traverse = func(elems []interface{}, indent string) {
		for _, elem := range elems {
			switch t := elem.(type) {
			case *asciiart.Rectangle:
				fmt.Println(indent, "rect", t.X, t.Y, t.W, t.H, fmtJsn(t.Ref))
				traverse(t.Elems, indent+"  ") // recursion
			case *asciiart.Line:
				fmt.Println(indent, "line", t.X1, t.Y1, t.X2, t.Y2)
			case *asciiart.Polyline:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Println(indent, "polyline", strings.Join(p, " "))
			case *asciiart.Polygon:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Println(indent, "polygon", strings.Join(p, " "))
			case *asciiart.Textline:
				fmt.Println(indent, "textline", t.X, t.Y, t.Text)
			}
		}
	}
	traverse(grid.Elems, " ")
}
