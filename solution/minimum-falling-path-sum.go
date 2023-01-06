package solution

import (
	"math"

	"github.com/imrajatmehta/leetcode/library"
)

func MinFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	max := math.MaxInt
	dp := library.CreateAndFillMatrixInt(n, m, math.MaxInt)

	for j := 0; j < n; j++ {
		max = library.MinInt(max, goBelow(matrix, dp, 0, j, n, m))
	}
	return max
}
func goBelow(mat, dp [][]int, i, j, n, m int) int {

	min := 0
	if dp[i][j] != math.MaxInt {
		return dp[i][j]
	}
	if library.IsValidIndexInMatrix(i+1, j, n, m) {
		min = goBelow(mat, dp, i+1, j, n, m)
	}
	if library.IsValidIndexInMatrix(i+1, j+1, n, m) {
		min = library.MinInt(goBelow(mat, dp, i+1, j+1, n, m), min)
	}
	if library.IsValidIndexInMatrix(i+1, j-1, n, m) {
		min = library.MinInt(goBelow(mat, dp, i+1, j-1, n, m), min)
	}
	dp[i][j] = mat[i][j] + min
	return dp[i][j]
}
