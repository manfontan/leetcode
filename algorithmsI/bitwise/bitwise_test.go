package bitwise_test

import (
	"bitwise"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {

	type testCase struct {
		n    int
		want bool
	}

	testCases := []testCase{
		{n: 0, want: false},
		{n: 1, want: true},
		{n: 2, want: true},
		{n: 3, want: false},
		{n: 4, want: true},
		{n: 5, want: false},
		{n: 6, want: false},
		{n: 7, want: false},
		{n: 8, want: true},
	}

	for _, tc := range testCases {
		got := bitwise.IsPowerOfTwo(tc.n)

		if got != tc.want {
			t.Errorf("for %d : got %v, want %v", tc.n, got, tc.want)
		}
	}
}
