package solution

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil && targetSum-root.Val == 0 {
		return true
	}
	if HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}
