package main

import (
	"fmt"
)

type List struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Insert(v int) *List {
	if list == nil {
		fmt.Println("list is not created.")
		return nil
	}
	temp := new(Node)
	temp.value = v
	temp.next = list.head
	list.head = temp
	return list
}

// Sum returns the sum of the list elements.
func (list *List) Sum() int {
	if list == nil {
		fmt.Println("list is not created.")
		return 0
	}
	temp := list.head
	sum := 0
	for temp != nil {
		sum += temp.value
		temp = temp.next
	}
	return sum
}

func (list *List) Print() {
	if list == nil {
		fmt.Println("list is not created.")
		return
	}
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}

func main() {
	// lst := new(List)
	lst := List{}
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	fmt.Println(lst.Sum())
	lst.Print()
}