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

//Given an array, rotate the array to the right by k steps, where k is non-negative.
func Rotate(nums []int, k int) []int {

	t := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		t[(i+k)%len(nums)] = nums[i]
	}

	return t
}

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
func MoveZeroes(nums []int) []int {

	l := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && nums[i+1] != 0 {
			nums[l] = nums[i+1]
			nums[i+1] = 0
			l++
		} else if nums[i] != 0 {
			l++
		}
	}
	return nums
}
