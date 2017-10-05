package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}

	return q.head.value
}

func (q *Queue) Add(value int) {
	temp := &Node{value}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *Queue) Remove() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}

	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

func (q *Queue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value, " ")
		temp = temp.next
	}
}

func main() {
	q := new(Queue)
	for i := 1; i <= 10; i++ {
		q.Add(i)
	}

	for i := 1; i <= 5; i++ {
		q.Remove()
	}

	q.Print()
}
