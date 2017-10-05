package main

import (
	"fmt"
)

type List struct {
	head  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

// Sum returns the sum of the list elements.
func (list *List) Sum() int {
	temp := list.head
	sum := 0
	for temp != nil {
		sum += temp.value
		temp = temp.next
	}
	return sum
}

func main() {
	lst := new(List)
	//lst := List{}
	lst.AddHead(1)
	lst.AddHead(2)
	lst.AddHead(3)
	lst.Print()
	lst2 := lst.CopyList()
	fmt.Println(lst.compareList(lst2))
	lst3 := lst.CopyListReversed()
	lst3.Print()
	fmt.Println(lst.IsPresent(7))
	fmt.Println(lst.IsPresent(4))
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return (list.count == 0)
}

func (list *List) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0, false
	}
	return list.head.value, true
}

func (list *List) AddHead(value int) {
	list.head = &Node{value, list.head}
	list.count++
}

func (list *List) addTail(value int) {
	curr := list.head
	newNode := &Node{value, nil}

	if curr == nil {
		list.head = newNode
		return
	}

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (list *List) Print() {
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}

func (list *List) SortedInsert(value int) {
	newNode := &Node{value, nil}
	curr := list.head

	if curr == nil || curr.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}

	newNode.next = curr.next
	curr.next = newNode
}

func (list *List) IsPresent(data int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.next
	list.count--
	return value, true
}

func (list *List) DeleteNode(delValue int) bool {
	temp := list.head
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return false
	}

	if delValue == list.head.value {
		list.head = list.head.next
		list.count--
		return true
	}

	for temp.next != nil {
		if temp.next.value == delValue {
			temp.next = temp.next.next
			list.count--
			return true
		}
		temp = temp.next
	}
	return false
}

func (list *List) DeleteNodes(delValue int) {
	currNode := list.head
	for currNode != nil && currNode.value == delValue {
		list.head = currNode.next
		currNode = list.head
	}

	for currNode != nil {
		nextNode := currNode.next
		if nextNode != nil && nextNode.value == delValue {
			currNode.next = nextNode.next
		} else {
			currNode = nextNode
		}
	}
}

func (list *List) FreeList() {
	list.head = nil
	list.count = 0
}

func (list *List) Reverse() {
	curr := list.head
	var prev, next *Node
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}

func (list *List) ReverseRecurse() {
	list.head = list.reverseRecurseUtil(list.head, nil)
}

func (list *List) reverseRecurseUtil(currentNode *Node, nextNode *Node) *Node {
	var ret *Node
	if currentNode == nil {
		return nil
	}
	if currentNode.next == nil {
		currentNode.next = nextNode
		return currentNode
	}

	ret = list.reverseRecurseUtil(currentNode.next, currentNode)
	currentNode.next = nextNode
	return ret
}

func (list *List) RemoveDuplicate() {
	curr := list.head

	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}
func (list *List) CopyListReversed() *List {
	var tempNode, tempNode2 *Node
	curr := list.head
	for curr != nil {
		tempNode2 = &Node{curr.value, tempNode}
		curr = curr.next
		tempNode = tempNode2
	}
	ll2 := new(List)
	ll2.head = tempNode
	return ll2
}

func (list *List) CopyList() *List {
	var headNode, tailNode, tempNode *Node
	curr := list.head

	if curr == nil {
		ll2 := new(List)
		ll2.head = nil
		return ll2
	}

	headNode = &Node{curr.value, nil}
	tailNode = headNode
	curr = curr.next

	for curr != nil {
		tempNode = &Node{curr.value, nil}
		tailNode.next = tempNode
		tailNode = tempNode
		curr = curr.next
	}
	ll2 := new(List)
	ll2.head = headNode
	return ll2
}

func (list *List) CompareList(ll *List) bool {
	return list.compareListUtil(list.head, ll.head)
}

func (list *List) compareListUtil(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if (head1 == nil) || (head2 == nil) || (head1.value != head2.value) {
		return false
	} else {
		return list.compareListUtil(head1.next, head2.next)
	}
}

func (list *List) FindLength() int {
	curr := list.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (list *List) NthNodeFromBegining(index int) (int, bool) {
	if index > list.Size() || index < 1 {
		fmt.Println("TooFewNodes")
		return 0, false
	}
	count := 0
	curr := list.head
	for curr != nil && count < index-1 {
		count++
		curr = curr.next
	}
	return curr.value, true
}

func (list *List) NthNodeFromEnd(index int) (int, bool) {
	size := list.findLength()
	if size != 0 && size < index {
		fmt.Println("TooFewNodes")
		return 0, false
	}
	startIndex := size - index + 1
	return list.NthNodeFromBegining(startIndex)
}

func (list *List) NthNodeFromEnd2(index int) (int, bool) {
	count := 1
	forward := list.head
	curr := list.head
	for forward != nil && count <= index {
		count++
		forward = forward.next
	}

	if forward == nil {
		fmt.Println("TooFewNodes")
		return 0, false
	}

	for forward != nil {
		forward = forward.next
		curr = curr.next
	}
	return curr.value, true
}

func (list *List) MakeLoop() {
	temp := list.head
	for temp != nil {
		if temp.next == nil {
			temp.next = list.head
			return
		}
		temp = temp.next
	}
}

func (list *List) LoopDetect() bool {
	slowPtr := list.head
	fastPtr := list.head

	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return true
		}
	}
	fmt.Println("loop not found")
	return false
}

func (list *List) ReverseListLoopDetect() bool {
	tempHead := list.head
	list.Reverse()
	if tempHead == list.head {
		list.Reverse()
		fmt.Println("loop found")
		return true
	}
	list.Reverse()
	fmt.Println("loop not found")
	return false
}

func (list *List) LoopTypeDetect() int {
	slowPtr := list.head
	fastPtr := list.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		if list.head == fastPtr.next || list.head == fastPtr.next.next {
			fmt.Println("circular list loop found")
			return 2
		}
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return 1
		}
	}
	fmt.Println("loop not found")
	return 0
}

func (list *List) RemoveLoop() {
	loopPoint := list.LoopPointDetect()
	if loopPoint == nil {
		return
	}

	firstPtr := list.head
	if loopPoint == list.head {
		for firstPtr.next != list.head {
			firstPtr = firstPtr.next
		}
		firstPtr.next = nil
		return
	}

	secondPtr := loopPoint
	for firstPtr.next != secondPtr.next {
		firstPtr = firstPtr.next
		secondPtr = secondPtr.next
	}
	secondPtr.next = nil
}

func (list *List) LoopPointDetect() *Node {
	slowPtr := list.head
	fastPtr := list.head

	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return slowPtr
		}
	}
	return nil
}

func (list *List) FindIntersection(head *Node, head2 *Node) *Node {
	l1 := 0
	l2 := 0
	tempHead := head
	tempHead2 := head2
	for tempHead != nil {
		l1++
		tempHead = tempHead.next
	}
	for tempHead2 != nil {
		l2++
		tempHead2 = tempHead2.next
	}
	var diff int
	if l1 < 12 {
		temp := head
		head = head2
		head2 = temp
		diff = l2 - l1
	} else {
		diff = l1 - l2
	}
	for ; diff > 0; diff-- {
		head = head.next
	}
	for head != head2 {
		head = head.next
		head2 = head2.next
	}
	return head
}
