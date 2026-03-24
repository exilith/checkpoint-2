// firstword

// Write a function that takes a string and return a string containing its first word,
// followed by a newline ('\n).
// - A word is sequence of characters delimited by spaces or by start/end of the argument

package piscine

func FirstWord(s string) string {
	slice := []rune(s)
	result := []rune{}

	for _, c := range slice {
		if c == ' ' {
			result = append(result, '\n')
			break
		}
		result = append(result, c)
	}

	return string(result)
}
