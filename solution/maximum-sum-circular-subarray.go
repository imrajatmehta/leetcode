package solution

import "math"

func MaxSubarraySumCircular(nums []int) int {
	maxPosSum := math.MinInt
	maxNegSum := math.MinInt
	negSum := 0
	posSum := 0
	totalSum := 0

	for i := range nums {
		totalSum += nums[i]
		posSum += nums[i]
		negSum += -nums[i]

		if maxPosSum < posSum {
			maxPosSum = posSum
		}
		if posSum < 0 {
			posSum = 0
		}
		if maxNegSum < negSum {
			maxNegSum = negSum
		}
		if negSum < 0 {
			negSum = 0
		}
	}
	// fmt.Println(totalSum,maxNegSum,maxPosSum)
	if maxPosSum < 0 {
		return maxPosSum
	}
	if totalSum+maxNegSum > maxPosSum {
		return totalSum + maxNegSum
	}
	return maxPosSum
}
