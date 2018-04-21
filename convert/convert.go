// Package convert implements converters for two-dimensional hierarchical
// ASCII art.
package convert

import (
	"io"
	"io/ioutil"

	"github.com/frankbraun/asciiart"
	"github.com/frankbraun/asciiart/svg"
)

// ASCIIArt2SVG converts two-dimensional ASCII art read from r to a SVG
// written to w.
// xScale denotes the number of pixels to scale each unit on the x-axis to.
// yScale denotes the number of pixels to scale each unit on the y-axis to.
func ASCIIArt2SVG(w io.Writer, r io.Reader, xScale, yScale float64) error {
	aa, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	p := asciiart.NewParser()
	if err := p.SetScale(xScale, yScale); err != nil {
		return err
	}
	grid, err := p.Parse(string(aa))
	if err != nil {
		return err
	}
	return svg.Generate(w, grid)
}
