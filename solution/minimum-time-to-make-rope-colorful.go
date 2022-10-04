package solution

import "fmt"

func MinCost(colors string, neededTime []int) int {
	n := len(neededTime)
	arr := []rune(colors)
	ans := 0
	for i := 0; i < n; i++ {
		lastInd := findLastInd(i, n, arr)
		fmt.Println(lastInd)
		if lastInd == i {
			continue
		} else {
			ans += sumAfterSubstMax(i, lastInd, neededTime)
		}
		i = lastInd
	}
	return ans

}
func findLastInd(i, n int, arr []rune) int {

	for i < n-1 && arr[i] == arr[i+1] {
		i++
	}
	return i
}
func sumAfterSubstMax(i, j int, arr []int) int {

	max := 0
	sum := 0
	for k := i; k <= j; k++ {
		sum += arr[k]
		if max < arr[k] {
			max = arr[k]
		}
	}
	if sum-max != 0 {
		return sum - max
	}
	return sum

}
