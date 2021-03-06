package bitwise_test

import (
	"bitwise"
	"strconv"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	t.Parallel()

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

func TestReverseBits(t *testing.T) {
	t.Parallel()

	type testCase struct {
		n, want uint32
	}

	testCases := []testCase{
		{n: parseUint32("00110001110011001100111000110011"), want: parseUint32("11001100011100110011001110001100")},
		{n: parseUint32("10111111111111111111111111111111"), want: parseUint32("11111111111111111111111111111101")},
	}

	for _, tc := range testCases {
		got := bitwise.ReverseBits(tc.n)

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

func parseUint32(s string) uint32 {
	u, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint32(u)
}

func TestHammingWeight(t *testing.T) {
	t.Parallel()

	type testCase struct {
		n    uint32
		want int
	}

	testCases := []testCase{
		{n: 000000000000000000000000000001011, want: 3},
		{n: 000000000000000000000000000000000, want: 0},
	}

	for _, tc := range testCases {
		got := bitwise.HammingWeight(tc.n)

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

func TestSingleNumber(t *testing.T) {
	t.Parallel()

	type testCase struct {
		nums []int
		want int
	}

	testCases := []testCase{
		{nums: []int{2, 2, 1}, want: 1},
		{nums: []int{2, 2, 3, 4, 4}, want: 3},
		{nums: []int{2, 2, 3, 4, 4, 7, 3}, want: 7},
	}

	for _, tc := range testCases {
		got := bitwise.SingleNumber(tc.nums)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
