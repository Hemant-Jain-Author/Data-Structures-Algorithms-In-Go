package main

import (
	"fmt"
)

type DoublyCircularLinkedList struct {
	head  *DCListNode
	tail  *DCListNode
	count int
}

type DCListNode struct {
	value int
	next  *DCListNode
	prev  *DCListNode
}

func (list *DoublyCircularLinkedList) Size() int {
	return list.count
}

func (list *DoublyCircularLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *DoublyCircularLinkedList) peekHead() int {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
	}
	return list.head.value
}

func (list *DoublyCircularLinkedList) IsPresent(key int) bool {
	temp := list.head
	if list.head == nil {
		return false
	}

	for true {
		if temp.value == key {
			return true
		}
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	return false
}

func (list *DoublyCircularLinkedList) FreeList() {
	list.head = nil
	list.tail = nil
	list.count = 0
}

func (list *DoublyCircularLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	fmt.Println("List size is ::", list.count)
	fmt.Print("List content :: ")
	temp := list.head
	for true {
		fmt.Print(temp.value, " ")
		temp = temp.next
		if temp == list.head {
			break
		}
	}
	fmt.Println()
}

func (list *DoublyCircularLinkedList) AddHead(value int) {
	newNode := new(DCListNode)
	newNode.value = value
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = list.head
		newNode.prev = list.head.prev
		list.head.prev = newNode
		newNode.prev.next = newNode
		list.head = newNode
	}
	list.count++
}

func (list *DoublyCircularLinkedList) AddTail(value int) {
	newNode := new(DCListNode)
	newNode.value = value
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = list.tail.next
		newNode.prev = list.tail
		list.tail.next = newNode
		newNode.next.prev = newNode
		list.tail = newNode
	}
	list.count++
}

func (list *DoublyCircularLinkedList) RemoveHead() (int, bool) {
	if list.count == 0 {
		fmt.Println("EmptyListException")
		return 0, false
	}

	value := list.head.value
	list.count--

	if list.count == 0 {
		list.head = nil
		list.tail = nil
		return value, true
	}

	next := list.head.next
	next.prev = list.tail
	list.tail.next = next
	list.head = next
	return value, true
}

func (list *DoublyCircularLinkedList) RemoveTail() (int, bool) {
	if list.count == 0 {
		fmt.Println("EmptyListException")
		return 0, false
	}

	value := list.tail.value
	list.count--
	if list.count == 0 {
		list.head = nil
		list.tail = nil
		return value, true
	}

	prev := list.tail.prev
	prev.next = list.head
	list.head.prev = prev
	list.tail = prev
	return value, true
}

func main() {
	ll := new(DoublyCircularLinkedList)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
}