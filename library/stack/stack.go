package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New() *Stack {
	return &Stack{}
}
func (this *Stack) Len() int {
	return this.length
}
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	poped := this.top
	this.top = poped.prev
	this.length--
	return poped.value

}

func (this *Stack) Push(value interface{}) {

	this.length++
	newV := &node{value, this.top}
	this.top = newV

}
