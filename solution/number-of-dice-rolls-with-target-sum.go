package solution

func NumRollsToTarget(n int, k int, target int) int {

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return numRollsToTargetHelper(n, k, target, dp)
}

var MOD = 1000000000 + 7

func numRollsToTargetHelper(i, k, target int, dp [][]int) int {
	if i <= 0 || target < 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	if dp[i][target] != -1 {
		return dp[i][target]
	}
	ans := 0
	for j := 1; j <= k; j++ {
		ans = (ans + numRollsToTargetHelper(i-1, k, target-j, dp)) % MOD
	}
	dp[i][target] = ans
	return ans
}
