package solution

import "sort"

type pair struct {
	Ind, Freq int
}

func FrequencySort(s string) string {
	ans := ""
	freqArr := make([]pair, 75)
	for i := range s {
		ind := int(s[i] - 48)
		// fmt.Println(ind)
		freqArr[ind].Ind = ind
		freqArr[ind].Freq++
	}

	sort.Slice(freqArr, func(i, j int) bool {
		return freqArr[i].Freq < freqArr[j].Freq
	})
	for i := range freqArr {
		for freqArr[74-i].Freq > 0 {
			ans += string(48 + freqArr[74-i].Ind)
			freqArr[74-i].Freq--
		}
	}
	return ans
}
