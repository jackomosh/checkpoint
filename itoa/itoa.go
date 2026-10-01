package main

import "fmt"

func Itoa(n int) string {

	if n == 0 {
		return "0"
	}

	negative := false

	if n < 0 {
		negative = true
		n = -n
	}

	result := ""

	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n = n / 10
	}

	if negative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
