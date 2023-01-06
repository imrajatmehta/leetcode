package solution

import "sort"

func MaxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	count := 0
	for i := range costs {
		coins -= costs[i]
		if coins < 0 {
			return count
		}
		count++
	}
}
