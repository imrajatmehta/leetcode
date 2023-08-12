package solution

func MaxSubArray(nums []int) int {
	gMax := -100000000000
	sum := 0
	l := 0
	for r, val := range nums {
		sum += val
		if sum > gMax {
			gMax = sum
		}

		for l < r && sum < 0 {
			sum -= nums[l]
			l++
			if sum > gMax {
				gMax = sum
			}
		}
		if sum < 0 {
			sum = 0
			l = r + 1
		}
	}
	return gMax
}
