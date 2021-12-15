package tree_test

import (
	"testing"
	"tree"
)

func TestSliceToTree(t *testing.T) {
	t.Parallel()

	type testCase struct {
		slice []interface{}
		want  tree.TreeNode
	}

	testCases := []testCase{
		{
			slice: []interface{}{1, 2, 3},
			want:  *tree.SliceToTree([]interface{}{1, 2, 3}),
		},
	}

	for _, tc := range testCases {
		got := tree.SliceToTree(tc.slice)

		if !tree.IsSameTree(got, &tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestMergeTrees(t *testing.T) {

	type testCase struct {
		root1, root2, want *tree.TreeNode
	}

	testCases := []testCase{
		{
			root1: tree.SliceToTree([]interface{}{1, 2, 3, 4, 5}),
			root2: tree.SliceToTree([]interface{}{2, 3, 4, 5, 6}),
			want:  tree.SliceToTree([]interface{}{3, 5, 7, 9, 11}),
		},
	}

	for _, tc := range testCases {
		got := tree.MergeTrees(tc.root1, tc.root2)

		if !tree.IsSameTree(got, tc.want) {
			t.Errorf("got %v, want %v", *got, tc.want)
		}
	}
}
