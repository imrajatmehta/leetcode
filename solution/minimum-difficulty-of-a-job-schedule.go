package solution

import "math"

func MinDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, d+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	return minDifficultyHelper(0, d, jobDifficulty, dp)

}
func minDifficultyHelper(i, d int, arr []int, dp [][]int) int {
	if i >= len(arr) {
		return 0
	}
	if dp[i][d] != math.MaxInt {
		return dp[i][d]
	}
	max := 0
	ans := math.MaxInt

	for k := i; k <= len(arr)-d; k++ {
		if max < arr[k] {
			max = arr[k]
		}
		if d == 1 {
			ans = max
		} else {
			rightAns := minDifficultyHelper(k+1, d-1, arr, dp)
			if ans > rightAns+max {
				ans = rightAns + max
			}
		}

	}
	dp[i][d] = ans
	return ans

}
