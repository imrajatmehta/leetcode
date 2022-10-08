package solution

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	min := math.MaxInt
	for i, _ := range nums {

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[l] + nums[r] + nums[i]

			if AbsInt(target-sum) < min {
				min = AbsInt(target - sum)
				ans = sum
			}
			if sum <= target {
				l++
			} else {
				r--
			}

		}

	}
	return ans
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
