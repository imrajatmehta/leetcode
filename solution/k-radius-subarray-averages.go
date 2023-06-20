package solution

func GetAverages(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n)
	prefSum := make([]int, n+2)
	prefSum[1] = nums[0]
	for i := 1; i < n; i++ {
		prefSum[i+1] = nums[i] + prefSum[i]
	}

	for i := range nums {
		if i < k || i+k >= n {
			ans[i] = -1
		} else {
			leftSum := prefSum[i] - prefSum[i-k]
			rightSum := prefSum[i+k+1] - prefSum[i+1]

			ans[i] = (leftSum + rightSum + nums[i]) / (k*2 + 1)
		}
	}
	return ans
}
