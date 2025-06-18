package functions

import "fmt"

func Add(a int, b int) (sum int) {
	sum = a + b
	return
}

func Subtract(a int, b int) (sub int) {
	sub = a - b
	return
}

func Multiply(a int, b int) (mul int) {
	mul = a * b
	return
}
func Divide(a int, b int) (div int) {
	div = a / b
	return
}

func Modulus(a int, b int) (mod int) {
	mod = a % b
	return
}

func ProcessOperator(a int, b int, op func(x int, y int) int) int {
	return op(a, b)
}

func ProcessOperatorWithInput(x string) func(int, int) int {
	switch x {
	case "+":
		return Add
	case "-":
		return Subtract
	case "*":
		return Multiply
	case "/":
		return Divide
	case "%":
		return Modulus
	default:
		fmt.Println("Invalid operator, returning addition as default.")
		return Add
	}
}
