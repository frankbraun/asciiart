// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package aa2d

import (
	"errors"
	"fmt"
)

// Possible scaling errors.
var (
	// ErrWrongXScale is returned when an illegal xScale is set.
	ErrWrongXScale = errors.New("aa2d: xScale must be at least one")
	// ErrWrongYScale is returned when an illegal yScale is set.
	ErrWrongYScale = errors.New("aa2d: yScale must be at least one")
)

// Possible parse errors.
var (
	// ErrExpRecLine is returned when a rectangle line was expected.
	ErrExpRecLine = errors.New("aa2d: expected rectangle line (-)")
	// ErrExpRecLineOrCorn is return when a rectangle line or corner was expected.
	ErrExpRecLineOrCorn = errors.New("aa2d: expected rectangle line (-) or corner")
)

// ParseError defines a ASCII art parsing error.
type ParseError struct {
	X   int   // x-axis coordinate where the error occurred
	Y   int   // y-axis coordinate where the error occured
	Err error // the actual error
}

func (e *ParseError) Error() string {
	return e.Err.Error() + " " + fmt.Sprintf("at (%d,%d)", e.X, e.Y)
}
