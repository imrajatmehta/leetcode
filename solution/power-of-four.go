package solution

func IsPowerOfFour(n int) bool {

	return n > 0 && n&(n-1) == 0 && 0x55555555&n != 0
}
