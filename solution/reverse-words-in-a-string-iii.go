package solution

import "strings"

func ReverseWords(s string) string {

	sList := strings.Split(s, " ")
	for i, v := range sList {
		sList[i] = reverseString(v)
	}
	return strings.Join(sList, " ")
}
func reverseString(s string) string {
	list := []rune(s)
	i := 0
	n := len(list) - 1
	for i < n {
		list[i], list[n] = list[n], list[i]
		i++
		n--
	}
	return string(list)
}
