package solution

import "sort"

func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	l := 0
	r := len(people) - 1
	ans := 0
	for l <= r {
		if l == r {
			l++
			r--
		} else if people[l]+people[r] <= limit {
			l++
			r--
		} else {
			r--
		}
		ans++
	}

	return ans
}
