package solution

func setZeroes(matrix [][]int) {
	firstRowZero := false
	firstColZero := false
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				if i == 0 {
					firstRowZero = true
				}
				if j == 0 {
					firstColZero = true
				}
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	for i := 0; firstColZero && i < len(matrix); i++ {
		matrix[i][0] = 0
	}

	for i := 0; firstRowZero && i < len(matrix[0]); i++ {
		matrix[0][i] = 0
	}

}
