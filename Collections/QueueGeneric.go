package main

import "fmt"

type Queue struct {
	que []interface{}
}

func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func (q *Queue) Front() interface{} {
	return q.que[0]
}

func (q *Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue) Len() int {
	return len(q.que)
}

func (q Queue) Print() {
	fmt.Println(q.que)
}

func main() {
	q := new(Queue)
	for i := 0; i < 5; i++ {
		q.Add(i)
	}

	for q.IsEmpty() == false {
		val := q.Remove()
		fmt.Print(val, " ")
	}
}
