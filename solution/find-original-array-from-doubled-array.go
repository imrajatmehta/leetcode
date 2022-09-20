package solution

func FindOriginalArray(changed []int) []int {
	cache := make(map[int]int)
	ans := make([]int, 0, len(changed))

	for _, val := range changed {
		cache[val]++
	}

	for val := 0; val <= 100000; val++ {

		for cache[val] > 0 {
			ans = append(ans, val)
			cache[val]--
			cache[val*2]--

			if cache[val*2] < 0 {
				return []int{}
			}

		}

	}

	// fmt.Println(ans, cache)

	return ans

}
