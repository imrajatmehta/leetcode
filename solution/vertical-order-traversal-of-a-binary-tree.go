package solution

import (
	"sort"

	"github.com/imrajatmehta/leetcode/library"
)

func VerticalTraversal(root *TreeNode) [][]int {
	var ans [][]int = make([][]int, 0)
	var cache map[int][][2]int = make(map[int][][2]int)
	var min, max int
	dfsForVerticalTraversal(root, 0, 0, &ans, cache, &min, &max)
	for i := min; i <= max; i++ {
		sort.Slice(cache[i], func(a, b int) bool {
			return cache[i][a][1] == cache[i][b][1] && cache[i][a][0] < cache[i][b][0] || cache[i][a][1] < cache[i][b][1]
		})
		var temp []int = []int{}
		for _, val := range cache[i] {
			temp = append(temp, val[0])
		}
		ans = append(ans, temp)
	}
	return ans
}
func dfsForVerticalTraversal(root *TreeNode, row, column int, ans *[][]int, cache map[int][][2]int, min, max *int) {
	if root == nil {
		return
	}

	if _, ok := cache[column]; !ok {
		cache[column] = [][2]int{}

	}
	*min = library.MinInt(*min, column)
	*max = library.MaxInt(*max, column)
	cache[column] = append(cache[column], [2]int{root.Val, row})
	dfsForVerticalTraversal(root.Left, row+1, column-1, ans, cache, min, max)
	dfsForVerticalTraversal(root.Right, row+1, column+1, ans, cache, min, max)
}
