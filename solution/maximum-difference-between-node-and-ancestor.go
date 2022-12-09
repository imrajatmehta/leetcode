package solution

import (
	"github.com/imrajatmehta/leetcode/library"
)

func MaxAncestorDiff(root *TreeNode) int {
	ans := 0
	MaxAncestorDiffA(root.Left, root.Val, root.Val, &ans)
	MaxAncestorDiffA(root.Left, root.Val, root.Val, &ans)
	return ans
}

func MaxAncestorDiffA(root *TreeNode, min, max int, ans *int) {
	if root == nil {
		return
	}
	(*ans) = library.MaxInt(library.AbsInt(root.Val-max), *ans)
	(*ans) = library.MaxInt(library.AbsInt(root.Val-min), *ans)
	MaxAncestorDiffA(root.Left, library.MinInt(root.Val, min), library.MaxInt(root.Val, max), ans)
	MaxAncestorDiffA(root.Right, library.MinInt(root.Val, min), library.MaxInt(root.Val, max), ans)
}
