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
		// global reference -> attach to grid
		if g.Refs == nil {
			g.Refs = make(map[string]map[string]interface{})
		}
		if g.Refs[key] != nil {
			return &ParseError{X: 0, Y: startY, Err: ErrRefTwice}
		}
		g.Refs[key] = refMap
	} else {
		// local reference -> store with parser for further processing
		if p.refs[key] != nil {
			return &ParseError{X: 0, Y: startY, Err: ErrRefTwice}
		}
		p.refs[key] = &reference{Y: startY, ref: refMap}
	}
	return nil
}
