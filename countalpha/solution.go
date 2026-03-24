// countalpha

// Write a function CountAlpha() that takes a string as an argument and
// returns the number of alphabetic characters in the string

package piscine

func isAlpha(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
		return true
	}
	return false
}

func CountAlpha(s string) int {
	if len(s) == 0 {
		// empty string
		return 0
	}

	slice := []rune(s)
	var count int

	for _, c := range slice {
		if isAlpha(c) {
			count++
		}
	}

	return count
}
