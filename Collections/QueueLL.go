package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	que *list.List
}

func NewQueue() *Queue {
	p := new(Queue)
	p.que = list.New()
	return p
}

func (q *Queue) Add(data int) {
	q.que.PushBack(data)
}

func (q *Queue) Remove() int {
	front := q.que.Front()
	val := front.Value.(int)
	q.que.Remove(front)
	return val
}

func (q *Queue) Front() int {
	return q.que.Front().Value.(int)
}

func (q Queue) Len() int {
	return q.que.Len()
}

func (q Queue) IsEmpty() bool {
	return q.que.Len() == 0
}

func (q Queue) Print() {
	for e := q.que.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {
	que := list.New()
	for i := 0; i < 5; i++ {
		que.PushBack(i)
	}

	for que.Len() > 0 {
		front := que.Front()
		fmt.Print(front.Value, " ")
		que.Remove(front)
	}

	fmt.Println()
	q := NewQueue()
	for i := 0; i < 5; i++ {
		q.Add(i)
	}
	q.Print()
	for q.Len() > 0 {
		fmt.Print(q.Remove(), " ")
	}

}

/*
0 1 2 3 4
*/
