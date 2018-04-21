/*
Package asciiart parses hierarchical ASCII art into an abstract representation.
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
Rectangles must be at least three times three cells large.

Lines:

  -------     |   ^   |   ^        /     ^
  ------>     |   |   |   |       /       \
  <------     |   |   |   |      /         \
  <----->     |   |   v   |     /           v

Lines can be horizontal, vertical, or diagonal and can have arrows
(<, >, ^, v) in one or two directions.
Lines must be at least two cells long.

Dotted lines:

  =======     :   ^   :   ^
  ======>     :   :   :   :
  <======     :   :   :   :
  <=====>     :   :   v   :

Lines can also be dotted by using (=) instead of (-) horizontally or by using
(:) instead of (|) vertically.
A line becomes dottes as soon as one cell on the line is dotted.
Diagonal lines cannot be dotted.

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

Rectangles and polygons can have up to one reference attached to them by
providing it at the edge of the element and defining the reference at the end
of the document.

Reference names are set between square brackets ([ and ]) and defined by
referencing the same name followed by a colon (:) at the end of the document.
After the colon a JSON object is expected.

Example:

  #-----#
  |[XB] |
  |     |
  #-----#

  [XB]: {fill: "none", stroke-width: 2}

All reference names starting with an underscore (_) are attached to the top-
level element (the grid). For example:

  [_SVG]: {"blur": false}
*/
package asciiart
