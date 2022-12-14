package solution

import "fmt"

func ClimbStairs(n int) int {
	prev := 1
	curr := 1
	for n > 1 {
		newCurr := prev + curr
		prev = curr
		curr = newCurr
		n--
		fmt.Println(prev, curr)
	}
	return curr
}
