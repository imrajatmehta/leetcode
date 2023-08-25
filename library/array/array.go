package array

import "math"

func GetMin(arr []int) int {
	min := math.MaxInt
	for i := range arr {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}
func GetMax(arr []int) int {
	max := math.MinInt
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
