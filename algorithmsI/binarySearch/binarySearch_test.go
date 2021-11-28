package binarySearch_test

import (
	"binarySearch"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	nums := []int{-1, 0, 3, 5, 9, 12}

	type testCase struct {
		nums         []int
		target, want int
	}

	testCases := []testCase{
		{nums: nums, target: -1, want: 0},
		{nums: nums, target: 0, want: 1},
		{nums: nums, target: 3, want: 2},
		{nums: nums, target: 5, want: 3},
		{nums: nums, target: 9, want: 4},
		{nums: nums, target: 12, want: 5},
	}

	for _, tc := range testCases {
		got := binarySearch.Search(tc.nums, tc.target)

		if tc.want != got {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}

func TestFirstBadVersion(t *testing.T) {
	t.Parallel()

	type testCase struct {
		n, want int
	}

	testCases := []testCase{
		{want: 1, n: 5},
		{want: 2, n: 5},
		{want: 3, n: 5},
		{want: 4, n: 5},
		{want: 5, n: 5},
	}

	for _, tc := range testCases {
		binarySearch.BadVersion = tc.want

		got := binarySearch.FirstBadVersion(tc.n)

		if tc.want != got {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}

func TestSearchInsert(t *testing.T) {
	t.Parallel()

	type testCase struct {
		nums         []int
		target, want int
	}
	nums := []int{1, 2, 4, 6}
	testCases := []testCase{
		{nums: nums, want: 3, target: 5},
		{nums: nums, want: 0, target: 0},
		{nums: nums, want: 4, target: 7},
		{nums: nums, want: 2, target: 3},
	}

	for _, tc := range testCases {
		got := binarySearch.SearchInsert(tc.nums, tc.target)

		if tc.want != got {
			t.Errorf("want %d, got %d", tc.want, got)
		}
	}
}
