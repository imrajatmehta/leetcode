package solution

import "strconv"

func CountAndSay(n int) string {
	ans := "1"
	for n > 1 {
		ans = countAndSayHelper(ans)
		n--
	}
	return ans
}
func countAndSayHelper(no string) string {
	count := 1
	prev := no[0]
	newNo := ""
	for i := 1; i < len(no); i++ {
		digit := no[i]
		if prev == digit {
			count++
		} else {
			newNo += strconv.Itoa(count) + string(prev)
			count = 1
			prev = digit
		}
	}

	newNo += strconv.Itoa(count) + string(prev)
	return newNo
}
