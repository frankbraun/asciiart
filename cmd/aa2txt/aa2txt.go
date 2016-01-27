// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// aa2svg converts ASCII diagrams to ASCII trees.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frankbraun/asciiart/convert"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [arguments...]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Read ASCII art diagram from stdin and write tree to stdout.\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.NArg() != 0 {
		usage()
	}
	if err := convert.ASCIIArt2Txt(os.Stdout, os.Stdin); err != nil {
		fatal(err)
	}
}
