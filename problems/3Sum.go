package main

import (
	"fmt"
	"reflect"
	"slices"
)

func ThreeSum(arr []int) [][3]int {
	res := make([][3]int, 0)

	slices.Sort(arr)

	for i := 0; i < len(arr)-2; i++ {
		if i != 0 && arr[i-1] == arr[i] {
			continue
		}
		start := i + 1
		end := len(arr) - 1

		for start < end {
			summ := arr[i] + arr[start] + arr[end]

			if summ > 0 {
				end -= 1
			} else if summ < 0 {
				start += 1
			} else {
				res = append(res, [3]int{arr[i], arr[start], arr[end]})
				for start < len(arr)-1 && arr[start] == arr[start+1] {
					start++
				}
				start++
				end--
			}
		}

	}

	return res
}

func CheckThreeSum() {
	tests := []struct {
		arr  []int
		want [][3]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][3]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][3]int{}},
		{[]int{0, 0, 0}, [][3]int{{0, 0, 0}}},
		{[]int{}, [][3]int{}},
	}
	for _, tc := range tests {
		got := ThreeSum(tc.arr)
		if !EqualTriplets(got, tc.want) {
			fmt.Printf("FAIL: ThreeSum(%v) = %v; want %v\n", tc.arr, got, tc.want)
		} else {
			fmt.Printf("PASS: ThreeSum(%v) = %v\n", tc.arr, got)
		}
	}
}

func EqualTriplets(a, b [][3]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, tripA := range a {
		found := false
		for j, tripB := range b {
			if used[j] {
				continue
			}
			if reflect.DeepEqual(tripA, tripB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
