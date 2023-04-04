package solution

func PartitionString(s string) int {
	ans := 1
	freq := make([]int, 26)
	for i := range s {
		ind := s[i] - 'a'
		if freq[ind] > 0 {
			ans++
			freq = make([]int, 26)
		}
		freq[ind]++

	}
	return ans
}
