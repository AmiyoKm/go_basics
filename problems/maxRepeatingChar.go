package main

import "fmt"

func MaxRepeating(s string, k int) int {
	res := 0
	currMaxFreq := make(map[rune]int)
	currMaxChar := 0
	left := 0

	for right, char := range s {
		currMaxFreq[char]++
		// maxx := 0

		// for _, val := range currMaxFreq {
		// 	if val >= maxx {
		// 		maxx = val
		// 	}
		// }

		currMaxChar := max(currMaxChar, currMaxFreq[char])

		for (right-left+1)-currMaxChar > k {
			currMaxFreq[rune(s[left])]--
			left++
		}

		res = max(res, right-left+1)
	}
	return res
}

func CheckMaxRepeating() {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"AABABBA", 1, 4},
		{"ABAB", 2, 4},
		{"AAAA", 2, 4},
		{"ABCDE", 1, 2},
		{"AABACCA", 2, 5},
	}
	for _, tc := range tests {
		got := MaxRepeating(tc.s, tc.k)
		if got != tc.want {
			fmt.Printf("FAIL: MaxRepeating(%q, %d) = %d; want %d\n", tc.s, tc.k, got, tc.want)
		} else {
			fmt.Printf("PASS: MaxRepeating(%q, %d) = %d\n", tc.s, tc.k, got)
		}
	}
}
