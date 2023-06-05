package solution

import "github.com/imrajatmehta/leetcode/library"

func MaximumDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := make([][]int, n)
	max := 1
	for i := range bombs {
		for j := range bombs {
			if i == j {
				continue
			}
			dist := library.GetDistanceBtwTwoPointsOnPlane(float64(bombs[i][0]), float64(bombs[i][1]), float64(bombs[j][0]), float64(bombs[j][1]))
			if dist <= float64(bombs[i][2]) {
				graph[i] = append(graph[i], j)
			}
		}
	}
	for i := range graph {
		count := dfsMaximumDetonation(graph, i, make([]bool, n))
		if count < max {
			max = count
		}
	}

	return max
}
func dfsMaximumDetonation(graph [][]int, i int, vis []bool) int {
	vis[i] = true
	ans := 0
	for _, child := range graph[i] {
		if !vis[child] {
			ans += dfsMaximumDetonation(graph, child, vis) + 1
		}

	}
	return ans
}
