package solution

func InsertInterval(in [][]int, newI []int) [][]int {
	i := 0
	done := false
	overlapS := newI[0]
	ans := [][]int{}
	for i < len(in) {
		start := newI[0]
		end := newI[1]
		currS := in[i][0]
		currE := in[i][1]
		if done {
			ans = append(ans, []int{in[i][0], in[i][1]})
			i++
		} else {
			if end < currS {
				ans = append(ans, []int{start, end})
				ans = append(ans, in[i])
				i++
				done = true
			} else if currE < start {
				ans = append(ans, in[i])
				i++
			} else {

				if start < currS {
					overlapS = start
				} else {
					overlapS = currS
				}
				for i < len(in) {
					currS := in[i][0]
					currE := in[i][1]
					if end < currS {
						ans = append(ans, []int{overlapS, end})
						done = true
						break
					} else if end <= currE {
						ans = append(ans, []int{overlapS, currE})
						done = true
						i++
						break
					}
					i++

				}

			}

		}

	}
	if !done {
		ans = append(ans, []int{overlapS, newI[1]})
	}
	return ans
}
