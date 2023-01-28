package solution

import (
	"fmt"

	"github.com/imrajatmehta/leetcode/library"
)

type SummaryRanges struct {
	NewArr [][]int
	LHash  map[int]int
	RHash  map[int]int
}

func SummaryConstructor() SummaryRanges {
	obj := SummaryRanges{}
	obj.NewArr = make([][]int, 0)
	obj.LHash = make(map[int]int)
	obj.RHash = make(map[int]int)
	return obj
}

func (this *SummaryRanges) AddNum(val int) {
	fmt.Println("Adding ", val)
	ind1, ind2 := this.findIndSummaryRanges(val)
	fmt.Println(ind1, ind2)
	if ind1 != -1 && ind2 != -1 {
		delete(this.RHash, this.NewArr[ind1][1])
		delete(this.LHash, this.NewArr[ind2][0])
		this.NewArr[ind1][1] = this.NewArr[ind2][1]
		this.NewArr = library.RemoveIndex(this.NewArr, ind2)

	} else if ind1 != -1 {
		delete(this.LHash, this.NewArr[ind1][0])
		this.NewArr[ind1][0] = val
		this.LHash[val] = ind1
	} else if ind2 != -1 {
		delete(this.RHash, this.NewArr[ind2][1])
		this.NewArr[ind2][1] = val
		this.RHash[val] = ind2
	} else {
		fmt.Println("Etnered")

		if i, isLap := this.OverLapIndex(val); !isLap {
			fmt.Println(i, isLap)
			this.NewArr = library.InsertAtIndex(this.NewArr, i, val, val)
			fmt.Println(this.NewArr)
			this.LHash[val] = i
			this.RHash[val] = i
		}
	}

}

func (this *SummaryRanges) GetIntervals() [][]int {

	fmt.Println(this.LHash, this.RHash)
	return this.NewArr
}
func (this *SummaryRanges) findIndSummaryRanges(val int) (int, int) {
	ind1 := -1
	ind2 := -1
	if _, ok := this.LHash[val+1]; ok {
		fmt.Println("ind2 = findEleSortedMatrix")
		ind2 = findEleSortedMatrix(this.NewArr, val+1, true)
	}
	if _, ok := this.RHash[val-1]; ok {
		fmt.Println("ind1 = findEleSortedMatrix")
		ind1 = findEleSortedMatrix(this.NewArr, val-1, false)
	}
	if ind2 == -1 || ind1 == -1 {
		return ind2, ind1
	} else {
		return ind1, ind2
	}

}
func (this *SummaryRanges) OverLapIndex(val int) (int, bool) {
	l := 0
	r := len(this.NewArr) - 1
	for l <= r {
		mid := l + (r-l)/2
		if this.NewArr[mid][0] <= val && val <= this.NewArr[mid][1] {
			return mid, true
		} else if val < this.NewArr[mid][0] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l, false
}
func findEleSortedMatrix(arr [][]int, val int, firstInd bool) int {
	r := len(arr) - 1
	l := 0
	for l <= r {
		mid := l + (r-l)/2
		// fmt.Println(mid, val)
		x := 0
		if firstInd {
			x = arr[mid][0]
		} else {
			x = arr[mid][1]
		}

		if x == val {
			return mid
		} else if x < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
