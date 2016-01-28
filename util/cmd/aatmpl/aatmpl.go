// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// aatmpl is the code and documentation generator for the ASCII art package.
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func main() {
	// work around defer not working after os.Exit()
	if err := aatmplMain(); err != nil {
		fatal(err)
	}
}

func aatmplMain() error {
	templates := filepath.Join("templates", "*.tmpl")
	tmpl, err := template.ParseGlob(templates)
	if err != nil {
		return err
	}

	// generate cmd/aa2txt/aa2txt.go
	fp, err := os.Create(filepath.Join("cmd", "aa2txt", "aa2txt.go"))
	if err != nil {
		return err
	}
	defer fp.Close()
	err = tmpl.ExecuteTemplate(fp, "aa2txt.tmpl", map[string]string{
		"MainFunc": "toTxt(asciiArt string)",
	})
	if err != nil {
		return err
	}

	// generate example.go
	example := filepath.Join("util", "cmd", "aaexample", "aaexample.go")
	fp, err = os.Create(example)
	if err != nil {
		return err
	}
	defer fp.Close()
	err = tmpl.ExecuteTemplate(fp, "exampleprog.tmpl", map[string]string{
		"MainFunc": "main()",
	})
	if err != nil {
		return err
	}

	// execute example.go and capture output for README.md
	output, err := runExample(example)
	if err != nil {
		return err
	}

	// generate README.md
	fp, err = os.Create("README.md")
	if err != nil {
		return err
	}
	defer fp.Close()
	err = tmpl.ExecuteTemplate(fp, "readme.tmpl", map[string]string{
		"MainFunc": "main()",
		"Output":   output,
	})
	if err != nil {
		return err
	}

	return nil
}

func runExample(example string) (string, error) {
	cmd := exec.Command("go", "run", example)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
