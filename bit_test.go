// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package asciiart

import (
	"testing"
)

func TestBitCount(t *testing.T) {
	if bitCount(0) != 0 {
		t.Error("bitCount(0) should be 0")
	}
	if bitCount(1) != 1 {
		t.Error("bitCount(1) should be 1")
	}
	if bitCount(7) != 3 {
		t.Error("bitCount(7) should be 3")
	}
	if bitCount(8) != 1 {
		t.Error("bitCount(8) should be 1")
	}
	if bitCount(255) != 8 {
		t.Error("bitCount(255) should be 8")
	}
}
