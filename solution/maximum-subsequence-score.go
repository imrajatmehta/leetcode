package solution

import (
	"container/heap"
	"sort"
)

func MaxScore(nums1 []int, nums2 []int, k int) int64 {

	var max int64 = 0
	arrs := TwoIntArr{nums2, nums1}
	sort.Sort(SortTwoIntArr(arrs))
	kSum := 0

	maxHeap := MaxHeap{}
	for i := 0; i < k-1; i++ {
		heap.Push(&maxHeap, nums1[i])
		kSum += nums1[i]
	}
	for i := k - 1; i < len(nums1); i++ {
		tem := int64(nums2[i] * (kSum + nums1[i]))
		if tem > max {
			max = tem
		}
		if maxHeap.Len() > 0 {
			poped := heap.Pop(&maxHeap).(int)
			if poped > nums1[i] {
				heap.Push(&maxHeap, poped)
			} else {
				heap.Push(&maxHeap, nums1[i])
				kSum += nums1[i] - poped
			}
		}
	}
	return max

}

type SortTwoIntArr TwoIntArr
type TwoIntArr struct {
	A []int
	B []int
}

func (arrs SortTwoIntArr) Len() int {
	return len(arrs.A)
}

func (arrs SortTwoIntArr) Swap(i, j int) {
	arrs.A[i], arrs.A[j] = arrs.A[j], arrs.A[i]
	arrs.B[i], arrs.B[j] = arrs.B[j], arrs.B[i]
}
func (arrs SortTwoIntArr) Less(i, j int) bool {
	return arrs.A[i] > arrs.A[j]
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
