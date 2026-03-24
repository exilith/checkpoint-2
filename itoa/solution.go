// itoa

// Write a function that simulates the behavior of the Itoa function in Go. Itoa
// transforms a number represented as an int in a number represented as a string.
// For this exercise the handling of the signs + or - does have to be taken into account

package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var isNegative bool

	if n < 0 {
		isNegative = true
		n = -n
	}

	digits := []rune{}

	for n > 0 {
		d := n % 10
		digits = append(digits, rune('0'+d))
		n /= 10
	}

	if isNegative {
		digits = append(digits, '-')
	}

	// reverse slice
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
