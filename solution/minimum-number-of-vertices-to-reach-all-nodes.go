package solution

import "github.com/imrajatmehta/leetcode/library"

func FindSmallestSetOfVertices(n int, edges [][]int) []int {
	ans := []int{}
	vis := make([]bool, n)
	ansMap := make(map[int]bool)
	graph := library.BuildDirGraph(edges, n)
	for i := 0; i < n; i++ {
		if !vis[i] {
			dfsFindSmallestSetOfVertices(graph, i, vis, ansMap)
			ansMap[i] = true
		}

	}
	for key := range ansMap {
		ans = append(ans, key)
	}
	return ans
}
func dfsFindSmallestSetOfVertices(graph [][]int, i int, vis []bool, ansMap map[int]bool) {
	vis[i] = true
	for _, child := range graph[i] {
		if !vis[child] {
			dfsFindSmallestSetOfVertices(graph, child, vis, ansMap)
		} else if ok := ansMap[child]; ok {
			delete(ansMap, child)
		}

	}
}
