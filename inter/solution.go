// inter

// Write a program that takes two string and displays, without doubles, the characters
// that appear in both string, in the order they appear in the first one
// - The display will be followed by a newline ('\n').
// - if the number of arguments is different from 2, the program displays nothing

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func contains(a []rune, char rune) bool {
	for _, c := range a {
		if c == char {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	str1 := []rune(args[0])
	str2 := []rune(args[1])

	doubles := []rune{}

	for i := range str1 {
		for j := range str2 {
			if str1[i] == str2[j] && !contains(doubles, str1[i]) {
				doubles = append(doubles, str1[i])
			}
		}
	}

	for _, c := range doubles {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
