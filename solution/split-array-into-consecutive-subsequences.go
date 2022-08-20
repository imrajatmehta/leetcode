package solution

import "fmt"

func IsPossible(nums []int) bool {
	var freq map[int]int = map[int]int{}
	var want map[int]int = map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	fmt.Println(freq)
	for _, num := range nums {
		if freq[num] <= 0 {
			continue
		}
		if want[num] > 0 {
			freq[num]--
			want[num]--
			want[num+1]++
		} else if freq[num] > 0 && freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			want[num+3]++
		} else {
			return false
		}
	}

	return true
}
