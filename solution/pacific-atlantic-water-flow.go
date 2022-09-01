package solution

type Ocean struct {
	Pacific  bool
	Atlantic bool
}

func PacificAtlantic(heights [][]int) [][]int {
	ans := [][]int{}
	n, m := len(heights), len(heights[0])
	var freq [][]Ocean = make([][]Ocean, n)
	for i, _ := range freq {
		freq[i] = make([]Ocean, m)
	}
	for i := 0; i < n; i++ {

		if !freq[i][0].Pacific {

			dfsForPacificAtlantic(freq, heights, i, 0, n, m, true)

		}

		if !freq[i][m-1].Atlantic {

			dfsForPacificAtlantic(freq, heights, i, m-1, n, m, false)
		}

	}
	for j := 0; j < m; j++ {
		if !freq[0][j].Pacific {

			dfsForPacificAtlantic(freq, heights, 0, j, n, m, true)

		}

		if !freq[n-1][j].Atlantic {

			dfsForPacificAtlantic(freq, heights, n-1, j, n, m, false)
		}

	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if freq[i][j].Atlantic && freq[i][j].Pacific {
				ans = append(ans, []int{i, j})
			}

		}

	}

	return ans

}

func dfsForPacificAtlantic(freq [][]Ocean, heights [][]int, i, j, n, m int, isPacific bool) {
	if i >= n || j >= m || i < 0 || j < 0 {
		return
	}
	if isPacific {
		freq[i][j].Pacific = true
		if (i+1) < n && heights[i+1][j] >= heights[i][j] && !freq[i+1][j].Pacific {
			dfsForPacificAtlantic(freq, heights, i+1, j, n, m, isPacific)
		}
		if (i-1) >= 0 && heights[i-1][j] >= heights[i][j] && !freq[i-1][j].Pacific {
			dfsForPacificAtlantic(freq, heights, i-1, j, n, m, isPacific)
		}
		if (j+1) < m && heights[i][j+1] >= heights[i][j] && !freq[i][j+1].Pacific {
			dfsForPacificAtlantic(freq, heights, i, 1+j, n, m, isPacific)
		}
		if (j-1) >= 0 && heights[i][j-1] >= heights[i][j] && !freq[i][j-1].Pacific {
			dfsForPacificAtlantic(freq, heights, i, j-1, n, m, isPacific)
		}
	} else {
		freq[i][j].Atlantic = true
		if (i+1) < n && heights[i+1][j] >= heights[i][j] && !freq[i+1][j].Atlantic {
			dfsForPacificAtlantic(freq, heights, i+1, j, n, m, isPacific)
		}
		if (i-1) >= 0 && heights[i-1][j] >= heights[i][j] && !freq[i-1][j].Atlantic {
			dfsForPacificAtlantic(freq, heights, i-1, j, n, m, isPacific)
		}
		if (j+1) < m && heights[i][j+1] >= heights[i][j] && !freq[i][j+1].Atlantic {
			dfsForPacificAtlantic(freq, heights, i, 1+j, n, m, isPacific)
		}
		if (j-1) >= 0 && heights[i][j-1] >= heights[i][j] && !freq[i][j-1].Atlantic {
			dfsForPacificAtlantic(freq, heights, i, j-1, n, m, isPacific)
		}
	}

}
