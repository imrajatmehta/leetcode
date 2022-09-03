package solution

import (
	"math"
)

func NumsSameConsecDiff(n int, k int) []int {
	var ans []int
	var queue []int = make([]int, 0)
	for i := 1; i <= 9; i++ {
		queue = append(queue, i)
	}
	numLen := int(math.Pow(10, float64(n-1)))

	for len(queue) != 0 {
		value := queue[0]
		queue = queue[1:]
		if numLen <= value {
			ans = append(ans, value)
			continue
		}
		lastDigit := value % 10
		value *= 10
		add := lastDigit + k
		substract := lastDigit - k
		if add < 10 {
			queue = append(queue, value+add)
		}
		if substract >= 0 && add != substract {
			queue = append(queue, value+substract)
		}

	}

	return ans
}
