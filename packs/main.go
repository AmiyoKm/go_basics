package main

import (
	"fmt"

	"example.com/functions"
)

func main() {
	// fmt.Println("Addition: ", functions.ProcessOperator(10, 5, functions.Add))
	// fmt.Println("Subtraction: ", functions.Add(10, 5))
	// fmt.Println("Multiplication: ", functions.Multiply(10, 5))
	// fmt.Println("Division: ", functions.ProcessOperator(10, 5, functions.Divide))
	// fmt.Println("Modulus: ", functions.ProcessOperator(12, 5, functions.Modulus))
	var input string
	fmt.Println("Enter the operator: ")
	fmt.Scan(&input)
	if input != "+" && input != "-" && input != "*" && input != "/" && input != "%" {
		fmt.Println("Invalid operator")
		return
	}
	fmt.Println("Enter the first number: ")
	var a int
	fmt.Scan(&a)
	fmt.Println("Enter the second number: ")
	var b int
	fmt.Scan(&b)
	if input == "+" {
		add := functions.ProcessOperatorWithInput(input)
		fmt.Println("Addition: ", add(a, b))
	}
	if input == "-" {
		sub := functions.ProcessOperatorWithInput(input)
		fmt.Println("Subtraction: ", sub(a, b))
	}
	if input == "*" {
		mul := functions.ProcessOperatorWithInput(input)
		fmt.Println("Multiplication: ", mul(a, b))
	}
	if input == "/" {
		div := functions.ProcessOperatorWithInput(input)
		fmt.Println("Division: ", div(a, b))
	}
	if input == "%" {
		mod := functions.ProcessOperatorWithInput(input)
		fmt.Println("Modulus: ", mod(a, b))
	}
}
