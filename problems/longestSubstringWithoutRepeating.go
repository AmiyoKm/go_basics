package main

import (
	"fmt"
)

func LongestUniqueSubStr(s string) int {
	res := 0

	lastSeen := make(map[rune]int)
	left := 0

	for right , char := range s {

		if i , ok := lastSeen[char] ; ok && i >= left {
			left = right + 1
		}

		lastSeen[char]=right
		res = max(res , right-left+1)
	}

	return res
}

func CheckLongestUniqueSubStr() {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"abba", 2},
	}
	for _, tc := range tests {
		got := LongestUniqueSubStr(tc.s)
		if got != tc.want {
			fmt.Printf("FAIL: LongestUniqueSubStr(%q) = %d; want %d\n", tc.s, got, tc.want)
		} else {
			fmt.Printf("PASS: LongestUniqueSubStr(%q) = %d\n", tc.s, got)
		}
	}
}
