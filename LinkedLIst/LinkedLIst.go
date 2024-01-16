package main

import (
	"fmt"
)

type List struct {
	head  *ListNode
	count int
}

type ListNode struct {
	value int
	next  *ListNode
}

func (list *List) Sum() int {
	temp := list.head
	sum := 0
	for temp != nil {
		sum += temp.value
		temp = temp.next
	}
	return sum
}

func (list *List) Size() int {
	return list.count
}

func (list *List) IsEmpty() bool {
	return list.count == 0
}

func (list *List) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("EmptyListException")
		return 0, false
	}
	return list.head.value, true
}

func (list *List) AddHead(value int) {
	list.head = &ListNode{value: value, next: list.head}
	list.count++
}

func (list *List) AddTail(value int) {
	newNode := &ListNode{value: value, next: nil}
	list.count++

	if list.head == nil {
		list.head = newNode
		return
	}

	curr := list.head
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
	newNode := &ListNode{value: value}
	if list.head == nil || list.head.value > value {
		newNode.next = list.head
		list.head = newNode
		return
	}
	curr := list.head

	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}

	newNode.next = curr.next
	curr.next = newNode
}

func (list *List) Find(data int) bool {
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
			list.count--
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
	var prev, next *ListNode
	current := list.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev
}

func (list *List) ReverseRecurse() {
	list.head = list.reverseRecurseUtil(list.head, nil)
}

func (list *List) reverseRecurseUtil(currentNode *ListNode, nextNode *ListNode) *ListNode {
	var ret *ListNode
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
			list.count--
		} else {
			curr = curr.next
		}
	}
}

func (list *List) CopyListReversed() *List {
	var tempNode, tempNode2 *ListNode
	curr := list.head
	for curr != nil {
		tempNode2 = &ListNode{curr.value, tempNode}
		curr = curr.next
		tempNode = tempNode2
	}
	ll2 := new(List)
	ll2.head = tempNode
	ll2.count = list.count
	return ll2
}

func (list *List) CopyList() *List {
	var headNode, tailNode, tempNode *ListNode
	curr := list.head

	if curr == nil {
		ll2 := new(List)
		ll2.head = nil
		return ll2
	}

	headNode = &ListNode{curr.value, nil}
	tailNode = headNode
	curr = curr.next

	for curr != nil {
		tempNode = &ListNode{curr.value, nil}
		tailNode.next = tempNode
		tailNode = tempNode
		curr = curr.next
	}
	ll2 := new(List)
	ll2.head = headNode
	ll2.count = list.count
	return ll2
}

func (list *List) CompareList(ll *List) bool {
	return list.compareListUtil(list.head, ll.head)
}

func (list *List) compareListUtil(head1 *ListNode, head2 *ListNode) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if (head1 == nil) || (head2 == nil) || (head1.value != head2.value) {
		return false
	} else {
		return list.compareListUtil(head1.next, head2.next)
	}
}

func (list *List) compareList2(head1 *ListNode, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.value != head2.value {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}

	if head1 == nil && head2 == nil {
		return true
	}
	return false
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
	size := list.FindLength()
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

	for fastPtr != nil && fastPtr.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next

		if slowPtr == fastPtr {
			fmt.Println("Loop found")
			return true
		}
	}
	fmt.Println("Loop not found")
	return false
}

func (list *List) ReverseListLoopDetect() bool {
	tempHead := list.head
	list.Reverse()
	if tempHead == list.head {
		list.Reverse()
		fmt.Println("Loop found")
		return true
	}
	list.Reverse()
	fmt.Println("Loop not found")
	return false
}

func (list *List) LoopTypeDetect() int {
	slowPtr := list.head
	fastPtr := list.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		if list.head == fastPtr.next || list.head == fastPtr.next.next {
			fmt.Println("Circular list loop found")
			return 2
		}
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("Loop found")
			return 1
		}
	}
	fmt.Println("Loop not found")
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

func (list *List) LoopPointDetect() *ListNode {
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

func (list *List) FindIntersection(list2 List) *ListNode {
	head := list.head
	head2 := list2.head
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

func main1() {
	ll := new(List) // or ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	fmt.Println("Size :", ll.Size())
	fmt.Println("IsEmpty :", ll.IsEmpty())
	fmt.Println(ll.Peek())
}

/*
3 2 1
Size : 3
IsEmpty : false
3 true
*/

func main2() {
	ll := List{}
	ll.SortedInsert(1)
	ll.SortedInsert(2)
	ll.SortedInsert(3)
	ll.SortedInsert(1)
	ll.SortedInsert(2)
	ll.SortedInsert(3)
	ll.Print()

}

/*
1 1 2 2 3 3
*/

func main3() {
	ll := new(List) // or ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(1)
	ll.AddHead(3)
	ll.Print()
	ll.DeleteNode(2)
	ll.Print()
	ll.DeleteNodes(1)
	ll.Print()
}

/*
3 1 2 1 2 1
3 1 1 2 1
*/

func main4() {
	ll := new(List) // or ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	ll.Reverse()
	ll.Print()
}

/*
3 2 1
1 2 3
*/

func main5() {
	ll := List{}
	ll.SortedInsert(1)
	ll.SortedInsert(2)
	ll.SortedInsert(3)
	ll.SortedInsert(1)
	ll.SortedInsert(2)
	ll.SortedInsert(3)
	ll.Print()
	ll.RemoveDuplicate()
	ll.Print()
}

/*
1 1 2 2 3 3
1 2 3
*/

func main6() {
	ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()

	ll2 := ll.CopyList()
	ll2.Print()
	fmt.Println(ll.CompareList(ll2))

	ll3 := ll.CopyListReversed()
	ll3.Print()
}

/*
3 2 1
3 2 1
true
1 2 3
*/

func main7() {
	ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	ll.MakeLoop()
	ll.LoopDetect()
	ll.ReverseListLoopDetect()
	ll.LoopTypeDetect()
	ll.RemoveLoop()
	ll.LoopDetect()
}

/*
3 2 1
loop found
loop found
circular list loop found
loop not found
*/

func main8() {
	ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll2 := List{}
	ll2.AddHead(3)
	ll2.head.next = ll.head
	ll.AddHead(4)
	ll2.AddHead(5)
	ll.Print()
	ll2.Print()
	nd := ll.FindIntersection(ll2)
	fmt.Print("Intersection :: ", nd.value)
}

/*
4 2 1
5 3 2 1
Intersection :: 2
*/

func main9() {
	ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	fmt.Println(ll.NthNodeFromBegining(2))
	fmt.Println(ll.NthNodeFromEnd(2))
	fmt.Println(ll.NthNodeFromEnd2(2))
}

func main10() {
	ll := List{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.AddTail(1)
	ll.AddTail(2)
	ll.AddTail(3)
	ll.Print()
	fmt.Println(ll.Size())
	fmt.Println(ll.IsEmpty())
	fmt.Println(ll.Peek())
	ll.DeleteNodes(3)
	ll.Print()
	fmt.Println(ll.Find(3))
	ll.RemoveHead()
	ll.Print()
	ll.FreeList()
	ll.Print()
}

/*
3 2 1 1 2 3
6
false
3 true
2 1 1 2
false
1 1 2
*/

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main10()
}
