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
		{input: []byte("Hannah"), want: []byte("hannaH")},
	}

	for _, tc := range testCases {
		got := mystrings.ReverseString(tc.input)

		if !cmp.Equal(tc.want, got) {
			t.Errorf("got %s, want %s", string(got), string(tc.want))
		}
	}
}

func TestReverseWords(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input, want string
	}

	testCases := []testCase{
		{input: "hello my friend", want: "olleh ym dneirf"},
		{input: "Let's take LeetCode contest", want: "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, tc := range testCases {
		got := mystrings.ReverseWords(tc.input)

		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
