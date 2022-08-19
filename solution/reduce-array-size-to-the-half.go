package solution

import (
	"sort"
)

func MinSetSize(arr []int) int {

	n := 100001
	var freqCount []int = make([]int, n)
	for _, val := range arr {
		freqCount[val]++
	}
	sort.Ints(freqCount)
	half := len(arr) / 2
	for i := n - 1; i >= 0; i-- {
		half -= freqCount[i]
		if half <= 0 {
			return n - i

		}
	}

	return 0
}
