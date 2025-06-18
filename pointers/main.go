package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changePerson(p *Person, n string, a int) {
	p.name = n
	p.age = a
}

func changePersonValue(p Person, n string, a int) Person {
	p.name = n
	p.age = a
	return p
}

func main() {
	x := 10
	y := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", y)

	person1 := Person{name: "John", age: 25}

	changePerson(&person1, "Doe", 32)
	fmt.Println(person1)
	fmt.Println("Address of Name:", &person1.name)
	fmt.Println("Address of Age:", &person1.age)

	changePersonValue(person1, "Dooki", 35)
	fmt.Println(person1)

	arr := make([]int, 100000)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println("Address of arr:", &arr)
	for i := range arr {
		fmt.Printf("Address of arr[%d]: %p , value = %d\n", i, &arr[i], *&arr[i])
	}
	person1.age = 35
	fmt.Println(person1)
}
