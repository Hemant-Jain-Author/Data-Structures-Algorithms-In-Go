package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	return s.head.value, true
}

func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}

	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

func (s *Stack) Print() {
	temp := s.head
	fmt.Print("Value stored in stck are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (s *Stack) insertAtBottom(value int) {
	if s.IsEmpty() {
		s.Push(value)
	} else {
		temp := s.Pop()
		s.insertAtBottom(value)
		s.Push(temp)
	}
}

func main() {
	s := new(Stack)
	for i := 1; i <= 100; i++ {
		s.Push(i)
	}
	fmt.Println("popped values are :: ")
	for i := 1; i <= 50; i++ {
		fmt.Print(s.Pop(), " ")
	}

	s.Print()
}
