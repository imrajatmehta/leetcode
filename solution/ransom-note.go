package solution

func CanConstruct(ransomNote string, magazine string) bool {
	var freq []int = make([]int, 26)
	for _, val := range magazine {
		freq[val-97]++
	}
	for _, val := range ransomNote {
		freq[val-97]--
	}
	for _, val := range freq {
		if val < 0 {
			return false
		}
	}
	return true
}
