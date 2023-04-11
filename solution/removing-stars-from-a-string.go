package solution

import "github.com/imrajatmehta/leetcode/library"

func RemoveStars(s string) string {
	n := len(s)
	ans := ""
	star := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '*' {
			star++
		} else {
			if star > 0 {
				star--
				continue
			} else {
				ans += string(s[i])
			}
		}

	}
	return library.Reverse(ans)
}
