package main

import (
	"fmt"
	"reflect"
)

func TransPoseMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		res[i] = make([]int, len(matrix))
	}
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			res[c][r] = matrix[r][c]
		}
	}
	return res
}

func CheckTransPoseMatrix() {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{},
			[][]int{},
		},
	}
	for _, tc := range tests {
		got := TransPoseMatrix(tc.matrix)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAIL: transPoseMatrix(%v) = %v; want %v\n", tc.matrix, got, tc.want)
		} else {
			fmt.Printf("PASS: transPoseMatrix(%v) = %v\n", tc.matrix, got)
		}
	}
}
