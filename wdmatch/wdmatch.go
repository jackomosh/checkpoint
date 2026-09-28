package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	runes := []rune(str1)
	count := 0

	// Iterate through str2 and match characters sequentially
	for _, ch := range str2 {
		if count < len(runes) && runes[count] == ch {
			count++
		}
	}

	// If we successfully matched every character in str1 in order
	if count == len(runes) {
		for _, ch := range runes {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}