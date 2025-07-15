package main

import (
	"fmt"
	"math"
)

func MaxSubProduct(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 0 {
		return 0
	}
	maxx := math.MinInt

	leftToRight := 1
	rightToLeft := 1

	for i := range arr {
		if leftToRight == 0 {
			leftToRight = 1
		}
		if rightToLeft == 0 {
			rightToLeft = 1
		}
		leftToRight *= arr[i]

		j := len(arr) - 1 - i
		rightToLeft *= arr[j]

		currMax := 0
		if rightToLeft >= leftToRight {
			currMax = rightToLeft
		} else {
			currMax = leftToRight
		}

		if maxx < currMax {
			maxx = currMax
		}
	}
	return maxx
}

func CheckMaxSubProduct() {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
		{[]int{0, 2}, 2},
		{[]int{-2}, -2},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		got := MaxSubProduct(tc.arr)
		if got != tc.want {
			fmt.Printf("FAIL: maxSubProduct(%v) = %d; want %d\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: maxSubProduct(%v) = %d\n", tc.arr, got)
		}
	}
}
