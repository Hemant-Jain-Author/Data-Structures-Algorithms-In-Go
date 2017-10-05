package main

import "fmt"

func main() {
	ht := new(HashTableSC)
	ht.Init()
	ht.Add(89)
	ht.Print()
	fmt.Println("89 found : ", ht.Find(89))
}

type Node struct {
	value int
	next  *Node
}

type HashTableSC struct {
	listArray [](*Node)
	tableSize int
}

func (h *HashTableSC) Init() {
	h.tableSize = 101
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
			fmt.Print("\nValues at index :: ", i, " Are :: ")
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
