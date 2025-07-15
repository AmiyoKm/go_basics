package main

import (
	"fmt"
	"math"
	"slices"
)

func FindMinDiff(arr []int, m int) int {
	slices.Sort(arr)
	res := math.MaxInt

	left := 0
	for right := m - 1; right < len(arr); right++ {
		curr := arr[right] - arr[left]
		res = min(res, curr)
		left++
	}
	return res
}

func CheckFindMinDiff() {
	tests := []struct {
		arr  []int
		m    int
		want int
	}{
		{[]int{12, 4, 7, 9, 2, 23, 25, 41, 30, 40, 28, 42, 30, 44, 48, 43, 50}, 7, 10},
		{[]int{3, 4, 1, 9, 56, 7, 9, 12}, 5, 6},
		{[]int{7, 3, 2, 4, 9, 12, 56}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 2, 1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
	}
	for _, tc := range tests {
		got := FindMinDiff(tc.arr, tc.m)
		if got != tc.want {
			fmt.Printf("FAIL: findMinDiff(%v, %d) = %d; want %d\n", tc.arr, tc.m, got, tc.want)
		} else {
			fmt.Printf("PASS: findMinDiff(%v, %d) = %d\n", tc.arr, tc.m, got)
		}
	}
}
