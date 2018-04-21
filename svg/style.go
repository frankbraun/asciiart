package svg

import (
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
        "effects": [
          {
            "feOffset": {
              "-dx": "3",
              "-dy": "3",
              "-in": "SourceGraphic",
              "-result": "offOut"
            }
          },
          {
            "feColorMatrix": {
              "-in": "offOut",
              "-result": "matrixOut",
              "-type": "matrix",
              "-values": "0.2 0 0 0 0 0 0.2 0 0 0 0 0 0.2 0 0 0 0 0 1 0"
            }
          },
          {
            "feBlend": {
              "-in": "SourceGraphic",
              "-in2": "matrixOut",
              "-mode": "normal"
            }
          }
        ]
      },
      {
        "-height": "150%",
        "-id": "dsFilter",
        "-width": "150%",
        "effects": [
          {
            "feOffset": {
              "-dx": "3",
              "-dy": "3",
              "-in": "SourceGraphic",
              "-result": "offOut"
            }
          },
          {
            "feColorMatrix": {
              "-in": "offOut",
              "-result": "matrixOut",
              "-type": "matrix",
              "-values": "0.2 0 0 0 0 0 0.2 0 0 0 0 0 0.2 0 0 0 0 0 1 0"
            }
          },
          {
            "feGaussianBlur": {
              "-in": "matrixOut",
              "-result": "blurOut",
              "-stdDeviation": "3"
            }
          },
          {
            "feBlend": {
              "-in": "SourceGraphic",
              "-in2": "blurOut",
              "-mode": "normal"
            }
          }
        ]
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
    "-fill": "white",
    "-stroke": "black",
    "-stroke-width": "2"
  },
  "line": {
    "-fill": "none",
    "-stroke": "black",
    "-stroke-width": "2"
  },
  "text": {
    "-fill": "black",
    "-stroke": "none",
    "-font-family": "Consolas,Monaco,Anonymous Pro,Anonymous,Bitstream Sans Mono,monospace"
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

func getAttributes(key string) ([]string, error) {
	paths, err := styleMap.ValuesForKey(key)
	if err != nil {
		return nil, err
	}
	if len(paths) > 1 {
		return nil, fmt.Errorf("svg: just one %s style expected", key)
	}
	p, ok := paths[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("svg: could not cast %s map", key)
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
		s[i] = fmt.Sprintf("%s=%q", strings.TrimPrefix(k, "-"), p[k].(string))
	}
	return s, nil
}

func rectStyle(blur bool) ([]string, error) {
	s, err := getAttributes("rect")
	if err != nil {
		return nil, err
	}
	if blur {
		s = append(s, `filter="url(#dsFilter)"`)
	} else {
		s = append(s, `filter="url(#dsFilterNoBlur)"`)
	}
	return s, nil
}

func lineStyle() ([]string, error) {
	s, err := getAttributes("line")
	if err != nil {
		return nil, err
	}
	return s, nil
}

func textStyle(yScale float64) ([]string, error) {
	s, err := getAttributes("text")
	if err != nil {
		return nil, err
	}
	// set font size "empirically", inspired by:
	// https://github.com/dhobsd/asciitosvg/blob/05f2ac06918247a79561b026a6a8011a64a98317/ASCIIToSVG.php#L1729-L1741
	s = append(s, fmt.Sprintf("font-size=\"%fpx\"", 0.95*yScale))
	return s, nil
}
