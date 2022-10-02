package solution

import (
	"fmt"
)

func PushDominoes(dominoes string) string {
	list := []rune(dominoes)

	left := make([]int, len(list))

	n := len(list) - 1
	if list[n] == 'L' {
		left[n] = n
	} else {
		left[n] = -1
	}
	for i := n - 1; i >= 0; i-- {
		if list[i] == 'L' {
			left[i] = i
		} else if list[i] == 'R' {
			left[i] = -1
		} else {
			left[i] = left[i+1]
		}

	}
	fmt.Println(string(list))
	i := 0
	for i = 0; i <= n; i++ {
		val := list[i]
		if val == 'R' {
			if i < n && left[i+1] != -1 {
				diff := left[i+1] - i - 1
				fmt.Println(diff)

				j := diff / 2
				fmt.Println(j)
				fillR(&list, i, j)
				if diff%2 == 0 {
					i += j
				} else {
					i += j + 1
				}
			}

		} else if val == '.' {

			if left[i] != -1 {
				list[i] = 'L'
			} else if i != 0 && list[i-1] == 'R' {
				list[i] = 'R'
			}
		}
		fmt.Println(string(list))

	}

	return string(list)
}
func fillR(list *[]rune, i, times int) {
	for times >= 0 {
		(*list)[i] = 'R'
		i++
		times--
	}

}
