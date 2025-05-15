// util is a package containing various shared utilities in wowzers
package util

// Point is a struct storing a 2D size or coordinates
type Point struct {
	// Position
	X, Y int
}

// Rect is a struct containing 2 Points one for Position and one for Size
type Rect struct {
	// Position and Size
	Pos, Size Point
}
