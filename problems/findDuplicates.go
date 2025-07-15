package main

import (
	"fmt"
	"reflect"
	"sort"
)

func FindDuplicates(arr []int) []int {
	mpp := make(map[int]int)

	for _, val := range arr {
		if v, ok := mpp[val]; ok {
			mpp[val] = v + 1
		} else {
			mpp[val] = 1
		}
	}
	res := make([]int, 0)
	for key, val := range mpp {
		if val > 1 {
			res = append(res, key)
		}
	}
	return res
}

func CheckFindDuplicates() {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 2, 2, 3, 4, 4, 5}, []int{2, 4}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{5, 4, 3, 2, 1}, []int{}},
	}
	for _, tc := range tests {
		got := FindDuplicates(tc.arr)
		sort.Ints(got)
		sort.Ints(tc.want)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAIL: findDuplicates(%v) = %v; want %v\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: findDuplicates(%v) = %v\n", tc.arr, got)
		}
	}
}
