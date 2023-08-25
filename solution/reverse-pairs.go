package solution

import (
	"github.com/imrajatmehta/leetcode/library/array"
	segmenttree_sum "github.com/imrajatmehta/leetcode/library/segmenttree"
)

func ReversePairs(nums []int) int {
	min := array.GetMin(nums) * 2
	max := array.GetMax(nums)
	if max > 0 {
		max *= 2
	}
	seg := segmenttree_sum.Build(min, max)
	n := len(nums)
	ans := 0
	seg.Update(1, 2*nums[n-1])

	for i := 1; i < n; i++ {
		ele := nums[n-i-1]
		ans += seg.GetSum(min, ele-1)
		seg.Update(1, ele*2)
	}
	return ans
}
