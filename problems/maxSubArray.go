package main

import (
	"fmt"
	"math"
)

func MaxSubArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxx := math.MinInt
	curr := 0

	for _, a := range arr {
		curr += a
		if curr > maxx {
			maxx = curr
		}
		if curr < 0 {
			curr = 0
		}
	}
	return maxx
}

func CheckMaxSubArray() {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 3, -8, 7, -1, 2, 3}, 11},
		{[]int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
		{[]int{-1, -2, -3, -4}, -1},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		got := MaxSubArray(tc.arr)
		if got != tc.want {
			fmt.Printf("FAIL: maxSubArray(%v) = %d; want %d\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: maxSubArray(%v) = %d\n", tc.arr, got)
		}
	}
}
