package main

import "fmt"

type CircularLinkedList struct {
	tail  *CircularLinkedListNode
	count int
}

type CircularLinkedListNode struct {
	value int
	next  *CircularLinkedListNode
}

func (list *CircularLinkedList) Size() int {
	return list.count
}

func (list *CircularLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *CircularLinkedList) Peek() int {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0
	}
	return list.tail.next.value
}

func (list *CircularLinkedList) AddHead(value int) {
	temp := &CircularLinkedListNode{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
	}
	list.count++
}

func (list *CircularLinkedList) AddTail(value int) {
	temp := &CircularLinkedListNode{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
		list.tail = temp
	}
	list.count++
}

func (list *CircularLinkedList) RemoveHead() int {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
	}
	value := list.tail.next.value

	if list.tail == list.tail.next {
		list.tail = nil
	} else {
		list.tail.next = list.tail.next.next
	}
	list.count--
	return value
}

func (list *CircularLinkedList) IsPresent(data int) bool {
	temp := list.tail
	for i := 0; i < list.count; i++ {
		if temp.value == data {
			return true
		}

		temp = temp.next
	}
	return false
}

func (list *CircularLinkedList) Print() {
	if list.IsEmpty() {
		return
	}
	temp := list.tail.next
	for temp != list.tail {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func (list *CircularLinkedList) FreeList() {
	list.tail = nil
	list.count = 0
}

func (list *CircularLinkedList) RemoveNode(key int) bool {
	if list.IsEmpty() {
		return false
	}

	prev := list.tail
	curr := list.tail.next
	head := list.tail.next

	if curr.value == key { // head and single node case.
		if curr == curr.next { // single node case
			list.tail = nil
		} else { // head case
			list.tail.next = list.tail.next.next
		}
		list.count--
		return true
	}

	prev = curr
	curr = curr.next

	for curr != head {
		if curr.value == key {
			if curr == list.tail {
				list.tail = prev
			}
			prev.next = curr.next
			list.count--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func (list *CircularLinkedList) CopyListReversed() *CircularLinkedList {
	cl := new(CircularLinkedList)
	curr := list.tail.next
	head := curr

	if curr != nil {
		cl.AddHead(curr.value)
		curr = curr.next
	}
	for curr != head {
		cl.AddHead(curr.value)
		curr = curr.next
	}
	cl.count = list.count
	return cl
}


func (list *CircularLinkedList) CopyList() *CircularLinkedList {
	cl := new(CircularLinkedList)
	curr := list.tail.next
	head := curr

	if curr != nil {
		cl.AddTail(curr.value)
		curr = curr.next
	}
	for curr != head {
		cl.AddTail(curr.value)
		curr = curr.next
	}
	cl.count = list.count
	return cl
}

func main() {
	ll := new(CircularLinkedList)

	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()

	ll.RemoveNode(3)
	ll.Print()
	ll.RemoveHead()
	ll.Print()

	ll2 := ll.CopyList()
	ll2.Print()
	ll3 := ll.CopyListReversed()
	ll3.Print()
}