package solution

func FindDuplicate1(nums []int) int {
	var sum int64 = 0
	for _, val := range nums {
		sum += int64(val)
	}
	var n int64 = int64(len(nums)) - 1
	return int(sum - (n*(n+1))/2)
}
func FindDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for ok := true; ok; ok = slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0] //head
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
