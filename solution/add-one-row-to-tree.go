package solution

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		var head *TreeNode = &TreeNode{val, nil, nil}
		head.Left = root
		return head
	}
	addOneRowHelper(root, val, depth-1)
	return root
}
func addOneRowHelper(root *TreeNode, val, dep int) {
	if root == nil {
		return
	}
	if dep == 1 {
		var newNodeL *TreeNode = &TreeNode{val, nil, nil}
		var newNodeR *TreeNode = &TreeNode{val, nil, nil}
		leftSubtree := root.Left
		rightSubtree := root.Right
		root.Left = newNodeL
		root.Right = newNodeR
		newNodeL.Left = leftSubtree
		newNodeR.Right = rightSubtree
		return
	}
	addOneRowHelper(root.Left, val, dep-1)
	addOneRowHelper(root.Right, val, dep-1)

}
