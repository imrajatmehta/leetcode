package solution

import "github.com/imrajatmehta/leetcode/library"

func uniquePaths(m int, n int) int {
	dp := library.CreateAndFillMatrixInt(m, n, -1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
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
