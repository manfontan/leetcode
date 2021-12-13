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

	l := linkedList.List{}

	testCases := []testCase{
		{head: l.Init([]int{1, 2, 3, 4, 5, 6}), want: []int{1, 2, 3, 4, 6}, n: 2},
		{head: l.Init([]int{1}), want: []int{}, n: 1},
		{head: l.Init([]int{1, 2}), want: []int{1}, n: 1},
		{head: l.Init([]int{1, 2}), want: []int{2}, n: 2},
	}

	for _, tc := range testCases {

		got := linkedList.RemoveNthFromEnd(tc.head, tc.n)
		lgot := linkedList.List{Head: got}

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

	l := linkedList.List{}

	testCases := []testCase{
		{head: l.Init([]int{1, 2, 3, 4}), want: 3},
		{head: l.Init([]int{1, 2, 4}), want: 2},
	}

	for _, tc := range testCases {
		got := linkedList.MiddleNode(tc.head).Val

		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}

func TestReverseList(t *testing.T) {

	type testCase struct {
		head *linkedList.ListNode
		want []int
	}

	l := linkedList.List{}

	testCases := []testCase{
		{head: nil, want: []int{}},
		{head: l.Init([]int{1, 2, 3, 4, 5}), want: []int{5, 4, 3, 2, 1}},
		{head: l.Init([]int{1, 2, 3}), want: []int{3, 2, 1}},
	}

	for _, tc := range testCases {
		got := linkedList.ReverseList(tc.head)
		lgot := linkedList.List{Head: got}

		if !cmp.Equal(lgot.ToSlice(), tc.want) {
			t.Errorf("got %v, want %v", lgot.ToSlice(), tc.want)
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	t.Parallel()

	type testCase struct {
		list1, list2 *linkedList.ListNode
		want         []int
	}

	l := linkedList.List{}

	testCases := []testCase{
		{list1: l.Init([]int{1, 3, 4}), list2: l.Init([]int{2, 5, 6}), want: []int{1, 2, 3, 4, 5, 6}},
		{list1: l.Init([]int{1, 2, 4}), list2: l.Init([]int{1, 3, 4}), want: []int{1, 1, 2, 3, 4, 4}},
	}

	for _, tc := range testCases {

		got := linkedList.MergeTwoLists(tc.list1, tc.list2)
		lgot := linkedList.List{Head: got}

		if !cmp.Equal(lgot.ToSlice(), tc.want) {
			t.Errorf("got %v, want %v", lgot.ToSlice(), tc.want)
		}

	}
}
