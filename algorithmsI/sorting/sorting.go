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

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Time complexity: O(n). The input array is traversed at most once. Thus the time complexity is O(n).
// Space complexity: O(1). We only use additional space to store two indices and the sum, so the space complexity is O(1).
func TwoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1

	for l < r {
		s := nums[l] + nums[r]
		if s < target {
			l++
		} else if s > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}

	return []int{-1, -1}
}
