package solution

func AverageOfLevels(root *TreeNode) []float64 {
	var sum []float64 = []float64{}
	var tot []float64 = []float64{}

	dfsForAverageOfLevels(root, &sum, &tot, 0)
	for i := 0; i < len(sum); i++ {
		sum[i] /= tot[i]
	}

	return sum
}

func dfsForAverageOfLevels(root *TreeNode, sum, tot *[]float64, level int) {
	if root == nil {
		return

	}
	if len(*sum) <= level {
		*sum = append(*sum, 0)
		*tot = append(*tot, 0)
	}
	(*sum)[level] += float64(root.Val)
	(*tot)[level]++
	dfsForAverageOfLevels(root.Left, sum, tot, level+1)
	dfsForAverageOfLevels(root.Right, sum, tot, level+1)
}
