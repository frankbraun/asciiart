package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/frankbraun/asciiart/aa2d"
)

var asciiArt = `
#--------------#
|              |
| #---#  #---# |
| |   |  |   | |
| #---#  #---# |
|              |
#--------------#
`

func main() {
	p := aa2d.NewParser()
	p.SetScale(1, 1)
	grid, err := p.Parse(asciiArt)
	if err != nil {
		log.Fatal(err)
	}
	fmtFlt := func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }
	var traverse func(elems []interface{}, indent string)
	fmt.Println("grid", grid.W, grid.H)
	traverse = func(elems []interface{}, indent string) {
		for _, elem := range elems {
			switch t := elem.(type) {
			case *aa2d.Rectangle:
				fmt.Println(indent, "rect", t.X, t.Y, t.W, t.H)
				traverse(t.Elems, indent+"  ") // recursion
			case *aa2d.Line:
				fmt.Println(indent, "line", t.X1, t.Y1, t.X2, t.Y2)
			case *aa2d.Polyline:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Println(indent, "polyline", strings.Join(p, " "))
			case *aa2d.Polygon:
				var p []string
				for i := 0; i < len(t.X); i++ {
					p = append(p, fmtFlt(t.X[i]), fmtFlt(t.Y[i]))
				}
				fmt.Println(indent, "polygon", strings.Join(p, " "))
				traverse(t.Elems, indent+"  ") // recursion
			case *aa2d.Textline:
				fmt.Println(indent, "textline", t.X, t.Y, t.Text)
			}
		}
	}
	traverse(grid.Elems, " ")
}
