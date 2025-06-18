package main

import "fmt"

const a = 20

var b = 30

func outer() func(x int) {
	money := 1000
	age := 23

	fmt.Println("Age", age)
	closure := func(x int) {
		sum := money + age + a + b + x
		fmt.Println("Sum", sum)
	}
	return closure
}
func adder() func(x int) int {
	sum := 0
	closure := func(x int) int {
		sum += x
		return sum
	}
	return closure
}

func call() {
	inner1 := outer()
	inner1(10)
	inner1(20)

	inner2 := outer()
	inner2(20)
	inner2(30)
}

func main() {
	//call()
	pos, neg := adder(), adder()

	for i := range 10 {
		fmt.Print(pos(i), " ")
		fmt.Println(neg(2 * -i))
	}
}

func init() {
	fmt.Println("====BANK====")
}
