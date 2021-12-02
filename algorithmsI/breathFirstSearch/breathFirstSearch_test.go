package breathFirstSearch_test

import (
	"breathFirstSearch"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFloodFill(t *testing.T) {
	t.Parallel()

	type testCase struct {
		image, want      [][]int
		sr, sc, newColor int
	}

	testCases := []testCase{
		{
			image: [][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}},
			sr:    1, sc: 1, newColor: 2,
			want: [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}},
		},
		{
			image: [][]int{[]int{0, 0, 0}, []int{0, 0, 0}},
			sr:    0, sc: 0, newColor: 2,
			want: [][]int{[]int{2, 2, 2}, []int{2, 2, 2}},
		},
	}

	for _, tc := range testCases {
		got := breathFirstSearch.FloodFill(tc.image, tc.sr, tc.sc, tc.newColor)

		if !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}