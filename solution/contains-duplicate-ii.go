package solution

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	m[nums[0]] = 1
	s := 0
	for i := 1; i < len(nums); i++ {

		if i > k {
			m[nums[s]]--
			s++
		}

		m[nums[i]]++
		val := m[nums[i]]
		if val > 1 {
			return true
		}

	}
	return false
}
