package main

import (
	"fmt"
)

func ZipString(s string) string {

	if s == "" {
		return s
	}

	runes := []rune(s)
	result := ""
	count := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, runes[i-1])
			count = 1
		}
	}
	result += fmt.Sprintf("%d%c", count, runes[len(runes)-1])
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
