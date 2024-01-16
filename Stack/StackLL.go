package main

import "fmt"

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

func (s *StackLinkedList) Peek() int {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return 0
    }
    return s.head.value
}

func (s *StackLinkedList) Push(value int) {
    s.head = &Node{value, s.head}
    s.size++
}

func (s *StackLinkedList) Pop() int {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return 0
    }

    value := s.head.value
    s.head = s.head.next
    s.size--
    return value
}

func (s *StackLinkedList) Print() {
    temp := s.head
    fmt.Print("[")
    for temp != nil {
        fmt.Print(temp.value, " ")
        temp = temp.next
    }
    fmt.Println("]")
}

func (s *StackLinkedList) insertAtBottom(value int) {
    if s.IsEmpty() {
        s.Push(value)
    } else {
        temp := s.Pop()
        s.insertAtBottom(value)
        s.Push(temp)
    }
}

// Testing code
func main() {
    stk := new(StackLinkedList)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Print()
	fmt.Print(stk.Pop(), " ")
	fmt.Print(stk.Pop(), " ")
}

/*
[3 2 1 ]
3 2 
*/

