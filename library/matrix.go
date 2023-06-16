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

func CreateAndFillMatrix(n, m int, val interface{}) [][]interface{} {
	mat := make([][]interface{}, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]interface{}, m)
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
func GetRightPascalTriange(n int, mod int) [][]int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, i+1)
		mat[i][0] = 1
		mat[i][i] = 1
	}
	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			mat[i][j] = (mat[i-1][j] + mat[i-1][j-1]) % mod
		}
	}
	return mat

}
