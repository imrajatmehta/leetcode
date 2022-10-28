package solution

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		temp := strings.Join(s, "")

		if _, ok := m[temp]; !ok {
			m[temp] = []string{}
		}
		m[temp] = append(m[temp], str)
	}
	ans := [][]string{}
	for _, val := range m {
		ans = append(ans, val)
	}
	return ans
}
