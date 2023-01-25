package solution

import (
	"fmt"
	"math"

	"github.com/imrajatmehta/leetcode/library"
)

func ClosestMeetingNode(edges []int, node1 int, node2 int) int {
	distNode1 := FindDistDFS(edges, node1)
	distNode2 := FindDistDFS(edges, node1)

	ans := math.MaxInt
	for i := range distNode1 {
		if distNode1[i] == -1 || distNode2[i] == -1 {
			continue
		}

		if ans > distNode1[i]+distNode2[i] {
			ans = distNode1[i] + distNode2[i]
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	fmt.Println(distNode1, distNode2)
	return ans

}
func FindDistDFS(edges []int, node1 int) []int {
	n := len(edges)
	distNode := make([]int, n)
	distNode = library.FillInt(distNode, -1)
	vis := make([]bool, n)
	dist := 0
	distNode[node1] = 0
	for edges[node1] != -1 && !vis[node1] {
		dist++
		distNode[edges[node1]] = dist
		vis[node1] = true
		node1 = edges[node1]
	}
	return distNode
}
