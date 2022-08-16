package solution

func FirstUniqChar(s string) int {
	n := len(s)
	freq := [26]int{}

	for pos := range s {
		freq[s[n-pos-1]-97]++

	}
	for pos := range s {
		if freq[s[pos]-97] == 1 {
			return pos
		}
	}

	return -1
}
