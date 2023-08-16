package solution

import (
	"math"
	"sort"
)

func NextPermutation(nums []int) {
	n := len(nums)
	prev := nums[n-1]
	i := n - 2
	for ; i >= 0; i-- {
		if prev <= nums[i] {
			prev = nums[i]
		} else {
			break
		}
	}
	if i >= 0 {
		j := i + 1
		diff := math.MaxInt
		minI := i

		for ; j < n; j++ {
			if nums[j]-nums[i] < diff && nums[j] > nums[i] {
				minI = j
				diff = nums[j] - nums[i]
			}
		}
		nums[i], nums[minI] = nums[minI], nums[i]
	}
	subset := nums[i+1:]
	sort.Slice(subset, func(a, b int) bool {
		return subset[a] < subset[b]
	})
}
