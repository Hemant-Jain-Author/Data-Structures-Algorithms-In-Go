package main

import "fmt"

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

func (list *DoublyCircularLinkedList) PeekHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0, false
	}
	return list.head.value, true
}

func (list *DoublyCircularLinkedList) IsPresent(key int) bool {
	if list.head == nil {
		return false
	}

	temp := list.head
	for temp != list.tail {
		if temp.value == key {
			return true
		}
		temp = temp.next
	}
	// Check tail node separately since the loop condition does not cover it
	if temp.value == key {
		return true
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

	temp := list.head
	for temp != list.tail {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func (list *DoublyCircularLinkedList) AddHead(value int) {
	newNode := &DCListNode{value: value}
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
	newNode := &DCListNode{value: value}
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

func main1() {
	ll := &DoublyCircularLinkedList{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	fmt.Println("Size:", ll.Size())
	fmt.Println("IsEmpty:", ll.IsEmpty())
}

/*
Output:
3 2 1
Size: 3
IsEmpty: false
*/

func main2() {
	ll := &DoublyCircularLinkedList{}
	ll.AddTail(1)
	ll.AddTail(2)
	ll.AddTail(3)
	ll.Print()
}

/*
Output:
1 2 3
*/

func main3() {
	ll := &DoublyCircularLinkedList{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	fmt.Println("IsPresent:", ll.IsPresent(3))
}

/*
Output:
3 2 1
IsPresent: true
*/

func main4() {
	ll := &DoublyCircularLinkedList{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	ll.RemoveHead()
	ll.Print()
}

/*
Output:
3 2 1
2 1
*/

func main5() {
	ll := &DoublyCircularLinkedList{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	ll.RemoveTail()
	ll.Print()
}

/*
Output:
3 2 1
3 2
*/

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
}
