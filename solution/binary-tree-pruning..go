package solution

func PruneTree(root *TreeNode) *TreeNode {
	if removeZeroSubtrees(root) {

		return nil
	}
	return root
}
func removeZeroSubtrees(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := removeZeroSubtrees(root.Left)
	right := removeZeroSubtrees(root.Right)
	if left && right && root.Val == 0 {
		return true
	}
	if left {
		root.Left = nil

	}
	if right {
		root.Right = nil
	}
	return false
}
