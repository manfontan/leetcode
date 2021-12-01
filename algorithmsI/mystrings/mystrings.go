package mystrings

import "strings"

// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
// Time complexity : O(N) to swap N/2 element.
// Space complexity : O(1), it's a constant space solution.
func ReverseString(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace
// and initial word order.
// Time complexity : O(n). where nnn is the length of the string.
// Space complexity : O(n). resresres of size nnn is used.

func ReverseWords(s string) string {
	r := []string{}
	for _, w := range strings.Split(s, " ") {
		r = append(r, reverse(w))
	}
	return strings.Join(r, " ")
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
