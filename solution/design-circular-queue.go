package solution

type MyCircularQueue struct {
	Queue                           []int
	FrontIndex, LastIndex, Capacity int
}

func Constructor(k int) MyCircularQueue {
	var newObj MyCircularQueue
	newObj.FrontIndex = 0
	newObj.LastIndex = -1
	newObj.Capacity = k
	newObj.Queue = make([]int, k)

	return newObj
}

func (cQ *MyCircularQueue) EnQueue(value int) bool {
	if cQ.IsFull() {
		return false
	}
	cQ.LastIndex++
	cQ.LastIndex %= cQ.Capacity
	cQ.Queue[cQ.LastIndex] = value

	return true
}

func (cQ *MyCircularQueue) DeQueue() bool {
	if cQ.IsEmpty() {
		return false
	}
	if cQ.LastIndex == cQ.FrontIndex {
		cQ.Reset()
	} else {
		cQ.FrontIndex++
		cQ.FrontIndex %= cQ.Capacity
	}

	return true
}

func (cQ *MyCircularQueue) Front() int {
	if cQ.IsEmpty() {
		return -1
	}
	return cQ.Queue[cQ.FrontIndex]
}

func (cQ *MyCircularQueue) Rear() int {
	if cQ.IsEmpty() {
		return -1
	}
	return cQ.Queue[cQ.LastIndex]
}

func (cQ *MyCircularQueue) IsEmpty() bool {
	return cQ.LastIndex == -1
}

func (cQ *MyCircularQueue) IsFull() bool {
	if cQ.LastIndex == -1 {
		return false
	}
	if cQ.LastIndex < cQ.FrontIndex && cQ.Capacity-cQ.FrontIndex+cQ.LastIndex == cQ.Capacity-1 {
		return true
	} else if cQ.LastIndex-cQ.FrontIndex == cQ.Capacity-1 {
		return true
	}
	return false

}
func (cQ *MyCircularQueue) Reset() {
	cQ.FrontIndex = 0
	cQ.LastIndex = -1
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
