package dynamicProgramming

import (
	"strings"
	"unicode"
)

// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.
// Time complexity: O(r⋅c)
// We perform two passes over the matrix and each pass requires O(r⋅c) time.
// Space complexity: O(1)
func UpdateMatrix(m [][]int) [][]int {
	r := len(m)
	c := len(m[0])
	dist := make([][]int, r)

	if r == 0 {
		return m
	}

	for i := 0; i < r; i++ {
		dist[i] = make([]int, c)
		for j := 0; j < c; j++ {
			dist[i][j] = r * c
			if m[i][j] == 0 {
				dist[i][j] = 0
			}
			if i > 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j > 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}

	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if m[i][j] == 0 {
				dist[i][j] = 0
			}
			if i > r-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j > c-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}

	return dist
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
// Return a list of all possible strings we could create. Return the output in any order.
func LetterCasePermutation(s string) []string {
	result := []string{""}

	for _, c := range s {
		n := len(result)
		if unicode.IsLetter(c) {
			for i := 0; i < n; i++ {
				result = append(result, result[i][:])
				result[i] += strings.ToLower(string(c))
				result[n+i] += strings.ToUpper(string(c))
			}
		} else {
			for i := 0; i < n; i++ {
				result[i] += string(c)
			}
		}
	}
	return result
}
