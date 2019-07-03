package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) Peek() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	return s.head.value, true
}

func (s *StackLinkedList) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}

	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

func (s *StackLinkedList) Print() {
	temp := s.head
	fmt.Print("Value stored in stck are :: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (s *StackLinkedList) insertAtBottom(value int) {
	if s.IsEmpty() {
		s.Push(value)
	} else {
		temp,_ := s.Pop()
		s.insertAtBottom(value)
		s.Push(temp)
	}
}


func main() {
	s := new(StackLinkedList)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	for s.IsEmpty() == false {
		val, _ := s.Pop()
		fmt.Print( val, " ")
	}
}