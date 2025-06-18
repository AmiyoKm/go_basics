package main

import (
	"fmt"

	"example.com/quick"
)

func bubble(arr []int, r, c int) (int int) {
	if r == 0 {
		return
	}
	if r > c {
		if arr[c] > arr[c+1] {
			arr[c], arr[c+1] = arr[c+1], arr[c]
		}
		bubble(arr, r, c+1)
	} else {
		bubble(arr, r-1, 0)
	}
	return
}

func selection(arr []int) []int {
	for i := range arr {
		last := len(arr) - 1 - i
		max := 0
		for j := max; j <= last; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[max], arr[last] = arr[last], arr[max]
	}
	return arr
}

func selectionRecursive(arr []int, r, c, max int) ([]int, int, int, int) {
	if r == 0 {
		return arr, r, c, max
	}
	if r > c {
		if arr[c] > arr[max] {
			selectionRecursive(arr, r, c+1, c)
		} else {
			selectionRecursive(arr, r, c+1, max)
		}
	} else {
		arr[max], arr[r-1] = arr[r-1], arr[max]
		selectionRecursive(arr, r-1, 0, 0)
	}
	return arr, r, c, max
}

func main() {
	// arr := []int{5, 4, 2, 3, 1}
	// bubble(arr, len(arr)-1, 0)
	// fmt.Println(arr)

	// arr2 := []int{5, 4, 2, 3, 1}
	// fmt.Println(selection(arr2))

	// arr3 := []int{5, 4, 2, 3, 1}
	// selectionRecursive(arr3, len(arr3), 0, 0)
	// fmt.Println(arr3)

	// arr4 := []int{5, 4, 2, 3, 1}
	// merge.MergeSortInPlace(arr4, 0, len(arr4))
	// fmt.Println(arr4)

	arr5 := []int{5, 4, 2, 3, 1}
	quick.QuickSort(arr5, 0, len(arr5)-1)
	fmt.Println(arr5)
}
