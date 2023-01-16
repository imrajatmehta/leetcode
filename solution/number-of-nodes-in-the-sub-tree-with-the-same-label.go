package solution

func CountSubTrees(n int, edges [][]int, labels string) []int {
	graph := make([][]int, n)
	ans := make([]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	vis := make([]bool, n)
	dfsGraph(0, graph, &ans, labels, vis)
	return ans
}
func dfsGraph(i int, graph [][]int, ans *[]int, labels string, vis []bool) []int {
	count := make([]int, 26)
	if vis[i] {
		return count
	}
	vis[i] = true
	for _, child := range graph[i] {
		tempCount := dfsGraph(child, graph, ans, labels, vis)
		for i := range tempCount {
			count[i] += tempCount[i]
		}
	}

	count[labels[i]-'a']++
	(*ans)[i] = count[labels[i]-'a']
	return count

}
