package solution

import "github.com/imrajatmehta/leetcode/library"

func HalvesAreAlike(s string) bool {
	owel1 := 0
	owel2 := 0
	n := len(s)
	for i := range s {
		if library.IsOwel(s[i]) {
			if n/2 > i {
				owel1++
			} else {
				owel2++
			}

		}
	}
	return owel1 == owel2
}
