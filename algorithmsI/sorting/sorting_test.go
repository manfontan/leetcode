package sorting_test

import (
	"github.com/google/go-cmp/cmp"
	"sorting"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	t.Parallel()

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

func TestRotate(t *testing.T) {
	t.Parallel()

	type testCase struct {
		nums, want []int
		k          int
	}

	testCases := []testCase{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, want: []int{5, 6, 7, 1, 2, 3, 4}, k: 3},
		{nums: []int{1, 2}, want: []int{2, 1}, k: 3},
	}

	for _, tc := range testCases {
		got := sorting.Rotate(tc.nums, tc.k)
		if !cmp.Equal(tc.want, got) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}

}

func TestMoveZeroes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		nums, want []int
	}

	testCases := []testCase{
		{nums: []int{1, 0, 0, 2, 3, 0, 4}, want: []int{1, 2, 3, 4, 0, 0, 0}},
		{nums: []int{0, 1, 0, 2, 3, 0, 4}, want: []int{1, 2, 3, 4, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := sorting.MoveZeroes(tc.nums)

		if !cmp.Equal(tc.want, got) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
