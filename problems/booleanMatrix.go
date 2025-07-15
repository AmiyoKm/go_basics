package main

import (
	"fmt"
	"reflect"
)

func BooleanMatrix(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 1 {
				for i := range rows {
					if matrix[i][c] == 0 {
						matrix[i][c] = -1
					}
				}
				for i := range cols {
					if matrix[r][i] == 0 {
						matrix[r][i] = -1
					}
				}
			}
		}
	}

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == -1 {
				matrix[r][c] = 1
			}
		}
	}
}

func CheckBooleanMatrix() {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}},
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
		},
		{
			[][]int{{0, 0}, {0, 1}},
			[][]int{{0, 1}, {1, 1}},
		},
		{
			[][]int{{0}},
			[][]int{{0}},
		},
	}
	for _, tc := range tests {
		input := make([][]int, len(tc.matrix))
		for i := range tc.matrix {
			input[i] = make([]int, len(tc.matrix[i]))
			copy(input[i], tc.matrix[i])
		}
		BooleanMatrix(input)
		if !reflect.DeepEqual(input, tc.want) {
			fmt.Printf("FAIL: booleanMatrix(%v) = %v; want %v\n", tc.matrix, input, tc.want)
		} else {
			fmt.Printf("PASS: booleanMatrix(%v) = %v\n", tc.matrix, input)
		}
	}
}
