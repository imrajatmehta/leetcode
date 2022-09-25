package solution

func PathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	pathSumDfs(root, targetSum, &ans, path)

	return ans
}
func pathSumDfs(root *TreeNode, targetSum int, ans *[][]int, path []int) {

	if root == nil {
		return

	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && targetSum == root.Val {

		temp := make([]int, len(path))
		copy(temp, path)

		*ans = append(*ans, temp)
	}

	pathSumDfs(root.Left, targetSum-root.Val, ans, path)
	pathSumDfs(root.Right, targetSum-root.Val, ans, path)
	path = path[:len(path)-1]
}
