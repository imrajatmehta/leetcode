package solution

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func MaxPerformance(n int, speed []int, efficiency []int, k int) int {
	maxPerformance := 0

	pairArr := make([][]int, 0)
	for i, val := range efficiency {
		pairArr = append(pairArr, []int{val, speed[i]})
	}
	sort.Slice(pairArr, func(i, j int) bool {
		return pairArr[i][0] > pairArr[j][0]

	})
	h := &IntHeap{}
	heap.Init(h)
	sum := 0
	for i, val := range pairArr {
		currSpeed := pairArr[i][1]
		sum += currSpeed
		multiplier := val[0]
		if maxPerformance < sum*multiplier {
			maxPerformance = sum * multiplier
		}
		heap.Push(h, currSpeed)
		if h.Len() == k {
			sa := heap.Pop(h).(int)
			sum -= sa
		}

	}

	return maxPerformance

}
