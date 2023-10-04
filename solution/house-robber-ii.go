package solution

import (
	"github.com/imrajatmehta/leetcode/library"
)

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	library.FillInt(dp, -1)
	a := RobUtil(0, n-1, nums, dp)
	library.FillInt(dp, -1)
	b := RobUtil(1, n, nums, dp)
	if a > b {
		return a
	}
	return b
}
func RobUtil(i, n int, nums, dp []int) int {
	if i >= n {
		return 0
	}
	if dp[i] != -1 {
		return dp[i]
	}
	b := RobUtil(i+1, n, nums, dp)
	a := nums[i] + RobUtil(i+2, n, nums, dp)
	if a > b {
		dp[i] = a
	} else {
		dp[i] = b
	}
	return dp[i]
}
