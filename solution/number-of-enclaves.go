package solution

import "github.com/imrajatmehta/leetcode/library"

func NumEnclaves(grid [][]int) int {

	ans := 0
	n, m := len(grid), len(grid[0])
	vis := library.CreateAndFillMatrix(n, m, false)
	for i := 0; i < n; i++ {
		if grid[i][0] == 1 && !vis[i][0].(bool) {
			NumEnclavesBFS(i, 0, n, m, grid, vis)
		}
		if grid[i][m-1] == 1 && !vis[i][m-1].(bool) {
			NumEnclavesBFS(i, m-1, n, m, grid, vis)
		}
	}
	for j := 0; j < m; j++ {
		if grid[0][j] == 1 && !vis[0][j].(bool) {
			NumEnclavesBFS(0, j, n, m, grid, vis)
		}
		if grid[n-1][j] == 1 && !vis[n-1][j].(bool) {
			NumEnclavesBFS(n-1, j, n, m, grid, vis)
		}
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if grid[i][j] == 1 && !vis[i][j].(bool) {
				ans++
			}
		}
	}
	return ans

}
func NumEnclavesBFS(i, j, n, m int, grid [][]int, vis [][]interface{}) {
	queue := library.Queue{}
	queue.Add([]int{i, j})
	vis[i][j] = true
	xDir := []int{0, 1, 0, -1}
	yDir := []int{-1, 0, 1, 0}
	for !queue.IsEmpty() {
		pop, _ := queue.Remove()
		i = pop.([]int)[0]
		j = pop.([]int)[1]
		for k := range xDir {
			x := i + xDir[k]
			y := j + yDir[k]
			if library.IsValidIndexInMatrix(x, y, n, m) && grid[x][y] == 1 && !vis[x][y].(bool) {
				vis[x][y] = true
				queue.Add([]int{x, y})
			}

		}
	}

}
