// retainfirsthalf

// Write a function called RetainFirstHalf() that takes a string as an argument
// and returns the first half of this string
// - if the length of the string is odd, round it down.
// - if the string is empty, return an empty string.
// - if the string length is equal to one, return the string

package piscine

func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

func RetainFirstHalf(str string) string {
	strLen := len(str)

	if strLen == 0 || strLen == 1 {
		return str
	}

	if isOdd(strLen) {
		strLen--
	}

	slice := []rune(str)
	return string(slice[:strLen/2])
}
