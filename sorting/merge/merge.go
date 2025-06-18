package merge

import "slices"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2

	leftCopy := slices.Clone(arr[:mid])
	left := MergeSort(leftCopy)
	rightCopy := slices.Clone(arr[mid:])
	right := MergeSort(rightCopy)
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// for i < len(left) {
	// 	result = append(result, left[i])
	// 	i++
	// }
	// for j < len(right) {
	// 	result = append(result, right[j])
	// 	j++
	// }
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func MergeSortInPlace(arr []int, start, end int) {
	if end-start <= 1 {
		return
	}
	mid := start + (end-start)/2
	MergeSortInPlace(arr, start, mid)
	MergeSortInPlace(arr, mid, end)
	MergeInPlace(arr, start, mid, end)
}

func MergeInPlace(arr []int, start, mid, end int) {
	mix := make([]int, 0, end-start)
	i, j := start, mid
	for i < mid && j < end {
		if arr[j] < arr[i] {
			mix = append(mix, arr[j])
			j++
		} else {
			mix = append(mix, arr[i])
			i++
		}
	}
	mix = append(mix, arr[i:mid]...)
	mix = append(mix, arr[j:end]...)

	copy(arr[start:end], mix)

}
