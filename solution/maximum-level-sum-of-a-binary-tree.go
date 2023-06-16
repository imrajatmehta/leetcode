package solution

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxLevelSum(root *TreeNode) int {
	sumMap := make(map[int]int)
	dfsMaxLevelSum(root, sumMap, 0)
	maxInd := 0
	maxSum := math.MaxInt
	for key := range sumMap {
		if maxSum < sumMap[key] {
			maxInd = key
			maxSum = sumMap[key]
		}
	}
	return maxInd
}
func dfsMaxLevelSum(root *TreeNode, sumMap map[int]int, level int) {
	if root == nil {
		return
	}
	if _, ok := sumMap[level]; !ok {
		sumMap[level] = 0
	}
	sumMap[level] += root.Val
	dfsMaxLevelSum(root.Left, sumMap, level+1)
	dfsMaxLevelSum(root.Right, sumMap, level+1)
}
