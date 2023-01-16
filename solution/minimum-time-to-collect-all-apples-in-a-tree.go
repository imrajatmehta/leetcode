package solution

func MinTime(n int, edges [][]int, hasApple []bool) int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	vis := make([]bool, n)
	ans, _ := dfsGraphMinTime(0, graph, vis, hasApple)
	return ans
}
func dfsGraphMinTime(i int, graph [][]int, vis, hasApple []bool) (int, bool) {
	if vis[i] {
		return 0, false
	}
	total := 0
	vis[i] = true
	for _, child := range graph[i] {
		steps, present := dfsGraphMinTime(child, graph, vis, hasApple)
		if present {
			total += steps + 2
		}

	}
	if total > 0 {
		return total, true
	} else if hasApple[i] {
		return 0, true
	}
	return 0, false
}
