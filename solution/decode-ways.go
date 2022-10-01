package solution

import (
	"fmt"
	"strconv"
)

func NumDecodings(s string) int {
	dp := make([]int, len(s))
	for i, _ := range dp {
		dp[i] = -1
	}
	return numDecodingsHelper(s, 0, dp)
}

func numDecodingsHelper(s string, i int, dp []int) int {
	if len(s) == i {
		return 1
	}
	if dp[i] != -1 {
		return dp[i]
	}
	max := 0
	no1, _ := strconv.Atoi(s[i : i+1])
	fmt.Println(no1)
	if no1 <= 26 && no1 > 0 {
		fmt.Println("HEY " + s[i:i+1])
		max += numDecodingsHelper(s, i+1, dp)
	}
	if i+1 < len(s) {
		no2, _ := strconv.Atoi(s[i : i+2])
		fmt.Println(no2)
		if no2 <= 26 && no2 > 0 && s[i : i+2][0] != '0' {
			fmt.Println("HEY2 " + s[i:i+2])
			max += numDecodingsHelper(s, i+2, dp)
		}

	}
	dp[i] = max
	return dp[i]

}
