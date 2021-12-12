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

// Reverse bits of a given 32 bits unsigned integer.
func ReverseBits(n uint32) uint32 {

	n = n>>16 | n<<16
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)

	return n
}
