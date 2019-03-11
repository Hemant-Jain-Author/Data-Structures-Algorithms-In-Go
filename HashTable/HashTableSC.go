package main

import "fmt"

func main() {
	ht := new(HashTableSC)
	ht.Init(101)
	ht.Add(1)
	ht.Add(2)
	ht.Add(3)
	ht.Print()
	fmt.Println("1 found : ", ht.Find(1))
	fmt.Println("4 found : ", ht.Find(4))
	fmt.Println("1 remove : ", ht.Remove(1))
	fmt.Println("4 remove : ", ht.Remove(4))
	ht.Print()
}

type Node struct {
	value int
	next  *Node
}

type HashTableSC struct {
	listArray [](*Node)
	tableSize int
}

func (h *HashTableSC) Init(tSize int) {
	h.tableSize = tSize
	h.listArray = make([](*Node), h.tableSize)

	for i := 0; i < h.tableSize; i++ {
		h.listArray[i] = nil
	}
}

func (h *HashTableSC) ComputeHash(key int) int {
	hashValue := key
	return hashValue % h.tableSize
}

func (h *HashTableSC) Add(value int) {
	index := h.ComputeHash(value)
	temp := new(Node)
	temp.value = value
	temp.next = h.listArray[index]
	h.listArray[index] = temp
}

func (h *HashTableSC) Remove(value int) bool {
	index := h.ComputeHash(value)
	var nextNode, head *Node
	head = h.listArray[index]
	if head != nil && head.value == value {
		h.listArray[index] = head.next
		return true
	}
	for head != nil {
		nextNode = head.next
		if nextNode != nil && nextNode.value == value {
			head.next = nextNode.next
			return true
		}
		head = nextNode
	}
	return false
}

func (h *HashTableSC) Print() {
	for i := 0; i < h.tableSize; i++ {
		head := h.listArray[i]
		if head != nil {
			fmt.Print("\nValues at index :: ", i, " are :: ")
		}
		for head != nil {
			fmt.Print(head.value, " ")
			head = head.next
		}
	}
	fmt.Println()
}

func (h *HashTableSC) Find(value int) bool {
	index := h.ComputeHash(value)
	head := h.listArray[index]
	for head != nil {
		if head.value == value {
			return true
		}
		head = head.next
	}
	return false
}
