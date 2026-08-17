package main

import "fmt"

func CountRepeats(s string) string {

	if s == "" {
		return s
	}

	runes := []rune(s)
	result := ""
	count := 1

	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			count++
			continue
		}
		result += string(runes[i])

		if count > 1 {
			result += intToString(count)
		}
	}
	return result
}

func intToString(n int) string{

	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n = n / 10
	}
	return result
}

func main() {
	fmt.Println(CountRepeats("ABCABC"))       // Output: ABCABC
	fmt.Println(CountRepeats("AAABBC"))       // Output: A3B2C
	fmt.Println(CountRepeats("JjjJohhnnnNn")) // Output: Jj3Joh2n3Nn
	fmt.Println(CountRepeats("     "))
	fmt.Println(CountRepeats("  Jaaccccckk   ")) // Output:  5
}
