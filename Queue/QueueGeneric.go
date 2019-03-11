package main

import "fmt"

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) Enqueue(value interface{}) {
	if q.size >= capacity {
		return
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) Dequeue() interface{} {
	var value interface{}
	if q.size <= 0 {
		return 0
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}

func main() {
	q := new(Queue)
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 10; i++ {
		val := q.Dequeue().(int)
		fmt.Print(val, " ")
	}
}
