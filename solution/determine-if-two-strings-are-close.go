package solution

import (
	"fmt"
	"sort"
)

func CloseStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	//checking both having same alpha
	f1 := make([]int, 26)
	f2 := make([]int, 26)
	for i := range a {
		ind := a[i] - 'a'
		f1[ind]++
	}
	for i := range b {
		ind := b[i] - 'a'
		f2[ind]++
	}
	for i := range f1 {
		if (f1[i] == 0 && f2[i] > 0) || (f1[i] > 0 && f2[i] == 0) {
			return false
		}
	}
	sort.Ints(f1)
	sort.Ints(f2)
	fmt.Println(f1, f2)
	for i := range f1 {
		if f1[i] != f2[i] {
			return false
		}
	}

	return true
}
