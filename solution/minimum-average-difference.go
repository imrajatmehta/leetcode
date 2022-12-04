package solution

import (
	"math"
)

func MinimumAverageDifference(nums []int) int {
	n := len(nums)
	leftSum := make([]float64, n)
	rightSum := make([]float64, n)
	leftSum[0] = float64(nums[0])
	rightSum[n-1] = float64(nums[n-1])
	for i := 1; i < n; i++ {
		leftSum[i] = float64(nums[i]) + leftSum[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		rightSum[i] = float64(nums[i]) + rightSum[i+1]
	}
	ans := math.MaxFloat64
	ansInd := -1
	for i := 0; i < n; i++ {
		leftAvg := leftSum[i] / float64(i+1)
		leftAvg = math.Floor(leftAvg)
		var rightAvg float64
		if i != n-1 {
			rightAvg = rightSum[i+1] / float64(n-i-1)

		} else {
			rightAvg = 0
		}
		rightAvg = math.Floor(rightAvg)

		diff := math.Abs(rightAvg - leftAvg)
		if ans > diff {
			ans = diff
			ansInd = i
		}
	}
	if ansInd == -1 {
		return n - 1
	}
	return ansInd
}
