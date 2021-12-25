package tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func SliceToTree(slice []interface{}) *TreeNode {
	root := new(*TreeNode)
	insertSliceToTree(root, slice, 0)
	return *root
}

func IsSameTree(a *TreeNode, b *TreeNode) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return IsSameTree(a.Left, b.Left) && IsSameTree(a.Right, b.Right)
}

func insertSliceToTree(n **TreeNode, slice []interface{}, index int) {
	if index >= len(slice) || slice[index] == nil {
		return
	}
	*n = new(TreeNode)
	(*n).Val = slice[index].(int)
	insertSliceToTree(&((*n).Left), slice, 2*index+1)
	insertSliceToTree(&((*n).Right), slice, 2*index+2)
}

// You are given two binary trees root1 and root2.
// Imagine that when you put one of them to cover the other,
// some nodes of the two trees are overlapped while the others are not.
// You need to merge the two trees into a new binary tree.
// The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node.
// Otherwise, the NOT null node will be used as the node of the new tree.
// Return the merged tree.
// Note: The merging process must start from the root nodes of both trees.
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)

	return root1
}

// You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
// Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.
