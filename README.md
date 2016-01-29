[comment]: # (This file is generated from templates/readme.tmpl, do not edit!)
## Parsing hierarchical ASCII art for fun and profit [![GoDoc](https://godoc.org/github.com/frankbraun/asciiart?status.png)](http://godoc.org/github.com/frankbraun/asciiart) [![Build Status](https://travis-ci.org/frankbraun/asciiart.png)](https://travis-ci.org/frankbraun/asciiart)

This is **alpha** software and the API is still in flux.


### Motivation

All ASCII art converters I know combine the ASCII parser and the graphics
generator into one package. This project started with the observation that
some ASCII art artifacts are *hierarchical* (often they comprise a *tree*) and
that it might be cool to have an abstract syntax tree (AST) of the parsed
ASCII art to use in *different* backends, a SVG generator just being one of
them.

The vision is to allow [model-oriented programming](https://github.com/imatix/gsl#model-oriented-programming) with hierarchical ASCII art as the modeling language...


### API example

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/frankbraun/asciiart"
)

var asciiArt = `
#-------------------#
|[REF]              |
|                   |
| .---.       .---. |
| |foo| ====> |bar| |
| '---'       '---' |
|                   |
|                   |
#-------------------#

[REF]: {"note":"outer"}
`

func main() {
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
				fmt.Println(indent, "polygon", strings.Join(p, " "), fmtJsn(t.Ref))
				traverse(t.Elems, indent+"  ") // recursion
			case *asciiart.Textline:
				fmt.Println(indent, "textline", t.X, t.Y, t.Text)
			}
		}
	}
	traverse(grid.Elems, " ")
}

```


#### Output

```
grid 21 10
  rect 0.5 1.5 20 8 {"note":"outer"}
    rect 2.5 4.5 4 2 null
      textline 3.5 5.5 foo
    rect 14.5 4.5 4 2 null
      textline 15.5 5.5 bar
    line 8.5 5.5 12.5 5.5
```


### Converting to SVG with `aa2svg`

This package also contains the tool `aa2svg` which uses the AST given by ASCII
art parser to generate SVGs.


#### Installation

```
go get -v github.com/frankbraun/asciiart/cmd/aa2svg
```


#### Usage

```
Usage of aa2svg:
  -f	overwrite existing output file
  -i string
    	path to input text file. If unspecified or set to '-', stdin is used
  -o string
    	path to output SVG file. If unspecified or set to '-', stdout is used
  -x float
    	number of pixels to scale each unit on the x-axis to (default 9)
  -y float
    	number of pixels to scale each unit on the y-axis to (default 16)
```


#### Example output
![Example SVG](https://rawgit.com/frankbraun/asciiart/master/templates/exampleart.svg)

Appending the following global reference to the ASCII art disables the blur
effect:

```
[_SVG]: {"blur": false}
```		


### Credits

This package was inspired by [ASCIIToSVG](https://github.com/dhobsd/asciitosvg/), the ASCII to SVG converter used to render the graphics in the
[ZeroMQ guide](http://zguide.zeromq.org/).


### Open research question

How does one formally define a grammar for two-dimensional ASCII art?
ENBF doesn't seem to be up for the job...