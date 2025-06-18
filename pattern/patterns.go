package main

import "fmt"

func triangle(n int) {
	for i := range n {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func triangleRecursive(n int) {
	if n == 0 {
		return
	}
	for range n {
		fmt.Print("* ")
	}
	fmt.Println()
	triangleRecursive(n - 1)

}
func triangleRecursive2(n int) {
	if n == 0 {
		return
	}
	triangleRecursive2(n - 1)
	for range n {
		fmt.Print("* ")
	}
	fmt.Println()
}

func triangleRecursive3(r, c int) {
	if r == 0 {
		return
	}
	if r > c {
		triangleRecursive3(r, c+1)
		fmt.Print("* ")
	} else {
		triangleRecursive3(r-1, 0)
		fmt.Println()
	}

}

func square(n int) {
	for range n {
		for range n {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
func main() {
	triangle(5)
	fmt.Println()
	square(5)
	fmt.Println()
	triangleRecursive(5)
	fmt.Println()
	triangleRecursive2(5)
	fmt.Println()
	triangleRecursive3(5, 0)
	fmt.Println()
}
