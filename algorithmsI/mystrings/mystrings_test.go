package mystrings_test

import (
	"github.com/google/go-cmp/cmp"
	"mystrings"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input, want []byte
	}

	testCases := []testCase{
		{input: []byte("hello"), want: []byte("olleh")},
	}

	for _, tc := range testCases {
		got := mystrings.ReverseString(tc.input)

		if !cmp.Equal(tc.want, got) {
			t.Errorf("got %s, want %s", string(got), string(tc.want))
		}
	}
}
