package library

func BuildNonDirGraph(edges [][]int, n int) [][]int {
	graph := make([][]int, n)

	for i := range edges {
		u := edges[i][0]
		v := edges[i][1]
		if len(graph[u]) == 0 {
			graph[u] = make([]int, 0)
		}
		graph[u] = append(graph[u], v)
		if len(graph[v]) == 0 {
			graph[v] = make([]int, 0)
		}
		graph[v] = append(graph[v], u)
	}
	return graph

}
func BuildDirGraph(edges [][]int, n int) [][]int {
	graph := make([][]int, n)

	for i := range edges {
		u := edges[i][0]
		v := edges[i][1]
		if len(graph[u]) == 0 {
			graph[u] = make([]int, 0)
		}
		graph[u] = append(graph[u], v)
	}
	return graph

}
