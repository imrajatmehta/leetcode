package library

import (
	"errors"
)

type Queue struct {
	Elements []interface{}
}

func (q *Queue) Add(elem interface{}) {
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Remove() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("UnderFlow")
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element, nil
	}
	q.Elements = q.Elements[1:]
	return element, nil
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
