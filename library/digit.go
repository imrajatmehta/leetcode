package library

func LenOfDigits(n int) int {
	var tot int
	for n != 0 {
		tot++
		n /= 10
	}
	return tot
}
func MinInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
