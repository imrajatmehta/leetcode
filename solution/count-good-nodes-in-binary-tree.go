package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GoodNodes(root *TreeNode) int {

	return goodNodesUtils(root, root.Val)

}
func goodNodesUtils(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	if root.Val >= max {
		return 1 + goodNodesUtils(root.Left, root.Val) + goodNodesUtils(root.Right, root.Val)
	} else {
		return goodNodesUtils(root.Left, max) + goodNodesUtils(root.Right, max)
	}
}
