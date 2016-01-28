// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

import (
	"encoding/json"
	"regexp"
	"strings"
)

var matchReference = regexp.MustCompile(`^\[(.+)\]:([ \t]+)(.*)$`)

func (p *Parser) parseReference(g *Grid, lines [][]byte, startY int) error {
	m := matchReference.FindStringSubmatch(string(lines[startY]))
	if m == nil {
		return &ParseError{X: 0, Y: startY, Err: ErrRefParseError}
	}
	key := m[1]
	mid := m[2]
	jsn := m[3]
	var ref interface{}
	if err := json.Unmarshal([]byte(jsn), &ref); err != nil {
		x := len(key) + 2 + len(mid)
		se, ok := err.(*json.SyntaxError)
		if ok {
			x += int(se.Offset)
		}
		return &ParseError{X: x, Y: startY, Err: ErrRefJSONUnmarshal}
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
