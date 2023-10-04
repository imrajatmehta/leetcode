package queue

type Queue struct {
	arr []any
}

func New(size int) Queue {
	return Queue{arr: []any{}}
}
func (q *Queue) Pop() any {
	poped := q.arr[0]
	q.arr = q.arr[1:]
	return poped
}

func (q *Queue) Add(a any) {
	q.arr = append(q.arr, a)
}
func (q *Queue) Front() any {
	return q.arr[0]
}

func (q *Queue) Len() int {
	return len(q.arr)
}
