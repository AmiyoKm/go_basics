package main

func TwoSum(arr []int, target int) [2]int {
	diffToIndex := make(map[int]int)

	for i, a := range arr {
		diff := target - a

		if v, ok := diffToIndex[a]; ok {
			return [2]int{v, i}
		}
		diffToIndex[diff] = i
	}
	return [2]int{-1, -1}
}
