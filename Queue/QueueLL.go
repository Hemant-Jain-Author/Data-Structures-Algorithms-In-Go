package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type QueueLinkedList struct {
	tail *Node
	size int
}

// Size returns the number of elements in the queue.
func (q *QueueLinkedList) Size() int {
	return q.size
}

// IsEmpty checks if the queue is empty.
func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

// Peek returns the element at the front of the queue without removing it.
func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	return q.tail.next.value
}

// Add adds an element to the rear of the queue.
func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.tail == nil {
		// If the queue is empty, set the tail to the new node and make it circular.
		q.tail = temp
		temp.next = temp
	} else {
		// Add the new node to the rear and update the tail.
		temp.next = q.tail.next
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

// Remove removes and returns the element at the front of the queue.
func (q *QueueLinkedList) Remove() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}

	// Remove the node at the front and update the tail if necessary.
	value := q.tail.next.value
	if q.tail == q.tail.next {
		q.tail = nil
	} else {
		q.tail.next = q.tail.next.next
	}
	q.size--
	return value
}

// Print prints the elements in the queue.
func (q *QueueLinkedList) Print() {
	if q.IsEmpty() {
		return
	}
	temp := q.tail.next
	for temp != q.tail {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func main() {
	que := new(QueueLinkedList)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println("IsEmpty:", que.IsEmpty())
	fmt.Println("Size:", que.Size())
	fmt.Println("Queue remove:", que.Remove())
	fmt.Println("Queue remove:", que.Remove())
}

/*
IsEmpty: false
Size: 3
Queue remove: 1
Queue remove: 2
*/
