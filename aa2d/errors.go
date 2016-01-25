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
	// ErrExpRecLineOrUpCorn is returned when a rectangle line or upper corner was expected.
	ErrExpRecLineOrUpCorn = errors.New("aa2d: expected rectangle line (-) or upper corner (#, .)")
	// ErrExpRecLineOrLowCorn is returned when a rectangle line or lower corner was expected.
	ErrExpRecLineOrLowCorn = errors.New("aa2d: expected rectangle line (|) or lower corner (#, ')")
	// ErrExpRecHorizontalLine is returned when a horizontal line was expected.
	ErrExpRecHorizontalLine = errors.New("aa2d: expected horizontal line (-)")
	// ErrExpRecVerticalLine is returned when a horizontal line was expected.
	ErrExpRecVerticalLine = errors.New("aa2d: expected vertical line (|)")
	// ErrExpRecLowCorn is returned when a lower corner was expected.
	ErrExpRecLowCorn = errors.New("aa2d: expected lower corner")
	// ErrNoRecUpRightCorn is return when the upper right corner of a rectangle could not be found.
	ErrNoRecUpRightCorn = errors.New("aa2d: could not find upper right rectangle corner")
	// ErrNoRecLowRightCorn is return when the lower right corner of a rectangle could not be found.
	ErrNoRecLowRightCorn = errors.New("aa2d: could not find lower right rectangle corner")
	// ErrNoRecLowLeftCorn is return when the lower left corner of a rectangle could not be found.
	ErrNoRecLowLeftCorn = errors.New("aa2d: could not find lower left rectangle corner")
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
