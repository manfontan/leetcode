package sorting

import "math"

// Description: Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
// Time Complexity: O(N) where N is the lenth of the slice
// Space Complexity: O(N) if we count the output O(1) otherwise.
func SortedSquares(nums []int) []int {

	l := 0
	r := len(nums) - 1

	sortedSquares := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {

		if math.Abs(float64(nums[r])) > math.Abs(float64(nums[l])) {
			sortedSquares[i] = nums[r] * nums[r]
			r--
		} else {
			sortedSquares[i] = nums[l] * nums[l]
			l++
		}
	}

	return sortedSquares
}
