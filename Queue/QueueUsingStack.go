package main

import "fmt"

type QueueUsingStack struct {
	stk1 Stack
	stk2 Stack
}

// Add adds an element to the queue.
func (que *QueueUsingStack) Add(value int) {
	que.stk1.Push(value)
}

// Remove removes and returns the element at the front of the queue.
func (que *QueueUsingStack) Remove() int {
	var value int
	if que.stk2.IsEmpty() == false {
		value = que.stk2.Pop().(int)
		return value
	}

	// Move elements from stk1 to stk2 to reverse the order.
	for que.stk1.IsEmpty() == false {
		value = que.stk1.Pop().(int)
		que.stk2.Push(value)
	}

	value = que.stk2.Pop().(int)
	return value
}

// Length returns the number of elements in the queue.
func (que *QueueUsingStack) Length() int {
	return que.stk1.Len() + que.stk2.Len()
}

// IsEmpty checks if the queue is empty.
func (que *QueueUsingStack) IsEmpty() bool {
	return que.Length() == 0
}

func main() {
	que := new(QueueUsingStack)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println(que.Length())
	fmt.Println(que.IsEmpty())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.IsEmpty())
}

/*
3
false
1
2
3
true
*/

type Stack struct {
	stk []interface{}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

// Pop removes and returns the element at the top of the stack.
func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[:n-1]
	return value
}

// Top returns the element at the top of the stack without removing it.
func (s *Stack) Top() interface{} {
	n := len(s.stk)
	return s.stk[n-1]
}

// Len returns the number of elements in the stack.
func (s Stack) Len() int {
	return len(s.stk)
}

// IsEmpty checks if the stack is empty.
func (s Stack) IsEmpty() bool {
	return len(s.stk) == 0
}

// Print prints the elements in the stack.
func (s Stack) Print() {
	fmt.Println(s.stk)
}
