package asciiart

import (
	"errors"
	"fmt"
)

// Possible scaling errors.
var (
	// ErrWrongXScale is returned when an illegal xScale is set.
	ErrWrongXScale = errors.New("asciiart: xScale must be at least one")
	// ErrWrongYScale is returned when an illegal yScale is set.
	ErrWrongYScale = errors.New("asciiart: yScale must be at least one")
)

// Possible parse errors.
var (
	// ErrExpRecLine is returned when a rectangle line was expected.
	ErrExpRecLine = errors.New("asciiart: expected rectangle line (-)")
	// ErrExpRecLineOrUpCorn is returned when a rectangle line or upper corner was expected.
	ErrExpRecLineOrUpCorn = errors.New("asciiart: expected rectangle line (-) or upper corner (#, .)")
	// ErrExpRecLineOrLowCorn is returned when a rectangle line or lower corner was expected.
	ErrExpRecLineOrLowCorn = errors.New("asciiart: expected rectangle line (|) or lower corner (#, ')")
	// ErrExpRecHorizontalLine is returned when a horizontal line was expected.
	ErrExpRecHorizontalLine = errors.New("asciiart: expected horizontal line (-)")
	// ErrExpRecVerticalLine is returned when a horizontal line was expected.
	ErrExpRecVerticalLine = errors.New("asciiart: expected vertical line (|)")
	// ErrExpRecLowCorn is returned when a lower corner was expected.
	ErrExpRecLowCorn = errors.New("asciiart: expected lower corner")
	// ErrNoRecUpRightCorn is returned when the upper right corner of a rectangle could not be found.
	ErrNoRecUpRightCorn = errors.New("asciiart: could not find upper right rectangle corner")
	// ErrNoRecLowRightCorn is returned when the lower right corner of a rectangle could not be found.
	ErrNoRecLowRightCorn = errors.New("asciiart: could not find lower right rectangle corner")
	// ErrNoRecLowLeftCorn is returned when the lower left corner of a rectangle could not be found.
	ErrNoRecLowLeftCorn = errors.New("asciiart: could not find lower left rectangle corner")
	// ErrRightArrow is returned when an right arrow is encountered that doesn't make sense.
	ErrRightArrow = errors.New("asciiart: right arrow doesn't make sense here")
	// ErrLeftArrow is returned when an left arrow is encountered that doesn't make sense.
	ErrLeftArrow = errors.New("asciiart: left arrow doesn't make sense here")
	// ErrRefParseError is returned when a reference could not be parsed.
	ErrRefParseError = errors.New("asciiart: cannot parse reference on line starting")
	// ErrRefJSONObj is returned when a reference definition is not a JSON object.
	ErrRefJSONObj = errors.New("asciiart: cannot parse reference JSON object on line starting")
	// ErrRefJSONUnmarshal is returned when an error occured during JSON unmarshalling.
	ErrRefJSONUnmarshal = errors.New("asciiart: cannot unmarshal JSON object")
	// ErrRefTwice is returned when a reference was defined twice.
	ErrRefTwice = errors.New("asciiart: reference defined twice on line starting")
	// ErrRefMiddle is returned when a reference was defined in the middle of the document.
	ErrRefMiddle = errors.New("asciiart: reference defined in the middle of document on line starting")
	// ErrRefMissingBracket is returned when the closing bracket of a reference is missing.
	ErrRefMissingBracket = errors.New("asciiart: closing bracket (]) of reference is missing")
	// ErrRefKeyEmpty is returned when a reference key is empty.
	ErrRefKeyEmpty = errors.New("asciiart: reference key is empty")
	// ErrRefKeyUndefined is returned when a reference key is undefined.
	ErrRefKeyUndefined = errors.New("asciiart: reference key is undefined")
	// ErrRefUnused is returned when a reference has been defined but wasn't used.
	ErrRefUnused = errors.New("ascii art: reference not used")
	// ErrLineTooShort is returned when a line is too short.
	ErrLineTooShort = errors.New("asciiart: line is too short")
	// ErrLineLeftArrow is returned when an unexpected left error was encountered.
	ErrLineLeftArrow = errors.New("asciiart: unexpected left error encountered")
	// ErrPolyCornerNoEdge is returned when a poly corner has no outgoing edges.
	ErrPolyCornerNoEdge = errors.New("asciiart: poly corner (+) has no outgoing edges")
	// ErrPolyCornerOneEdge is returned when a poly corner has only one edge.
	ErrPolyCornerOneEdge = errors.New("asciiart: poly corner (+) has only one edge")
	// ErrPolyCornerTooManyEdges is returned when a poly corner has too many edges.
	ErrPolyCornerTooManyEdges = errors.New("asciiart: poly corner (+) has too many edges")
	// ErrPolygonNotClosed is returned when polygon is not closed.
	ErrPolygonNotClosed = errors.New("asciiart: non-closed polygon starting")
)

// ParseError defines an ASCII art parsing error.
type ParseError struct {
	X   int   // x-axis coordinate where the error occurred
	Y   int   // y-axis coordinate where the error occurred
	Err error // the actual error
}

func (e *ParseError) Error() string {
	return e.Err.Error() + " " + fmt.Sprintf("at (%d,%d)", e.X, e.Y)
}
