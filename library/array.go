package library

func FillBool(arr []bool, flag bool) []bool {
	for i := range arr {
		arr[i] = flag
	}
	return arr
}
func FillInt(arr []int, flag int) []int {
	for i := range arr {
		arr[i] = flag
	}
	return arr
}
