// slice

// The function receives a slice of strings and one or more integers, and returns a
// slice of strings. The returned slice is part of the received one but cut from the
// position indicated in the first int, until the position indicated by the second int
// In case there only exists one int, the result slice begins in the position indicated
// by the int and ends at the end of the received slice.
// The integers can be negative.

package piscine

// l - length of slice
// n - possible negative index
func parseIdx(n int, l int) int {
	if n < 0 {
		return l + n
	}
	return n
}

func isValidIdx(start, stop int) bool {
	if start > stop || stop < start {
		return false
	}
	return true
}

func Slice(a []string, nbrs ...int) []string {
	argLen := len(a)

	if len(nbrs) == 1 {
		// starting point was provided
		return a[parseIdx(nbrs[0], argLen):]
	}

	if len(nbrs) == 2 && isValidIdx(nbrs[0], nbrs[1]) {
		return a[parseIdx(nbrs[0], argLen):parseIdx(nbrs[1], argLen)]
	}

	return []string(nil)
}
