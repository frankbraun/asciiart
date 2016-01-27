// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"encoding/json"
	"regexp"
	"strings"
)

var matchReference = regexp.MustCompile(`^\[(.+)\]: (.*)$`)

func (p *Parser) parseReference(g *Grid, lines [][]byte, startY int) error {
	m := matchReference.FindStringSubmatch(string(lines[startY]))
	if m == nil {
		return &ParseError{X: 0, Y: startY, Err: ErrRefParseError}
	}
	key := m[1]
	jsn := m[2]
	var ref interface{}
	if err := json.Unmarshal([]byte(jsn), &ref); err != nil {
		return err
	}
	refMap, ok := ref.(map[string]interface{})
	if !ok {
		return &ParseError{X: 0, Y: startY, Err: ErrRefJSONObj}
	}
	if strings.HasPrefix(key, "_") {
		if g.Refs == nil {
			g.Refs = make(map[string]map[string]interface{})
		}
		if g.Refs[key] != nil {
			return &ParseError{X: 0, Y: startY, Err: ErrRefTwice}
		}
		g.Refs[key] = refMap
	} else {
		if p.refs[key] != nil {
			return &ParseError{X: 0, Y: startY, Err: ErrRefTwice}
		}
		p.refs[key] = refMap
	}
	return nil
}
