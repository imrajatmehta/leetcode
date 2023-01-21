package solution

import "fmt"

func MinFlipsMonoIncr(s string) int {
	n := len(s)
	leftZero := make([]int, n)
	leftOne := make([]int, n)
	rightZero := make([]int, n)
	rightOne := make([]int, n)

	for i := 0; i < n; i++ {

		if s[i] == '0' {
			leftZero[i] = 1

		} else {
			leftOne[i] = 1
		}
		if i > 0 {
			leftOne[i] += leftOne[i-1]
			leftZero[i] += leftZero[i-1]
			rightZero[n-i-1] = rightZero[n-i]
			rightOne[n-i-1] = rightOne[n-i]
		}

		if s[n-i-1] == '0' {
			rightZero[n-i-1] += 1
		} else {
			rightOne[n-i-1] += 1
		}

	}
	ans := rightOne[0] //if all zero
	fmt.Println(ans)
	if ans > rightZero[0] {
		ans = rightZero[0] //if all one
		fmt.Println(ans)
	}
	for i := range leftOne {
		fmt.Println(leftOne[i], rightZero[i])
		if ans > (leftOne[i] + rightZero[i]) {
			ans = leftOne[i] + rightZero[i]
			fmt.Println(ans)
		}
	}
	return ans
}
