package main

import "fmt"

func FindMin(arr []int) int {
	start, end := 0, len(arr)-1

	for start < end {
		if arr[start] < arr[end] {
			return arr[start]
		}

		mid := (start + end) / 2

		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return arr[start]
}

func CheckFindMin() {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{3, 1, 4, 1, 5, 9}, 1},
		{[]int{10, 20, 30}, 10},
		{[]int{-1, -5, 0, 2}, -5},
		{[]int{}, 0},
		{[]int{7}, 7},
	}
	for _, tc := range tests {
		got := FindMin(tc.arr)
		if got != tc.want {
			fmt.Printf("FAIL: findMin(%v) = %d; want %d\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: findMin(%v) = %d\n", tc.arr, got)
		}
	}
}
