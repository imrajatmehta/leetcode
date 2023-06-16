package solution

import "github.com/imrajatmehta/leetcode/library"

var mod int = 1e9 + 7

func NumOfWays(nums []int) int {
	n := len(nums)
	combMat := library.GetRightPascalTriange(n, mod)
	return dfsNumOfWays(nums, combMat) % mod
}

func dfsNumOfWays(nums []int, combMap [][]int) int {
	m := len(nums)
	if m < 3 {
		return 1
	}
	leftSize := make([]int, 0)
	rightSize := make([]int, 0)
	for i := 1; i < m; i++ {
		if nums[i] < nums[0] {
			leftSize = append(leftSize, nums[i])
		} else {
			rightSize = append(rightSize, nums[i])
		}
	}
	leftCr := dfsNumOfWays(leftSize, combMap) % mod
	rightCr := dfsNumOfWays(rightSize, combMap) % mod
	return (((leftCr * rightCr) % mod) * combMap[m-1][len(leftSize)]) % mod
}
