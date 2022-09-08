package solution

func InorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	// dfsInorderTraversal(root, &ans)
	return predecessorInorderTraversal(root, &ans)
}
func dfsInorderTraversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	dfsInorderTraversal(root.Left, ans)
	dfsInorderTraversal(root.Right, ans)
}
func predecessorInorderTraversal(root *TreeNode, ans *[]int) []int {

	curr := root
	pre := root

	for curr != nil {

		if curr.Left == nil {
			*ans = append(*ans, curr.Val)
			curr = curr.Right
		} else {
			pre = curr.Left
			for pre.Right != nil && pre.Right != curr {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = curr
				curr = curr.Left
			} else {
				pre.Right = nil
				*ans = append(*ans, curr.Val)
				curr = curr.Right
			}
		}

	}
	return *ans
}
