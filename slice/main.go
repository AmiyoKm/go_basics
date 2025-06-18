package main

import "fmt"

func main() {

	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// arr1 := arr[1:5]
	// fmt.Println(arr1)
	// fmt.Println(len(arr1))
	// fmt.Println(cap(arr1))

	// arr2 := arr1[2:3]
	// fmt.Println(arr2)
	// fmt.Println(len(arr2))
	// fmt.Println(cap(arr2))

	// arr3 := make([]int, 5, 10)
	// arr3 = append(arr3, 1, 2, 3, 4, 5)
	// fmt.Println(len(arr3))
	// fmt.Println(cap(arr3))

	// var arr4 []int
	// fmt.Println(arr4)

	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 5)
	y = append(y, 4)

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)
	// arr := []int{}
	// for i := range 600 {
	// 	arr = append(arr, i)
	// 	arr[i] = i
	// 	fmt.Println(arr[i])
	// 	fmt.Println("array length:", len(arr))
	// 	fmt.Println("array capacity:", cap(arr))
	// }
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr1[0:1]
	fmt.Println(s, cap(s), len(s))
	s = s[1:5]
	fmt.Println(s, cap(s), len(s))

	s = s[:7]
	fmt.Println(s, cap(s), len(s))

	s = s[4:8]
	fmt.Println(s, cap(s), len(s))

	s = s[:2]
	fmt.Println(s, cap(s), len(s))
}
