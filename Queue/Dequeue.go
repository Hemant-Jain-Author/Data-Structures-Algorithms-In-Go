package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev *Node
}

type Dequeue struct {
	head *Node
	tail *Node
	size int
}

func (q *Dequeue) Size() int {
	return q.size
}

func (q *Dequeue) IsEmpty() bool {
	return q.size == 0
}

func (q *Dequeue) PeekFront() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	return q.head.value
}

func (q *Dequeue) PeekBack() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	return q.tail.value
}

func (q *Dequeue) AddBack(value int) {
	temp := &Node{value, nil, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		temp.prev = q.tail
		q.tail = temp
	}
	q.size++
}

func (q *Dequeue) AddFront(value int) {
	temp := &Node{value, nil, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		temp.next = q.head
		q.head.prev = temp
		q.head = temp
	}
	q.size++
}


func (q *Dequeue) RemoveFront() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	q.size--
	value := q.head.value
	q.head = q.head.next
	
	if q.head != nil {
		q.head.prev = nil
	} else {
		q.tail = nil
	}
	
	return value
}

func (q *Dequeue) RemoveBack() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	q.size--
	value := q.tail.value
	q.tail = q.tail.prev
	if q.tail != nil {
		q.tail.next = nil
	} else {
		q.head = nil
	}
	return value
}

func (q *Dequeue) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func main() {
	q := new(Dequeue)
	q.AddFront(1)
	q.AddFront(1)
	q.AddBack(2)
	q.AddBack(2)
	q.Print()
	fmt.Println(q.RemoveBack())
	q.Print()
	fmt.Println(q.RemoveFront())
	q.Print()		
}