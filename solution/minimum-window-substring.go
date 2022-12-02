package solution

func MinWindow(s string, t string) string {
	ans := ""
	freq := make([]int, 26)
	for _, val := range t {
		freq[val-97]++
	}
	i, j := 0, 0
	for j != 0 {

		j++
	}

	return ans
}
