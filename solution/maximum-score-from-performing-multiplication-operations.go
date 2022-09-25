package solution

import "github.com/imrajatmehta/leetcode/library"

func MaximumScore(nums []int, multipliers []int) int {

	dp := make([][]int, len(multipliers))
	for i, _ := range dp {
		dp[i] = make([]int, len(multipliers))
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	return maximumScoreDfs(0, 0, nums, multipliers, dp)

}
func maximumScoreDfs(i, left int, nums []int, multipliers []int, dp [][]int) int {
	if i >= len(multipliers) {
		return 0
	}
	if dp[i][left] != -1 {
		return dp[i][left]
	}
	ans := 0
	ans = multipliers[i]*nums[0] + maximumScoreDfs(i+1, left+1, nums, multipliers, dp)
	ans = library.MaxInt(multipliers[i]*nums[len(nums)-1-left]+maximumScoreDfs(i+1, left, nums, multipliers, dp), ans)
	dp[i][left] = ans
	return ans

}
