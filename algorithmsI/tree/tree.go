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
