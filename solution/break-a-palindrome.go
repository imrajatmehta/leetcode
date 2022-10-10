package solution

func BreakPalindrome(palindrome string) string {
	list := []rune(palindrome)
	changed := false
	oddLen := len(palindrome)%2 != 0

	for i, _ := range list {
		if oddLen && i == len(list)/2 {
			continue
		}
		if list[i] != 'a' {
			list[i] = 'a'
			changed = true
			break
		}

	}
	if !changed {

		for i := len(list) - 1; i >= 0; i-- {
			if oddLen && i == len(list)/2 {
				continue
			}
			if list[i] != 'z' {
				list[i]++
				changed = true
				break
			}

		}
	}
	if changed {

		return string(list)
	}
	return ""
}
