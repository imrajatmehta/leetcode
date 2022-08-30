package solution



func NumIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	var freq [][]bool = make([][]bool, n)
	for i, _ := range freq {
		freq[i] = make([]bool, m)
	}

	var totalIsland int = 0
	for i, arr := range grid {
		for j, val := range arr {
			if !freq[i][j] && val == '1' {
				totalIsland++
				dfs(grid, i, j, n, m, freq)
			}
		}
	}

	return totalIsland

}

func dfs(grid [][]byte, i, j, n, m int, freq [][]bool) {
	if i >= n || j >= m || i < 0 || j < 0 {
		return
	}
	freq[i][j] = true
	
	if (i+1) < n && grid[i+1][j] == '1' && !freq[i+1][j] {
		dfs(grid, i+1, j, n, m, freq)
	}
	if (i-1) >= 0 && grid[i-1][j] == '1' && !freq[i-1][j] {
		dfs(grid, i-1, j, n, m, freq)
	}
	if (j+1) < m && grid[i][j+1] == '1' && !freq[i][j+1] {
		dfs(grid, i, 1+j, n, m, freq)
	}
	if (j-1) >= 0 && grid[i][j-1] == '1' && !freq[i][j-1] {
		dfs(grid, i, j-1, n, m, freq)
	}

}
