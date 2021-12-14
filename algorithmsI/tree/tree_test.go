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
