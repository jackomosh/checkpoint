package main

import "fmt"

func RepeatAlpha(s string) string {

	result := ""
	count := 1

	for _, ch := range s {

		if ch >= 'A' && ch <= 'Z' {
			count = int(ch) - 'A' + 1
		} else if ch >= 'a' && ch <= 'z' {
			count = int(ch) - 'a' + 1
		} else {
			count = 1
		}

		for i := 0; i < count; i++ {
			result += string(ch)
		}
	}

	return result

}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
	fmt.Println(RepeatAlpha("c!"))
}
