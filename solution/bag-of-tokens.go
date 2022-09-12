package solution

import (
	"sort"
)

func BagOfTokensScore(tokens []int, power int) int {
	maxSrc, src := 0, 0

	sort.Sort((sort.IntSlice(tokens)))
	left, right := 0, len(tokens)-1
	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			src++
			if maxSrc < src {
				maxSrc = src
			}
			left++
		} else if src > 0 {
			src--
			power += tokens[right]
			right--

		} else {
			return src
		}
	}

	return maxSrc

}
