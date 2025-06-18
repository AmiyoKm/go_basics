package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World! 🐼")
	fmt.Printf("Hello, World! 🐼\n")
	var ten int = 10
	fmt.Println("Value of var is", ten)

	different := "THis is different 🐲🐲"
	fmt.Println(different)
	const nah = "This is a constant"
	var nah2 = "This is a variable"
	nah2 = "This is a variable2"
	fmt.Println(nah2)

	something := map[string]string{
		"first":  "first",
		"second": "second",
	}
	something["third"] = "third"
	fmt.Println(something)

}
