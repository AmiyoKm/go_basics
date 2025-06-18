package quick

func QuickSort(nums []int, low int, high int) {
	if low > high {
		return
	}
	start := low
	end := high
	mid := start + (end-start)/2
	pivot := nums[mid]

	for start <= end {
		for nums[start] < pivot {
			start++
		}
		for nums[end] > pivot {
			end--
		}
		if start <= end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
		// now pivot is at its correct position in the array
		// now we need to sort the left and right sub arrays
		QuickSort(nums, low, end)
		QuickSort(nums, start, high)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0, m+n)
	if n == 0 {
		return
	}
	i := 0
	j := 0

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for i < m {
		result = append(result, nums1[i])
		i++
	}
	for j < n {
		result = append(result, nums2[j])
		j++
	}
	copy(nums1, result)

}
