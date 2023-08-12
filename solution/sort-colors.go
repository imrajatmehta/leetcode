package solution

func SortColors(nums []int) {
	zeroInd := 0
	twoInd := len(nums) - 1
	for i := 0; i < len(nums) && i <= twoInd; {

		for zeroInd < i && nums[i] == 0 {
			nums[i], nums[zeroInd] = nums[zeroInd], nums[i]
			zeroInd++
		}

		for twoInd > i && nums[i] == 2 {
			nums[i], nums[twoInd] = nums[twoInd], nums[i]
			twoInd--
		}

		if !(nums[i] == 0 && zeroInd < i || nums[i] == 2 && twoInd > i) {
			i++
		}
	}
}
