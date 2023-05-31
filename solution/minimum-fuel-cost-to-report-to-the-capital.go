package solution

import (
	"github.com/imrajatmehta/leetcode/library"
)

func MinimumFuelCost(roads [][]int, seats int) int64 {
	graph := library.BuildNonDirGraph(roads, len(roads)+1)
	vis := make([]bool, len(graph))
	var ss int64 = int64(seats)
	ans, _ := minimumFuelCostHelper(graph, 0, len(graph), ss, vis)
	return ans
}
func minimumFuelCostHelper(graph [][]int, i, n int, seats int64, vis []bool) (int64, int64) {
	if i >= n {
		return 0, 0
	}
	var totN int64 = 0
	var totDist int64 = 0
	vis[i] = true
	for k := range graph[i] {
		child := graph[i][k]
		if !vis[child] {
			dist, n := minimumFuelCostHelper(graph, child, n, seats, vis)
			totDist += dist
			totN += n
		}
	}
	if totN+1 > seats {
		if (totN+1)%seats == 0 {
			return totDist + (totN+1)/seats, totN + 1
		} else {
			return 1 + totDist + (totN+1)/seats, totN + 1
		}

	} else {
		return totDist + 1, totN + 1
	}

}
