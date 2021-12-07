package breathFirstSearch_test

import (
	"breathFirstSearch"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFloodFill(t *testing.T) {
	t.Parallel()

	type testCase struct {
		image, want      [][]int
		sr, sc, newColor int
	}

	testCases := []testCase{
		{
			image: [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:    1, sc: 1, newColor: 2,
			want: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			image: [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:    0, sc: 0, newColor: 2,
			want: [][]int{{2, 2, 2}, {2, 2, 2}},
		},
	}

	for _, tc := range testCases {
		got := breathFirstSearch.FloodFill(tc.image, tc.sr, tc.sc, tc.newColor)

		if !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	t.Parallel()

	type testCase struct {
		grid [][]int
		want int
	}

	testCases := []testCase{
		{grid: [][]int{
			{1, 1, 0},
			{1, 0, 0},
			{1, 0, 1}},
			want: 4},
		{grid: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
			want: 6},
	}

	for _, tc := range testCases {
		got := breathFirstSearch.MaxAreaOfIsland(tc.grid)

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
