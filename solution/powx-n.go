package solution

func MyPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / MyPowHelper(x, n)
	}
	return MyPowHelper(x, n)
}
func MyPowHelper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	a := 1.0
	if n%2 != 0 {
		a = x
	}
	ak := MyPowHelper(x, n/2)
	return ak * ak * a
}
