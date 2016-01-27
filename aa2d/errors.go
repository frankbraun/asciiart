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
	// ErrNoRecUpRightCorn is returned when the upper right corner of a rectangle could not be found.
	ErrNoRecUpRightCorn = errors.New("aa2d: could not find upper right rectangle corner")
	// ErrNoRecLowRightCorn is returned when the lower right corner of a rectangle could not be found.
	ErrNoRecLowRightCorn = errors.New("aa2d: could not find lower right rectangle corner")
	// ErrNoRecLowLeftCorn is returned when the lower left corner of a rectangle could not be found.
	ErrNoRecLowLeftCorn = errors.New("aa2d: could not find lower left rectangle corner")
	// ErrUnknownCharacter is returned when an unknown character was encountered.
	ErrUnknownCharacter = errors.New("aa2d: unknown character encountered")
	// ErrRightArrow is returned when an right arrow is encountered that doesn't make sense.
	ErrRightArrow = errors.New("aa2d: right arrow doesn't make sense here")
	// ErrLeftArrow is returned when an left arrow is encountered that doesn't make sense.
	ErrLeftArrow = errors.New("aa2d: left arrow doesn't make sense here")
	// ErrRefParseError is returned when a reference could not be parsed.
	ErrRefParseError = errors.New("aa2d: cannot parse reference on line starting")
	// ErrRefJSONObj is returned when a reference definition is not a JSON object.
	ErrRefJSONObj = errors.New("aa2d: cannot parse reference JSON object on line starting")
	// ErrRefTwice is returned when a reference was defined twice.
	ErrRefTwice = errors.New("aa2d: reference defined twice on line starting")
	// ErrRefMiddle is returned when a reference was defined in the middle of the document.
	ErrRefMiddle = errors.New("aa2d: reference defined in the middle of document on line starting")
	// ErrLineTooShort is returned when a line is too short.
	ErrLineTooShort = errors.New("aa2d: line is too short")
	// ErrLineLeftArrow is returned when an unexpected left error was encountered.
	ErrLineLeftArrow = errors.New("aa2d: unexpected left error encountered")
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
