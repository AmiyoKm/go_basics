package recursion

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
)

func RemoveA(s string) string {
	if len(s) == 0 {
		return ""
	}

	ch := rune(s[0])

	if ch == 'a' || ch == 'A' {
		return RemoveA(s[1:])
	} else {
		return string(ch) + RemoveA(s[1:])
	}
}

func Subsequence(processed string, unprocessed string) {
	if len(unprocessed) == 0 {
		fmt.Println(processed)
		return
	}

	ch := unprocessed[0]

	Subsequence(processed+string(ch), unprocessed[1:])
	Subsequence(processed, unprocessed[1:])
}

func SubsequenceReturn(processed string, unprocessed string) []string {
	if len(unprocessed) == 0 {
		return []string{processed}
	}

	ch := unprocessed[0]

	left := SubsequenceReturn(processed+string(ch), unprocessed[1:])
	right := SubsequenceReturn(processed, unprocessed[1:])

	left = append(left, right...)
	return left
}

func SubsequenceReturnASCII(processed string, unprocessed string) []string {
	if len(unprocessed) == 0 {
		return []string{processed}
	}

	ch := unprocessed[0]

	left := SubsequenceReturnASCII(processed+string(ch), unprocessed[1:])
	middle := SubsequenceReturnASCII(processed+strconv.Itoa(int(ch)), unprocessed[1:])
	right := SubsequenceReturnASCII(processed, unprocessed[1:])

	right = append(right, middle...)
	left = append(left, right...)
	return left
}

func Subset(nums []int) [][]int {
	outer := make([][]int, 0)

	outer = append(outer, []int{})

	for _, num := range nums {
		n := len(outer)

		for i := range n {
			inter := make([]int, 0)
			inter = append(inter, outer[i]...)
			inter = append(inter, num)
			outer = append(outer, inter)
		}
	}
	return outer
}
func SubsetDuplicate(nums []int) [][]int {
	outer := make([][]int, 0)

	outer = append(outer, []int{})
	start := 0
	end := 0
	for index, num := range nums {
		n := len(outer)
		start = 0
		if index > 0 && nums[index] == nums[index-1] {
			start = end + 1
		}
		end = n - 1
		for i := start; i < n; i++ {
			inter := make([]int, 0)
			inter = append(inter, outer[i]...)
			inter = append(inter, num)
			outer = append(outer, inter)
		}
	}
	return outer
}

func Permutation(processed string, unprocessed string) []string {
	if len(unprocessed) == 0 {
		return []string{processed}
	}
	store := make([]string, 0)
	ch := unprocessed[0]
	for i := range len(processed) + 1 {
		first := processed[0:i]
		second := processed[i:]
		store = append(store, Permutation(first+string(ch)+second, unprocessed[1:])...)
	}
	return store
}

func Permute(nums []int) [][]int {
	sort.Ints(nums)
	return helper([]int{}, nums)
}

func helper(p []int, up []int) [][]int {
	if len(up) == 0 {
		res := make([]int, len(p))
		copy(res, p)
		return [][]int{res}
	}
	result := make([][]int, 0)

	number := up[0]

	for i := 0; i <= len(p); i++ {
		newP := append(slices.Clone(p[0:i]), number)
		newP = append(newP, p[i:]...)
		result = append(result, helper(newP, up[1:])...)
	}
	return result

}
func PermuteUnique(nums []int) [][]int {
	frequencies := make(map[int]int)
	for _, n := range nums {
		frequencies[n]++
	}

	res := make([][]int, 0)
	perms := make([]int, 0, len(nums))
	res = backtrack(res, 0, nums, frequencies, perms)
	return res
}

func backtrack(res [][]int, idx int, nums []int, frequencies map[int]int, perms []int) [][]int {
	if idx == len(nums) {
		temp := make([]int, len(perms))
		copy(temp, perms)
		res = append(res, temp)
		return res
	}

	for n, freq := range frequencies {
		if freq == 0 {
			continue
		}

		frequencies[n]--
		perms = append(perms, n)

		res = backtrack(res, idx+1, nums, frequencies, perms)

		frequencies[n]++
		perms = perms[:len(perms)-1]
	}

	return res
}
