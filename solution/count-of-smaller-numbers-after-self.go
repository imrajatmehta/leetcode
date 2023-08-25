package solution

import (
	"fmt"

	"github.com/imrajatmehta/leetcode/library/array"
	segmenttree_sum "github.com/imrajatmehta/leetcode/library/segmenttree"
)

func CountSmaller(nums []int) []int {
	min := array.GetMin(nums)
	max := array.GetMax(nums)
	seg := segmenttree_sum.Build(min, max)
	n := len(nums)
	ans := make([]int, n)
	seg.Update(1, nums[n-1])

	for i := 1; i < n; i++ {
		ele := nums[n-i-1]
		fmt.Println(ele)
		seg.Update(1, ele)
		ans[n-i-1] = seg.GetSum(min, ele-1)
	}
	return ans
}
