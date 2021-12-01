package mystrings

// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Time complexity : O(N)\mathcal{O}(N)O(N) to swap N/2N/2N/2 element.
// Space complexity : O(1)\mathcal{O}(1)O(1), it's a constant space solution.

func ReverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
