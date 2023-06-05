package solution

// Finding slope using  y2-y1/x2-x1 will not work, coz x2-x1 can be 0, and also if one x or y is at 0 then ti will alsways be 0, so that why we are using
// ^Y1/^X1 == ^Y2/^X2 for 3 points to be straight. and to avoid above case, we will be ^Y1*^X2 == ^Y2*^X1.
// can debug this [[0,0],[0,5],[5,0]]
// ^X1=0
// ^Y1=5
// ^X2=5
// ^Y2=0

func CheckStraightLine(coordinates [][]int) bool {
	deltaX := getDiffX(coordinates[0], coordinates[1])
	deltaY := getDiffY(coordinates[0], coordinates[1])
	for i := 2; i < len(coordinates); i++ {
		if deltaX*getDiffY(coordinates[0], coordinates[i]) != deltaY*getDiffX(coordinates[0], coordinates[i]) {
			return false
		}
	}
	return true
}
func getDiffX(cord1, cord2 []int) int {
	return cord2[0] - cord1[0]
}
func getDiffY(cord1, cord2 []int) int {
	return cord2[1] - cord1[1]
}
