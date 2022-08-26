package solution

func ReorderedPowerOf2(n int) bool {

	var freq []int = make([]int, 10)

	for n != 0 {
		freq[n%10]++
		n /= 10
	}

	for i := 0; i < 33; i++ {
		val := 2 << i

		var temp []int = make([]int, 10)
		copy(temp, freq)
		if checkIfValCanBeMade(temp, val) {
			return true
		}
	}

	return false
}
func checkIfValCanBeMade(arr []int, n int) bool {
	for n != 0 {
		arr[n%10]--
		n /= 10
	}
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}

	return true
}
