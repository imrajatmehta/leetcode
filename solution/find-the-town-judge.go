package solution

func FindJudge(n int, trust [][]int) int {
	outgoing := make(map[int]int)
	ingoing := make(map[int]int)
	for i := range trust {
		outgoing[trust[i][0]]++
		ingoing[trust[i][1]]++
	}

	for i := 1; i <= n; i++ {
		if outgoing[i] == 0 && ingoing[i] == n-1 {
			return i
		}
	}

	return -1
}
