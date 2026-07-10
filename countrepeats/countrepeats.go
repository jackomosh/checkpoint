// package main

// import "fmt"

// func CountRepeats(s string) string {
// 	if s == "" {
// 		return s
// 	}

// 	runes := []rune(s)
// 	result := ""
// 	counter := 1

// 	for i := 0; i < len(runes); i++ {
// 		// Increment count if the next character matches
// 		if i+1 < len(runes) && runes[i] == runes[i+1] {
// 			counter++
// 			continue
// 		}

// 		// Append character directly
// 		result += string(runes[i])

// 		// Append count using fmt if character repeated
// 		if counter > 1 {
// 			result += fmt.Sprintf("%d", counter)
// 		}
// 		counter = 1
// 	}

// 	return result
// }

// func main() {
// 	fmt.Println(CountRepeats("ABCABC"))       // Output: ABCABC
// 	fmt.Println(CountRepeats("AAABBC"))       // Output: A3B2C
// 	fmt.Println(CountRepeats("JjjJohhnnnNn")) // Output: Jj3Joh2n3Nn
// 	fmt.Println(CountRepeats("     ")) 
// 	fmt.Println(CountRepeats("  Jaaccccckk   "))         // Output:  5
// }


package main

import "fmt"

func CountRepeats(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	result := ""
	counter := 1
	for i := 0; i < len(runes); i++ {
		if i + 1 < len(runes) && runes[i] == runes[i+1] {
			counter++
			continue
		}
		result += string(runes[i])
		if counter > 0{
			result += fmt.Sprintf("%d", counter) 
		}
		counter = 1
	}
	return result
}

func main() {
	fmt.Println(CountRepeats("ABCABC"))       // Output: ABCABC
	fmt.Println(CountRepeats("AAABBC"))       // Output: A3B2C
	fmt.Println(CountRepeats("JjjJohhnnnNn")) // Output: Jj3Joh2n3Nn
	fmt.Println(CountRepeats("     ")) 
	fmt.Println(CountRepeats("  Jaaccccckk   "))         // Output:  5
}