package solution

import (
	"github.com/imrajatmehta/leetcode/library"
)

func ClosedIsland(grid [][]int) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	vis := library.CreateAndFillMatrix(n, m, false)
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 0 && !vis[i][j].(bool) {
				if closedIslandDFS(i, j, n, m, grid, vis) {
					ans++
				}
			}
		}
	}
	return ans
}
func closedIslandDFS(i, j, n, m int, grid [][]int, vis [][]interface{}) bool {
	if vis[i][j].(bool) {
		return true
	}

	vis[i][j] = true
	a := true

	if !isFine(i+1, j, n, m, grid, vis) {
		a = false
	}
	if !isFine(i, j+1, n, m, grid, vis) {
		a = false
	}
	if !isFine(i-1, j, n, m, grid, vis) {
		a = false
	}
	if !isFine(i, j-1, n, m, grid, vis) {
		a = false
	}
	return a
}
func isCornerIndex(i, j, n, m int) bool {
	return i == 0 || j == 0 || i == n-1 || j == m-1
}
func isFine(i, j, n, m int, grid [][]int, vis [][]interface{}) bool {
	ans := true
	if library.IsValidIndexInMatrix(i, j, n, m) && grid[i][j] == 0 && !vis[i][j].(bool) {
		if isCornerIndex(i, j, n, m) {
			ans = false
			closedIslandDFS(i, j, n, m, grid, vis)
		} else {
			ans = closedIslandDFS(i, j, n, m, grid, vis)
		}
	}
	return ans

}
