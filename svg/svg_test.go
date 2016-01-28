// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package svg

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/frankbraun/asciiart"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	p := asciiart.NewParser()
	fis, err := ioutil.ReadDir("testdata")
	assert.NoError(t, err)
	for _, fi := range fis {
		inFN := fi.Name()
		if strings.HasSuffix(inFN, ".txt") {
			input, err := ioutil.ReadFile(filepath.Join("testdata", inFN))
			assert.NoError(t, err)
			g, err := p.Parse(string(input))
			assert.NoError(t, err)
			var svg bytes.Buffer
			err = Generate(&svg, g)
			assert.NoError(t, err)
			outFN := strings.TrimSuffix(inFN, ".txt") + ".svg"
			output, err := ioutil.ReadFile(filepath.Join("testdata", outFN))
			assert.NoError(t, err)
			assert.Equal(t, string(output), svg.String())
		}
	}
}
