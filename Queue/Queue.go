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
	que := new(Queue)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println("IsEmpty :", que.IsEmpty());
    fmt.Println("Size :", que.Len());
    fmt.Println("Queue remove :", que.Remove());
    fmt.Println("Queue remove :", que.Remove());

}

/*
IsEmpty : false
Size : 3
Queue remove : 1
Queue remove : 2
*/
