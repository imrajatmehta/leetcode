package solution

func reverseVowels(s string) string {
	list := []rune(s)
	i := 0
	j := len(list) - 1
	for i < j {
		firstVowel := findNextVowel(list, i, true)
		lastVowel := findNextVowel(list, j, false)

		if firstVowel != -1 && lastVowel != -1 && firstVowel < lastVowel {
			list[firstVowel], list[lastVowel] = list[lastVowel], list[firstVowel]
			i = firstVowel + 1
			j = lastVowel - 1
		} else {
			return string(list)
		}

	}
	return string(list)
}
func findNextVowel(list []rune, i int, forward bool) int {
	for i < len(list) && i >= 0 {
		if list[i] == 'u' || list[i] == 'o' || list[i] == 'i' || list[i] == 'e' || list[i] == 'a' || list[i] == 'U' || list[i] == 'O' || list[i] == 'I' || list[i] == 'E' || list[i] == 'A' {
			return i
		}
		if forward {
			i++
		} else {
			i--
		}
	}
	return -1
}
