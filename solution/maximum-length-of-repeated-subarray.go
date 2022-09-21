package solution

var res int

func FindLength(nums1 []int, nums2 []int) int {
	res = 0
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	dfsFindLength(0, 0, nums1, nums2, dp)
	return res
}

func dfsFindLength(i, j int, nums1 []int, nums2 []int, dp [][]int) int {
	if i >= len(nums1) || j >= len(nums2) {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	ans := 0
	if nums1[i] == nums2[j] {
		ans = 1 + dfsFindLength(i+1, j+1, nums1, nums2, dp)
		if res < ans {
			res = ans
		}
	}
	dfsFindLength(i+1, j, nums1, nums2, dp)
	dfsFindLength(i, j+1, nums1, nums2, dp)
	dp[i][j] = ans
	return dp[i][j]
}
