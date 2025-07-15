package main

import "fmt"

func MaxWater(heights []int) int {
	left := 0
	right := len(heights) - 1
	water := 0
	for left < right {

		minHeight := min(heights[left], heights[right])
		curr := (right - left) * minHeight
		water = max(water, curr)

		if heights[left] > heights[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return water
}

func CheckMaxWater() {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		got := MaxWater(tc.heights)
		if got != tc.want {
			fmt.Printf("FAIL: maxWater(%v) = %d; want %d\n", tc.heights, got, tc.want)
		} else {
			fmt.Printf("PASS: maxWater(%v) = %d\n", tc.heights, got)
		}
	}
}
