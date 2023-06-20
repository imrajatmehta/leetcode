package solution

import (
	"fmt"

	"github.com/imrajatmehta/leetcode/library"
)

func CountPaths(grid [][]int) int {
	ans := 0
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = library.FillInt(make([]int, len(grid[i])), -1)
	}
	for i := range grid {
		for j := range grid[i] {
			ans = (ans + dfsCountPaths(grid, dp, i, j, len(grid), len(grid[i]), -1)) % mod
			fmt.Println("HEY")
		}
	}
	return ans % mod

}
func dfsCountPaths(grid, dp [][]int, i, j, n, m, prev int) int {

	if !library.IsValidIndexInMatrix(i, j, n, m) || prev >= grid[i][j] {
		return 0
	}
	if dp[i][j] > -1 {
		return dp[i][j]
	}
	fmt.Println(i, j, prev)
	ans := 1

	dirs := library.GetFourXYDirections()
	for k := range dirs {
		ans = (ans + dfsCountPaths(grid, dp, i+dirs[k][0], j+dirs[k][1], n, m, grid[i][j])) % mod
	}
	dp[i][j] = ans % mod
	return dp[i][j]
}
