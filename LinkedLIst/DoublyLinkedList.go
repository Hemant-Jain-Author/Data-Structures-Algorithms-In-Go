package main

import (
	"fmt"
)

type DoublyLinkedList struct {
	head  *DLLNode
	tail  *DLLNode
	count int
}

type DLLNode struct {
	value int
	next  *DLLNode
	prev  *DLLNode
}

func (list *DoublyLinkedList) Size() int {
	return list.count
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *DoublyLinkedList) Peek() int {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
	}
	return list.head.value
}

func (list *DoublyLinkedList) AddHead(value int) {
	newNode := &DLLNode{value, nil, nil}
	if list.count == 0 {
		list.tail = newNode
		list.head = newNode
	} else {
		list.head.prev = newNode
		newNode.next = list.head
		list.head = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) AddTail(value int) {
	newNode := &DLLNode{value, nil, nil}
	if list.count == 0 {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.count++
}

func (list *DoublyLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0, false
	}

	value := list.head.value
	list.head = list.head.next

	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	list.count--
	return value, true
}

func (list *DoublyLinkedList) RemoveNode(key int) bool {
	curr := list.head
	if curr == nil { // empty list
		return false
	}
	if curr.value == key { // head is the Node with value key.
		curr = curr.next
		list.count--
		if curr != nil {
			list.head = curr
			list.head.prev = nil
		} else {
			list.tail = nil // only one element in list.
		}
		return true
	}
	for curr.next != nil {
		if curr.next.value == key {
			curr.next = curr.next.next
			if curr.next == nil { // last element case.
				list.tail = curr
			} else {
				curr.next.prev = curr
			}
			list.count--
			return true
		}
		curr = curr.next
	}
	return false
}

func (list *DoublyLinkedList) IsPresent(key int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *DoublyLinkedList) FreeList() {
	list.tail = nil
	list.head = nil
	list.count = 0
}

func (list *DoublyLinkedList) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (list *DoublyLinkedList) ReverseList() {
	curr := list.head
	var tempNode *DLLNode
	for curr != nil {
		tempNode = curr.next
		curr.next = curr.prev
		curr.prev = tempNode
		if curr.prev == nil {
			list.tail = list.head
			list.head = curr
			return
		}
		curr = curr.prev
	}
	return
}

func (list *DoublyLinkedList) CopyListReversed( ) *DoublyLinkedList {
	dll := new(DoublyLinkedList)

	curr := list.head
	for curr != nil {
		dll.AddHead(curr.value)
		curr = curr.next
	}
	dll.count = list.count
	return dll
}

func (list *DoublyLinkedList) CopyList( ) *DoublyLinkedList {
	dll := new(DoublyLinkedList)
	curr := list.head
	for curr != nil {
		dll.AddTail(curr.value)
		curr = curr.next
	}
	dll.count = list.count
	return dll
}

func (list *DoublyLinkedList) SortedInsert(value int) {
	temp := &DLLNode{value, nil, nil}
	curr := list.head
	if curr == nil { // first element
		list.head = temp
		list.tail = temp
		return
	}

	if list.head.value <= value { // at the begining
		temp.next = list.head
		list.head.prev = temp
		list.head = temp
		return 
	}

	for curr.next != nil && curr.next.value > value { // treversal
		curr = curr.next
	}

	if curr.next == nil { // at the end
		list.tail = temp
		temp.prev = curr
		curr.next = temp
	} else { // all other
		temp.next = curr.next
		temp.prev = curr
		curr.next = temp
		temp.next.prev = temp
	}
}

func (list *DoublyLinkedList) RemoveDuplicate() {
	curr := list.head
	var deleteMe *DLLNode
	for curr != nil {
		if (curr.next != nil) && curr.value == curr.next.value {
			deleteMe = curr.next
			curr.next = deleteMe.next
			curr.next.prev = curr
			if deleteMe == list.tail {
				list.tail = curr
			}
		} else {
			curr = curr.next
		}
	}
}

func main() {
	ll := new(DoublyLinkedList)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.AddHead(4)
	ll.AddHead(5)
	ll.AddHead(6)
	ll.Print()
	ll.RemoveHead()
	ll.Print()
	ll.FreeList()
	ll.Print()

	ll.AddHead(11)
	ll.AddHead(21)
	ll.AddHead(31)
	ll.AddHead(41)
	ll.AddHead(51)
	ll.AddHead(61)
	ll.Print()
	ll2 := ll.CopyList()
	ll2.Print()
	ll2 = ll.CopyListReversed()
	ll2.Print()

}
