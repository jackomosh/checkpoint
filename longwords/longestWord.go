package main

import (
	"fmt"
)

func longestWord(str string) string {

	longest := ""
	current := ""

	for i := 0; i < len(str); i++ {

		// set characters at the str[i]
		ch := str[i]

		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			current += string(ch)
		} else {
			if len(current) > len(longest) {
				longest = current
			}
			current = ""
		}
	}

	if len(current) > len(longest) {
		longest = current
	}

	return longest
}

func main() {
	test1 := "The Tiger and the Pandaaa"
	test2 := "Check this: apple, banana, cherries."
	test3 := "122 2314 25"

	// Using %v allows Go to automatically handle strings and integers safely
	fmt.Printf("Test 1 (Equal length): %v\n", longestWord(test1))
	fmt.Printf("Test 2 (Punctuation): %v\n", longestWord(test2))
	fmt.Printf("Test 3 (Numbers): %v\n", longestWord(test3))
}
