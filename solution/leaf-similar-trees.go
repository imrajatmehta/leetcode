package solution

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leftLeafs := make([]int, 0)
	rightLeafs := make([]int, 0)
	FindAllLeafsInorder(root2, &leftLeafs)
	FindAllLeafsInorder(root1, &rightLeafs)
	if len(leftLeafs) != len(rightLeafs) {
		return false
	}
	for i := range leftLeafs {
		if leftLeafs[i] != rightLeafs[i] {
			return false
		}
	}
	return true
}
func FindAllLeafsInorder(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}
	FindAllLeafsInorder(root.Left, leafs)
	if root.Left == nil && root.Right == nil {
		(*leafs) = append(*leafs, root.Val)
	}
	FindAllLeafsInorder(root.Right, leafs)
}
