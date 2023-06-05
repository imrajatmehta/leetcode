package library

import "math"

func GetDistanceBtwTwoPointsOnPlane(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}
func TwoCirclesIntersects(x1, y1, x2, y2, r1, r2 float64) bool {
	dist := GetDistanceBtwTwoPointsOnPlane(x1, y1, x2, y2)
	if dist <= r1-r2 || dist <= r2-r1 || dist <= r1+r2 {
		return true
	}
	return false
}
