package main

import (
	"fmt"
)

func split(sum int) (b, a int) {
	a = sum * 4 / 9
	b = sum - a
	return
}

func swap(x, y int) (int, int) {
	return y, x
}
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func reverse(s string) string {
	if len(s) == 0 {
		return s
	}
	var result string = reverse(s[1:]) + s[:1]
	return result
}

func reverse2(s string) string {
	var result string = ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func gcm(a, b int) int {
	if b == 0 {
		return a
	}
	return gcm(b, a%b)
}
func main() {
	fmt.Println(split(98))
	fmt.Println(swap(10, 20))
	fmt.Println(swap(20, 10))
	fmt.Println(fibonacci(15))
	fmt.Println(factorial(5))
	fmt.Println(reverse("Hello, World!"))
	fmt.Println(reverse2("Hello, World!"))
	fmt.Println(gcm(10, 20))
	fmt.Println(gcm(20, 10))
}
