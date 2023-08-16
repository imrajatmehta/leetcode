package solution

import "github.com/imrajatmehta/leetcode/library"

func uniquePaths(m int, n int) int {
	dp := library.CreateAndFillMatrixInt(m, n, -1)
	return uniquePathsDFS(0, 0, m, n, dp)
}
func uniquePathsDFS(i, j, m, n int, dp [][]int) int {
	if !library.IsValidIndexInMatrix(i, j, m, n) {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if i == m-1 && j == n-1 {
		return 1
	}
	dp[i][j] = uniquePathsDFS(i+1, j, m, n, dp) + uniquePathsDFS(i, j+1, m, n, dp)
	return dp[i][j]
}
