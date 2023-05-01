package main

import (
	"container/list"
	"fmt"
)

// Queue
type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Enqueue(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Dequeue() interface{} {
	front := q.v.Front()
	if front != nil {
		//Remove() => 비우면서 비운값 반환
		return q.v.Remove(front)
	}
	return nil
}

// Stack
type Stack struct {
	v *list.List
}

func newStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	v := list.New()

	e4 := v.PushBack(4)
	e1 := v.PushBack(1)

	v.InsertBefore(3, e4)
	v.InsertBefore(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	qValue := q.Dequeue()

	for qValue != nil {
		fmt.Printf("%v => ", qValue)
		qValue = q.Dequeue()
	}

	fmt.Println()

	s := newStack()
	books := []string{"어린왕자", "겨울왕국", "노인과 바다", "짱구"}

	for i := 0; i < len(books); i++ {
		s.Push(books[i])
	}

	val := s.Pop()
	for val != nil {
		fmt.Printf("%v => ", val)
		val = s.Pop()
	}
}
