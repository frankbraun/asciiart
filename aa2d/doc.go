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

  -------     | ^ | ^        /     ^
  ------>     | | | |       /       \
  <------     | | | |      /         \
  <----->     | | v |     /           v

Lines can be horizontal, vertical, or diagonal and can have arrows
(<, >, ^, v) in one or two directions.




References

Reference names starting with an underscore (_) are reserved.
*/
package aa2d
