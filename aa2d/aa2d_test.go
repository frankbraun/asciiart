// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"fmt"
	"reflect"
	"testing"
)

/* smallest possible rectangle
#-#
| |
#-#
*/

const rectangle = `
#-----#
|     |
|     |
#-----#
`

func TestRectangle(t *testing.T) {
	p := NewParser()
	p.SetScale(1, 1)
	g, err := p.Parse(rectangle)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", g)
	r := &Grid{W: 7, H: 6}
	if !reflect.DeepEqual(g, r) {
		t.Error("wrong grid")
	}
}
