package solution

import "github.com/imrajatmehta/leetcode/library"

func MaxProduct(root *TreeNode) int {
	totalSum := NodeSumTree(root)
	ans := 0
	maxProductHelper(root, &totalSum, &ans)
	return ans
}
func maxProductHelper(root *TreeNode, totalSum, ans *int) int {
	if root == nil {
		return 0
	}

	leftS := maxProductHelper(root.Left, totalSum, ans)
	rightS := maxProductHelper(root.Right, totalSum, ans)

	(*ans) = library.MaxInt(*ans, library.MaxInt((*totalSum-leftS)*leftS, (*totalSum-rightS)*rightS))

	return leftS + root.Val + rightS
}
func NodeSumTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return NodeSumTree(root.Left) + NodeSumTree(root.Right) + root.Val
}
