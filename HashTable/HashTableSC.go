package main

import "fmt"

type Node struct {
	key int
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
	return key % h.tableSize
}

func (h *HashTableSC) Add(key int, args ...int) {
	value := key
	if len(args) > 0 {
		value = args[0]
	}

	index := h.ComputeHash(key)
	temp := new(Node)
	temp.key = key
	temp.value = value
	temp.next = h.listArray[index]
	h.listArray[index] = temp
}

func (h *HashTableSC) Remove(key int) bool {
	index := h.ComputeHash(key)
	var nextNode, head *Node
	head = h.listArray[index]
	if head != nil && head.key == key {
		h.listArray[index] = head.next
		return true
	}
	for head != nil {
		nextNode = head.next
		if nextNode != nil && nextNode.key == key {
			head.next = nextNode.next
			return true
		}
		head = nextNode
	}
	return false
}

func (h *HashTableSC) Find(key int) bool {
	index := h.ComputeHash(key)
	head := h.listArray[index]
	for head != nil {
		if head.key == key {
			return true
		}
		head = head.next
	}
	return false
}

func (h *HashTableSC) Get(key int) int {
	index := h.ComputeHash(key)
	head := h.listArray[index]
	for head != nil {
		if head.key == key {
			return head.value
		}
		head = head.next
	}
	return 0
}

func (h *HashTableSC) Print() {
	fmt.Print("\nValues Stored in HashTable are::")
	for i := 0; i < h.tableSize; i++ {
		head := h.listArray[i]
		if head != nil {
			fmt.Print("\nValues at index ", i, " :: ")
		}
		for head != nil {
			fmt.Print(head.value, " ")
			head = head.next
		}
	}
	fmt.Println("\n")	
}

func main() {
	ht := new(HashTableSC)
	ht.Init()

	ht.Add(1, 10)
	ht.Add(2, 20)
	ht.Add(3, 30)
	ht.Print()

	fmt.Println("Find key 2 : ", ht.Find(2))
	fmt.Println("Value at key 2 : ",ht.Get(2))

	ht.Remove(2)
	fmt.Println("\nAfter deleting node with key 2..")
	fmt.Println("Find key 2 : ", ht.Find(2))
}
/*
Values Stored in HashTable are::
Values at index 1 :: 10 
Values at index 2 :: 20 
Values at index 3 :: 30 

Find key 2 :  true
Value at key 2 :  20

After deleting node with key 2..
Find key 2 :  false

Values Stored in HashTable are::
Values at index 1 :: 10 
Values at index 3 :: 30 
*/