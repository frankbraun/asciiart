[comment]: # (This file is generated from templates/readme.tmpl, do not edit!)
## Parsing hierarchical ASCII art for fun and profit [![GoDoc](https://godoc.org/github.com/frankbraun/asciiart?status.png)](http://godoc.org/github.com/frankbraun/asciiart) [![Build Status](https://travis-ci.org/frankbraun/asciiart.png)](https://travis-ci.org/frankbraun/asciiart)

This is **alpha** software and the API is still in flux.


### API example

```go
package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/frankbraun/asciiart"
)

var asciiArt = `
#-----------------#
|                 |
| #---#     #---# |
| |foo| --> |bar| |
| #---#     #---# |
|                 |
#-----------------#
`

func main() {
	p := asciiart.NewParser()
	p.SetScale(1, 1)
	grid, err := p.Parse(asciiArt)
	if err != nil {
		log.Fatal(err)
	}
	var traverse func(elems []interface{}, indent string)
	fmtFlt := func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }
	fmt.Println("grid", grid.W, grid.H)
	traverse = func(elems []interface{}, indent string) {
		for _, elem := range elems {
			switch t := elem.(type) {
			case *asciiart.Rectangle:
				fmt.Println(indent, "rect", t.X, t.Y, t.W, t.H)
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
				traverse(t.Elems, indent+"  ") // recursion
			case *asciiart.Textline:
				fmt.Println(indent, "textline", t.X, t.Y, t.Text)
			}
		}
	}
	traverse(grid.Elems, " ")
}

```


### Output

```
grid 19 8
  rect 0.5 1.5 18 6
    rect 2.5 3.5 4 2
      textline 3.5 4.5 foo
    rect 12.5 3.5 4 2
      textline 13.5 4.5 bar
    line 8.5 4.5 10.5 4.5
```


## Converting to SVG

The tool `aa2svg` uses abstract representation delivered by the ASCII parser to generate SVGs like this:

![Example SVG](https://rawgit.com/frankbraun/asciiart/master/templates/exampleart.svg)


### Credits

This package was inspired by [ASCIIToSVG](https://github.com/dhobsd/asciitosvg/), the ASCII to SVG converter used to render the graphics in the
[ZeroMQ guide](http://zguide.zeromq.org/).
