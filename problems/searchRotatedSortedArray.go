package main

import "fmt"

func Search(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}
	first := FindPivot(arr)

	if first == 0 {
		return BinarySearch(arr, 0, len(arr)-1, target)
	}

	if arr[0] <= target {
		return BinarySearch(arr, 0, first-1, target)
	}
	return BinarySearch(arr, first, len(arr)-1, target)
}

func FindPivot(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {

		if arr[start] < arr[end] {
			return start
		}

		mid := (start + end) / 2

		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func BinarySearch(arr []int, start, end, target int) int {
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func CheckSearchRotatedSortedArray() {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 0, -1},
		{[]int{1, 3}, 3, 1},
		{[]int{5, 1, 3}, 3, 2},
	}
	for _, tc := range tests {
		got := Search(tc.arr, tc.target)
		if got != tc.want {
			fmt.Printf("FAIL: search(%v, %d) = %d; want %d\n", tc.arr, tc.target, got, tc.want)
		} else {
			fmt.Printf("PASS: search(%v, %d) = %d\n", tc.arr, tc.target, got)
		}
	}
}
