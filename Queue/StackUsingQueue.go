package main

import "fmt"

type StackUsingQueue struct {
	que1 Queue
	que2 Queue
	size int
}

// Push adds an element to the stack.
func (stk *StackUsingQueue) Push(value int) {
	stk.que1.Add(value)
	stk.size++
}

// Pop removes and returns the element at the top of the stack.
func (stk *StackUsingQueue) Pop() int {
	var value int
	s := stk.que1.Length()
	for s > 0 {
		value = stk.que1.Remove().(int)
		if s > 1 {
			stk.que2.Add(value)
		}
		s--
	}
	temp := stk.que1
	stk.que1 = stk.que2
	stk.que2 = temp
	stk.size--
	return value
}

// Pop2 removes and returns the element at the top of the stack.
func (stk *StackUsingQueue) Pop2() int {
	var value int
	s := stk.que1.Length()
	for s > 0 {
		value = stk.que1.Remove().(int)
		if s > 1 {
			stk.que1.Add(value)
		}
		s--
	}
	stk.size--
	return value
}

func main1() {
	stk := new(StackUsingQueue)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	fmt.Println("Stack pop:", stk.Pop())
	fmt.Println("Stack pop:", stk.Pop())
}

func main2() {
	stk := new(StackUsingQueue)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	fmt.Println("Stack pop:", stk.Pop2())
	fmt.Println("Stack pop:", stk.Pop2())
}

func main() {
	main1()
	main2()
}

/*
Stack pop: 3
Stack pop: 2


Stack pop: 3
Stack pop: 2

*/

type Queue struct {
	que []interface{}
}

// Add adds an element to the queue.
func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

// Remove removes and returns the element at the front of the queue.
func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

// RemoveBack removes and returns the element at the back of the queue.
func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

// Front returns the element at the front of the queue without removing it.
func (q *Queue) Front() interface{} {
	return q.que[0]
}

// Back returns the element at the back of the queue without removing it.
func (q *Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}

// Length returns the number of elements in the queue.
func (q *Queue) Length() int {
	return len(q.que)
}

// Print prints the elements in the queue.
func (q Queue) Print() {
	fmt.Println(q.que)
}
