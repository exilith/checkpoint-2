// rectperimeter

// Write a function that takes two int as arguments, representing the
// length of width and the height of a rectangle and returning the
// perimeter of the rectangle
// - if one of the arguments is negative it should return -1

package piscine

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}

	// P = 2(l + w)
	return 2 * (w + h)
}
