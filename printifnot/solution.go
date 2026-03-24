// printifnot

// Write a function that takes a string as an argument and returns
// the letter "G" if the argument length is less than 3, otherwise
// returns "Invalid Input" followed by a newline "\n".
// - if it's an empty string return "G" followed by a new line \n

package piscine

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {}
