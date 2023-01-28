package library

func IsValidIndexInMatrix(i, j, n, m int) bool {
	if i < 0 || i >= n || j >= m || j < 0 {
		return false
	}
	return true
}
func CreateAndFillMatrixInt(n, m, val int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < m; j++ {
			mat[i][j] = val
		}
	}
	return mat
}
func RemoveIndex(arr [][]int, i int) [][]int {
	return append(arr[:i], arr[i+1:]...)
}
func InsertAtIndex(a [][]int, i, c1, c2 int) [][]int {
	x := make([][]int, 1)
	x[0] = []int{c1, c2}
	return append(a[:i], append(x, a[i:]...)...)
}
