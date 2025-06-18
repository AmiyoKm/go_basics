package main

import (
	"fmt"
	"math"
	"sort"
)

func kthFactor(n int, k int) int {
	var factors []int
	for i :=1 ; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	sort.Ints(factors)

	if k > len(factors) {
		return -1
	}

	return factors[k-1]
}

func sortedSquares(nums []int) []int {
    for i:=range nums {
		nums[i] = nums[i]*nums[i]
	}
	sort.Ints(nums)
	return nums
}
func rotate(matrix [][]int)  {
	//rotate 90 degree clockwise
	n := len(matrix)
	for i :=0 ; i <n ; i++ {
		for j:=i; j <n; j++ {
			matrix[i][j] , matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i:=0; i<n; i++ {
		for j :=0 ; j <n/2 ; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
func isValidSudoku(board [][]byte) bool {
    for i:=0 ; i <9 ; i++ {
		row := make(map[byte]bool)
		col := make(map[byte]bool)
		box := make(map[byte]bool)
		for j:=0 ; j <9 ; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}
			boxRow := 3*(i/3) + j/3
			boxCol := 3*(i%3) + j%3
			if board[boxRow][boxCol] != '.' {
				if box[board[boxRow][boxCol]] {
					return false
				}
				box[board[boxRow][boxCol]] = true
			}
		}
	}
	return true
}
func maxSubArray(nums []int) int {
    if len(nums)==0 {
		return 0
	}
	maxSum :=math.MinInt32
	currentSum := 0

	for i:=range nums {
		currentSum = max(nums[i] , currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
func lengthOfLongestSubstring(s string) int {
    charMap := make(map[byte]int)
	maxLength := 0
	left :=0

	for right:=range len(s) {
		index , found := charMap[s[right]]
		if found && index >=left {
			left = index+1
		}
		charMap[s[right]] = right
		maxLength = max(maxLength,right-left+1)
	}
	return maxLength
}

func main() {

	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(kthFactor(n, k))

}
