package main

import "fmt"

func modify(arr []int) {
	fmt.Printf("address %p\n",arr)
	arr[0] = 1
	arr = []int{3,4,5}
	fmt.Printf("address %p\n",arr)
	fmt.Print(arr)
}
func main() {
	arr := []int{2, 3, 4}
	fmt.Printf("address %p\n",arr)

	modify(arr)
	fmt.Print(arr)
	fmt.Printf("address %p\n",arr)

}
