package main

import "fmt"

func calculation() (result int) {
	fmt.Println("1st ", result)

	show := func() {
		result += 10
		fmt.Println("defer ", result)
	}
	defer show()
	result = 10
	fmt.Println("2nd", result)
	return
}

func calc() int {
	result := 0
	fmt.Println("1st ", result)

	show := func() {
		result += 10
		fmt.Println("defer ", result)
	}
	defer show()
	result = 10
	fmt.Println("2nd", result)
	return result
}
func final() int {
	result := 0
	fmt.Println("1st ", result)

	show := func() {
		result += 10
		fmt.Println("defer ", result)
	}
	defer show()

	result = 10
	defer fmt.Println("2nd", result)
	fmt.Println("3rd", result)

	defer func(result int) {
		fmt.Println("4th ", result)
	}(result)

	result++
	return result
}
func main() {
	fmt.Println(calculation())
	fmt.Println(calc())
	fmt.Println(final())
}
