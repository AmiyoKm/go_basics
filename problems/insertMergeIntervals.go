package main

import (
	"fmt"
	"reflect"
)

func MergeInterval(intervals [][2]int, newInterval [2]int) [][2]int {
	res := make([][2]int, 0)

	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	res = append(res, intervals[i:]...)

	return res
}

func CheckMergeInterval() {
	tests := []struct {
		intervals   [][2]int
		newInterval [2]int
		want        [][2]int
	}{
		{
			[][2]int{{1, 3}, {6, 9}},
			[2]int{2, 5},
			[][2]int{{1, 5}, {6, 9}},
		},
		{
			[][2]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[2]int{4, 8},
			[][2]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			[][2]int{},
			[2]int{5, 7},
			[][2]int{{5, 7}},
		},
		{
			[][2]int{{1, 5}},
			[2]int{2, 3},
			[][2]int{{1, 5}},
		},
	}
	for _, tc := range tests {
		got := MergeInterval(tc.intervals, tc.newInterval)
		if !reflect.DeepEqual(got, tc.want) {
			fmt.Printf("FAIL: mergeInterval(%v, %v) = %v; want %v\n", tc.intervals, tc.newInterval, got, tc.want)
		} else {
			fmt.Printf("PASS: mergeInterval(%v, %v) = %v\n", tc.intervals, tc.newInterval, got)
		}
	}
}
