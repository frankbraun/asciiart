// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package svg

import (
	"fmt"
	"sort"
)

// style defines the default SVG style.
var style = map[string]map[string]string{
	"rect": {
		"fill":         "none",
		"stroke":       "black",
		"stroke-width": "2",
	},
}

func rectStyle() []string {
	var keys []string
	for k := range style["rect"] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	s := make([]string, len(keys))
	for i, k := range keys {
		s[i] = fmt.Sprintf("%s=%q", k, style["rect"][k])
	}
	return s
}
