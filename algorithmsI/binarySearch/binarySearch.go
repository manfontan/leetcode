package binarySearch

func Search(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if target > nums[median] {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(nums) || nums[low] != target {
		return -1
	}

	return low
}

func FirstBadVersion(n int) int {
	return 0
}
