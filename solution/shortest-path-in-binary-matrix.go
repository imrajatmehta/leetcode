package solution

import (
	"math"

	"github.com/imrajatmehta/leetcode/library"
)

func ShortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 {
		return -1
	}
	dp := library.CreateAndFillMatrixInt(len(grid), len(grid), -1)
	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid))
	}

	ans, reached := dfsShortestPathBinaryMatrix(&grid, &dp, 0, 0, &vis)

	if (ans == 0 && len(grid) != 1 && len(grid[0]) != 0) || !reached {
		return -1
	}
	return ans
}
func dfsShortestPathBinaryMatrix(grid, dp *[][]int, i, j int, vis *[][]bool) (int, bool) {
	if (*dp)[i][j] > -1 || (*dp)[i][j] == -2 {
		return (*dp)[i][j], (*vis)[i][j]
	}
	if i == len((*grid))-1 && j == len((*grid))-1 {
		(*vis)[i][j] = true
		(*dp)[i][j] = 0
		return 0, true

	}
	(*vis)[i][j] = false
	(*dp)[i][j] = -2
	dirs := library.GetEightXYDirections()
	min := math.MaxInt
	for k := range dirs {
		x := dirs[k][0] + i
		y := dirs[k][1] + j
		if library.IsValidIndexInMatrix(x, y, len((*grid)[i]), len((*grid)[i])) && (*grid)[x][y] == 0 {
			dist, reached := dfsShortestPathBinaryMatrix(grid, dp, x, y, vis)
			if reached && dist < min {
				min = dist + 1
			}
		}
	}
	if min != math.MaxInt {
		(*dp)[i][j] = min
		(*vis)[i][j] = true
		return min, true
	}

	return (*dp)[i][j], (*vis)[i][j]
}
