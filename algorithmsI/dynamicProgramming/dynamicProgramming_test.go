package dynamicProgramming_test

import (
	"dynamicProgramming"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUpdateMatrix(t *testing.T) {
	t.Parallel()

	type testCase struct {
		m, want [][]int
	}

	testCases := []testCase{
		{m: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
	}

	for _, tc := range testCases {
		got := dynamicProgramming.UpdateMatrix(tc.m)

		if !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestLetterCasePermutation(t *testing.T) {
	t.Parallel()

	type testCase struct {
		s    string
		want []string
	}

	testCases := []testCase{
		{s: "a1b2", want: []string{"a1b2", "A1b2", "a1B2", "A1B2"}},
	}

	for _, tc := range testCases {
		got := dynamicProgramming.LetterCasePermutation(tc.s)

		if !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
