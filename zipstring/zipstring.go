package main

func ZipString(s string) string {

	if s == "" {
		return s
	}

	runes := []rune(s)
	result := ""
	count := 1

	for i := 1; i <= len(runes); i++ {
		if i < len(runes) && runes[i] == runes[i-1] {
			count++
		} else {
			result += intToString(count) + string(runes[i-1])
			count = 1
		}
	}
	return result
}

func intToString(n int) string {

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
	println(ZipString("YouuungFellllas"))
	println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	println(ZipString("Helloo Therre!"))
	println(ZipString("Haaaaaaaaaa"))
}
