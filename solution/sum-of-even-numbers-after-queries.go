package solution

func SumEvenAfterQueries(nums []int, queries [][]int) []int {
	ans := make([]int, 0)
	totalSum := 0
	for _, val := range nums {
		if val%2 == 0 {
			totalSum += val
		}
	}
	for _, val := range queries {
		ind := val[1]
		newVal := val[0]
		if nums[ind]%2 == 0 {
			if (nums[ind]+newVal)%2 == 0 {
				totalSum += newVal

			} else {
				totalSum -= nums[ind]
			}
		} else if (nums[ind]+newVal)%2 == 0 {
			totalSum += nums[ind] + newVal
		}
		ans = append(ans, totalSum)
		nums[ind] += newVal
	}
	return ans

}
