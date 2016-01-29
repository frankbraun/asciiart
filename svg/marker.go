// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package svg

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ajstarks/svgo/float"
)

func setMarker(s *svg.SVG) error {
	paths, err := styleMap.ValuesForPath("defs.marker")
	if err != nil {
		return err
	}
	for _, path := range paths {
		p, ok := path.(map[string]interface{})
		if !ok {
			return errors.New("svg: could not cast filter map")
		}
		var (
			id     string   // id attribute
			x      float64  // refX
			y      float64  // refY
			width  float64  // markerWidth
			height float64  // markerHeight
			attr   []string // other attributes
			path   map[string]interface{}
		)
		for k, v := range p {
			switch {
			case k == "-id":
				id = v.(string)
			case k == "-refX":
				x, err = strconv.ParseFloat(v.(string), 64)
				if err != nil {
					return err
				}
			case k == "-refY":
				y, err = strconv.ParseFloat(v.(string), 64)
				if err != nil {
					return err
				}
			case k == "-markerWidth":
				width, err = strconv.ParseFloat(v.(string), 64)
				if err != nil {
					return err
				}
			case k == "-markerHeight":
				height, err = strconv.ParseFloat(v.(string), 64)
				if err != nil {
					return err
				}
			case strings.HasPrefix(k, "-"):
				attr = append(attr, k)
			case k == "path":
				path = v.(map[string]interface{})
			default:
				return fmt.Errorf("svg: unknown filter: %s", k)
			}
		}
		if id == "" {
			return errors.New("svg: marker id undefined")
		}
		d, ok := path["-d"]
		if !ok {
			return errors.New("svg: marker path has no 'd' attribute")
		}
		sort.Strings(attr)
		a := make([]string, len(attr))
		for i, k := range attr {
			a[i] = fmt.Sprintf("%s=%q", strings.TrimPrefix(k, "-"),
				p[k].(string))
		}
		s.Marker(id, x, y, width, height, a...)
		s.Path(d.(string))
		s.MarkerEnd()
	}
	return nil
}
