package solution

var romanMapping map[byte]int = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func RomanToInt(s string) int {

	n := len(s)
	convertedNo := 0

	for i := 0; i < n; i++ {
		if i < n-1 && romanMapping[s[i]] < romanMapping[s[i+1]] {
			convertedNo -= romanMapping[s[i]]
		} else {
			convertedNo += romanMapping[s[i]]
		}

	}

	return convertedNo
}
