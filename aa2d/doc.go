// Copyright (c) 2016 Frank Braun <frank@cryptogroup.net>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package aa2d parses two-dimensional hierarchical ASCII art into an abstract
representation.
The abstract representation is a two-dimensional grid containing the top-level
elements with contained elements attached to them.


Elements

The following types of elements are supported.

Rectangles:

  #-----#     .-----.     #-----.
  |     |     |     |     |     |
  |     |     |     |     |     |
  #-----#     '-----'     '-----#

Rectangles can have 90 degree corners (#) or rounded corners (.) or (').
The two corner types can be mixed.

Lines:

  -------     |   ^   |   ^        /     ^
  ------>     |   |   |   |       /       \
  <------     |   |   |   |      /         \
  <----->     |   |   v   |     /           v

Lines can be horizontal, vertical, or diagonal and can have arrows
(<, >, ^, v) in one or two directions.

Dotted lines:

  =======     :   ^   :   ^
  ======>     :   :   :   :
  <======     :   :   :   :
  <=====>     :   :   v   :

Lines can also be dotted by using (=) instead of (-) horizontally or by using
(:) instead of (|) vertically. Diagonal lines can not be dotted.

Polylines:

  ------+
        |
        |     +----->
        |     |
        +-----+

A polyline is an open set of connected straight line segments (horizontal or
vertical). The segments must be connected with (+). The segments can be dotted.

Polygons:

     +---+
     |   |
 +---+   +---+
 |           |
 +---+   +---+
     |   |
     +---+

A polygon is a closed set of connected straight line segments (horizontal or
vertical). The segments must be connected with (+). The segments can be dotted.

Text lines:

  This is a text line.

A text line is a single line of text (without newlines).


References

All elements except text lines can have up to one reference attached to them
by providing it at the edge of the element and defining the reference
somewhere in the document at the start of a line (usually at the bottom).

Reference names are set between square brackets ([ and ]) and defined by
referencing the same name followed by a colon (:). After the colon a JSON object is expected.

Example:

  #-----#
  |[XB] |
  |     |
  #-----#

  [XB]: {fill: "none", stroke-width: 2}

Reference names starting with an underscore (_) are reserved. For example:

  [_ROOT]: {}
*/
package aa2d
