package solution

import (
	"math"

	"github.com/imrajatmehta/leetcode/library"
)

func MaxPathSumHelper(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt
	left := MaxPathSumHelper(root.Left, result)
	right := MaxPathSumHelper(root.Right, result)
	ans = library.MaxInt(ans, root.Val)
	ans = library.MaxInt(ans, root.Val+left)
	(*result) = library.MaxInt(*result, root.Val+right)
	(*result) = library.MaxInt(*result, ans)
	return ans

}
func MaxPathSum(root *TreeNode) int {
	ans := 0
	MaxPathSumHelper(root.Left, &ans)
	return ans

}
