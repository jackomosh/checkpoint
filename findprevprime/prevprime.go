package main

import "fmt"

func FindPrevPrime(nb int) int {

	for nb >= 2 {
		if isPrime(nb) {
			return nb
		}
		nb--
	}
	return 0
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
