package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	que *list.List
}

func NewQueue() *Queue {
	q := new(Queue)
	q.que = list.New()
	return q
}

func (q *Queue) Add(value interface{}) {
	q.que.PushBack(value)
}

func (q *Queue) Remove() interface{} {
	front := q.que.Front()
	val := front.Value
	q.que.Remove(front)
	return val
}

func (q *Queue) Front() interface{} {
	return q.que.Front().Value
}

func (q Queue) Len() int {
	return q.que.Len()
}

func (q Queue) IsEmpty() bool {
	return q.que.Len() == 0
}

func main() {
	q := NewQueue()
	for i := 0; i < 5; i++ {
		q.Add(i)
	}

	for q.IsEmpty() == false {
		val := q.Remove()
		fmt.Print(val, " ")
	}
}
