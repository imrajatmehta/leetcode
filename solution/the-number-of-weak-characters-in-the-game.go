package solution

import (
	"sort"
)

func NumberOfWeakCharacters(properties [][]int) int {
	ans := 0
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})
	maxDefence := properties[0][1]
	for _, val := range properties {

		if maxDefence > val[1] {
			ans++
		} else {
			maxDefence = val[1]
		}
	}
	return ans

}
