package main

import "fmt"

func FirstWord(s string) string {
    if len(s) == 0 {
		return "\n"
	}
	// runes := []rune(s)
	// i := 0
	// result := ""
	// if runes[0] != ' ' && runes[i-1] == ' ' {
	// 	result = append(runes, result)
	// }
	result := ""
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := start
	for end < len(s) && s[end] != ' ' {
		end++
		result += string()
	}
	return s[start:end] + "\n"
}

func main() {
    fmt.Print(FirstWord("hello, there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("   hello   .........  bye"))
}