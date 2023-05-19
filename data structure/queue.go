package queue

import "fmt"

type Queue []interface{}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue(data interface{}) (interface{}, error) {
	if q.isEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	} else {
		top := len(*q) - 1
		out := (*q)[0]
		*q = (*q)[1:top]
		return out, nil
	}
}
