package solution

import "bytes"

func EqualPairs(grid [][]int) int {
	ans := 0
	colMap := make(map[int]string)
	rowMap := make(map[int]string)
	for i := range grid {
		var buffer bytes.Buffer
		for j := range grid[i] {
			buffer.WriteString(string(grid[i][j]))
		}
		colMap[i] = buffer.String()
	}
	for j := range grid[0] {
		var buffer bytes.Buffer
		for i := range grid {
			buffer.WriteString(string(grid[i][j]))
		}
		rowMap[j] = buffer.String()
	}
	for i := range grid {
		for j := range grid[i] {
			if colMap[i] == rowMap[j] {
				ans++
			}
		}

	}
	return ans
}
