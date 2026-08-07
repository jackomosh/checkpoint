package main

import (
	"fmt"
	"strings"
	"unicode"
)

func longestWord(str string) string {

	if len(str) == 0 {
		return ""
	}

	words := strings.Fields(str)
	longword := ""

	for _, word := range words {
		cleaned := ""
		for _, ch := range word {
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
				cleaned += string(ch)
			}
		}
		if len(cleaned) > len(longword) {
			longword = cleaned
		}
	}
	return longword
}

func main() {
	test1 := "The Tiger and the Panda!!"
	test2 := "Check this: @apple, #banana, &cherry."
	test3 := "122 2314 25"

	// Using %v allows Go to automatically handle strings and integers safely
	fmt.Printf("Test 1 (Equal length): %v\n", longestWord(test1))
	fmt.Printf("Test 2 (Punctuation): %v\n", longestWord(test2))
	fmt.Printf("Test 3 (Numbers): %v\n", longestWord(test3))
}
