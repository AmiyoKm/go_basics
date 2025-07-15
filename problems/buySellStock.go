package main

import "fmt"

func MaxProfit(prices []int) int {
	maxx := 0
	left := 0

	for right := range prices {
		profit := prices[right] - prices[left]

		if profit > maxx {
			maxx = profit
		}
		if prices[right] < prices[left] {
			left = right
		}
	}
	return maxx
}
func CheckMaxProfit() {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{2, 4, 1}, 2},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		got := MaxProfit(tc.prices)
		if got != tc.want {
			fmt.Printf("FAIL: maxProfit(%v) = %d; want %d\n", tc.prices, got, tc.want)
		} else {
			fmt.Printf("PASS: maxProfit(%v) = %d\n", tc.prices, got)
		}
	}
}
