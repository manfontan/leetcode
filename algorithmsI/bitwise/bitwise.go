package bitwise

// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2^x.
// Time complexity : O(1).
// Space complexity : O(1).
func IsPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(-n) == n {
		return true
	}

	return false
}
