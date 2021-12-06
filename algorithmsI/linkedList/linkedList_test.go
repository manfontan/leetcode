package linkedList_test

import (
	"linkedList"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		head *linkedList.ListNode
		want []int
		n    int
	}

	l := linkedList.List{Head: nil, Len: 0}
	l.Init([]int{1, 2, 3, 4, 5, 6})
	l2 := linkedList.List{Head: nil, Len: 0}
	l2.Init([]int{1})
	l3 := linkedList.List{Head: nil, Len: 0}
	l3.Init([]int{1, 2})
	l4 := linkedList.List{Head: nil, Len: 0}
	l4.Init([]int{1, 2})

	testCases := []testCase{
		{head: l.Head, want: []int{1, 2, 3, 4, 6}, n: 2},
		{head: l2.Head, want: []int{}, n: 1},
		{head: l3.Head, want: []int{1}, n: 1},
		{head: l4.Head, want: []int{2}, n: 2},
	}

	for _, tc := range testCases {

		got := linkedList.RemoveNthFromEnd(tc.head, tc.n)
		lgot := linkedList.List{Head: got, Len: 0}

		if !cmp.Equal(lgot.ToSlice(), tc.want) {
			t.Errorf("got %v, want %v", lgot.ToSlice(), tc.want)
		}
	}
}

func TestMiddleNode(t *testing.T) {
	t.Parallel()

	type testCase struct {
		head *linkedList.ListNode
		want int
	}

	l := linkedList.List{Head: nil, Len: 0}
	l.Init([]int{1, 2, 3, 4})
	l2 := linkedList.List{Head: nil, Len: 0}
	l2.Init([]int{1, 2, 4})

	testCases := []testCase{
		{head: l.Head, want: 3},
		{head: l2.Head, want: 2},
	}

	for _, tc := range testCases {
		got := linkedList.MiddleNode(tc.head).Val

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
