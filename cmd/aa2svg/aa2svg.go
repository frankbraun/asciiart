// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// aa2svg converts ASCII diagrams to SVGs.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/asciiart/aa2d"
	"github.com/frankbraun/asciiart/convert"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: [arguments...]", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var (
		in     string
		out    string
		xScale int
		yScale int
	)
	flag.StringVar(&in, "i", "", "path to input text file. If unspecified or "+
		"set to '-', stdin is used")
	flag.StringVar(&out, "o", "", "path to output SVG file. If unspecified or "+
		"set to '-', stdout is used")
	flag.IntVar(&xScale, "x", aa2d.XScale,
		"number of pixels to scale each unit on the x-axis to")
	flag.IntVar(&yScale, "y", aa2d.YScale,
		"number of pixels to scale each unit on the y-axis to")
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
	}
	// work around defer not working after os.Exit()
	if err := aa2svgMain(out, in, xScale, yScale); err != nil {
		fatal(err)
	}
}

func aa2svgMain(out, in string, xScale, yScale int) error {
	var (
		outFP *os.File
		inFP  *os.File
		err   error
	)
	if out == "" || out == "-" {
		outFP = os.Stdout
	} else {
		if _, err := os.Stat(out); err == nil {
			return fmt.Errorf("output file '%s' exists already", out)
		}
		outFP, err = os.Create(out)
		if err != nil {
			return err
		}
		defer outFP.Close()
	}
	if in == "" || in == "-" {
		inFP = os.Stdin
	} else {
		inFP, err = os.Open(in)
		if err != nil {
			return err
		}
		defer inFP.Close()
	}
	return convert.ASCIIArt2SVG(outFP, inFP, xScale, yScale)
}
