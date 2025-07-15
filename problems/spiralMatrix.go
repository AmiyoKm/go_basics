package main

import (
	"fmt"
	"reflect"
)

func SpiralMatrix(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows := len(matrix)
	cols := len(matrix[0])

	top := 0
	bottom := rows
	left := 0
	right := cols

	res := make([]int, 0)
	for left < right && top < bottom {

		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right-1])
		}
		right--

		if left >= right || top >= bottom {
			break
		}

		for i := right-1; i >= left; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--

		for i := bottom-1; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func CheckSpiralMatrix() {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{{1}},
			[]int{1},
		},
		{
			[][]int{},
			[]int{},
		},
	}
	for _, tc := range tests {
		got := SpiralMatrix(tc.matrix)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAIL: SpiralMatrix(%v) = %v; want %v\n", tc.matrix, got, tc.want)
		} else {
			fmt.Printf("PASS: SpiralMatrix(%v) = %v\n", tc.matrix, got)
		}
	}
}
