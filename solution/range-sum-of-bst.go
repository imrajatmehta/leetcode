package solution

func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	ans := 0

	if low <= root.Val && root.Val <= high {
		ans = root.Val
	}
	if low < root.Val {
		ans += RangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		ans += RangeSumBST(root.Right, low, high)
	}

	return ans
}
