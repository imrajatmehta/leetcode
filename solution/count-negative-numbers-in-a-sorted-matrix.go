package solution

func CountNegatives(grid [][]int) int {
	ans := 0
	n := len(grid)
	m := len(grid[0])
	i := 0
	j := m - 1
	for i < n {
		for j >= 0 && grid[i][j] < 0 {
			j--
		}

		ans += m - j - 1
		i++
	}
	return ans
}
