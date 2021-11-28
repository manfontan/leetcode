package sorting_test

import (
	"github.com/google/go-cmp/cmp"
	"sorting"
	"testing"
)

func TestSortedSquares(t *testing.T) {

	type testCase struct {
		nums []int
		want []int
	}

	testCases := []testCase{
		{nums: []int{-4, -1, 2, 3}, want: []int{1, 4, 9, 16}},
		{nums: []int{-2, -1, 5, 7}, want: []int{1, 4, 25, 49}},
		{nums: []int{-1, -2, 3, 4}, want: []int{1, 4, 9, 16}},
		{nums: []int{-7, -5, 1, 2}, want: []int{1, 4, 25, 49}},
	}

	for _, tc := range testCases {

		got := sorting.SortedSquares(tc.nums)

		if !cmp.Equal(tc.want, got) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
