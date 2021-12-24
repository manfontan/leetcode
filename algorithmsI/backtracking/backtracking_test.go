package backtracking_test

import (
	"backtracking"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPermute(t *testing.T) {
	t.Parallel()

	type testCase struct {
		nums []int
		want [][]int
	}

	testCases := []testCase{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
	}

	for _, tc := range testCases {
		got := backtracking.Permute(tc.nums)

		if !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
