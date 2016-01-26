// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package svg

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/clbanning/mxj"
)

// styleJSON defines the default SVG style in JSON.
const styleJSON = `
{
  "defs": {
    "filter": [
      {
        "-height": "150%",
        "-id": "dsFilterNoBlur",
        "-width": "150%",
        "feBlend": {
          "-in": "SourceGraphic",
          "-in2": "matrixOut",
          "-mode": "normal"
        },
        "feColorMatrix": {
          "-in": "offOut",
          "-result": "matrixOut",
          "-type": "matrix",
          "-values": "0.2 0 0 0 0 0 0.2 0 0 0 0 0 0.2 0 0 0 0 0 1 0"
        },
        "feOffset": {
          "-dx": "3",
          "-dy": "3",
          "-in": "SourceGraphic",
          "-result": "offOut"
        }
      },
      {
        "-height": "150%",
        "-id": "dsFilter",
        "-width": "150%",
        "feBlend": {
          "-in": "SourceGraphic",
          "-in2": "blurOut",
          "-mode": "normal"
        },
        "feColorMatrix": {
          "-in": "offOut",
          "-result": "matrixOut",
          "-type": "matrix",
          "-values": "0.2 0 0 0 0 0 0.2 0 0 0 0 0 0.2 0 0 0 0 0 1 0"
        },
        "feGaussianBlur": {
          "-in": "matrixOut",
          "-result": "blurOut",
          "-stdDeviation": "3"
        },
        "feOffset": {
          "-dx": "3",
          "-dy": "3",
          "-in": "SourceGraphic",
          "-result": "offOut"
        }
      }
    ],
    "marker": [
      {
        "-id": "iPointer",
        "-markerHeight": "7",
        "-markerUnits": "strokeWidth",
        "-markerWidth": "8",
        "-orient": "auto",
        "-refX": "5",
        "-refY": "5",
        "-viewBox": "0 0 10 10",
        "path": {
          "-d": "M 10 0 L 10 10 L 0 5 z"
        }
      },
      {
        "-id": "Pointer",
        "-markerHeight": "7",
        "-markerUnits": "strokeWidth",
        "-markerWidth": "8",
        "-orient": "auto",
        "-refX": "5",
        "-refY": "5",
        "-viewBox": "0 0 10 10",
        "path": {
          "-d": "M 0 0 L 10 5 L 0 10 z"
        }
      }
    ]
  },
  "rect": {
    "-fill": "none",
    "-stroke": "black",
    "-stroke-width": "2"
  }
}`

var styleMap mxj.Map

func init() {
	var err error
	styleMap, err = mxj.NewMapJson([]byte(styleJSON))
	if err != nil {
		log.Panic(err)
	}
}

func rectStyle() ([]string, error) {
	paths, err := styleMap.ValuesForKey("rect")
	if err != nil {
		return nil, err
	}
	if len(paths) > 1 {
		return nil, errors.New("svg: just one rectangle style expected")
	}
	p, ok := paths[0].(map[string]interface{})
	if !ok {
		return nil, errors.New("svg: could not cast rectangle map")
	}
	var keys []string
	for k := range p {
		// only process attributes
		if strings.HasPrefix(k, "-") {
			keys = append(keys, k)
		} else {
			return nil, fmt.Errorf("svg: did expect attribute: %s", k)
		}
	}
	sort.Strings(keys)
	s := make([]string, len(keys))
	for i, k := range keys {
		s[i] = fmt.Sprintf("%s=%q", strings.TrimPrefix(k, "-"), p[k])
	}
	return s, nil
}
