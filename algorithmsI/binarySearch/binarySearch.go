package binarySearch

func Search(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {
		median := (low + high) / 2

		if target == nums[median] {
			return median
		} else if target > nums[median] {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return -1
}

func FirstBadVersion(n int) int {
	l := 1
	r := n

	for l < r {
		m := (l + r) / 2

		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

var BadVersion = 1

func isBadVersion(version int) bool {
	return version >= BadVersion
}

func SearchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
