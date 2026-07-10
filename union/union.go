package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Invalid Operation Provide atleast two args")
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]
	combined := str1 + str2

	seen := make(map[rune]bool)
	var newString []rune

	for _, ch := range combined {
		if !seen[ch] {
			seen[ch] = true
			newString = append(newString, ch)
		}
	}
	fmt.Println(string(newString))
}