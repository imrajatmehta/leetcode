package library

func LenOfDigits(n int) int {
	var tot int
	for n != 0 {
		tot++
		n /= 10
	}
	return tot
}
