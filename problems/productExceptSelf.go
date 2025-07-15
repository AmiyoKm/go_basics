package main

import (
	"fmt"
	"reflect"
)

func ProductExceptSelf(arr []int) []int {
	res := make([]int, len(arr))
	for i := range res {
		res[i] = 1
	}

	prefix := 1
	for i := range arr {
		res[i] = prefix
		prefix *= arr[i]
	}
	postfix := 1
	for i := len(arr) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= arr[i]
	}
	return res
}

func CheckProductExceptSelf() {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{2, 3, 4, 5}, []int{60, 40, 30, 24}},
		{[]int{1, 0, 3, 4}, []int{0, 12, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{5}, []int{1}},
	}
	for _, tc := range tests {
		got := ProductExceptSelf(tc.arr)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAIL: productExceptSelf(%v) = %v; want %v\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: productExceptSelf(%v) = %v\n", tc.arr, got)
		}
	}
}
