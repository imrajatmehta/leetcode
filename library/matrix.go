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
