package main

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
}

type HashTableSC struct {
	listArray [](*Node)
	tableSize int
}

func NewHashTableSC() *HashTableSC {
	h := &HashTableSC{
		tableSize: 101,
		listArray: make([](*Node), 101),
	}
	return h
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
	temp := &Node{
		key:   key,
		value: value,
		next:  h.listArray[index],
	}
	h.listArray[index] = temp
}

func (h *HashTableSC) Remove(key int) bool {
	index := h.ComputeHash(key)
	head := h.listArray[index]

	if head != nil && head.key == key {
		h.listArray[index] = head.next
		return true
	}

	var prev *Node
	for head != nil {
		if head.key == key {
			prev.next = head.next
			return true
		}
		prev = head
		head = head.next
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
	fmt.Print("Hash Table contains :: ")
	for i := 0; i < h.tableSize; i++ {
		head := h.listArray[i]
		for head != nil {
			fmt.Print("(", i, "=>", head.value, ") ")
			head = head.next
		}
	}
	fmt.Println()
}

// Testing code.
func main() {
	ht := NewHashTableSC()

	ht.Add(1, 10)
	ht.Add(2, 20)
	ht.Add(3, 30)
	ht.Print()

	fmt.Println("Find key 2:", ht.Find(2))
	fmt.Println("Value at key 2:", ht.Get(2))

	ht.Remove(2)
	fmt.Println("Find key 2:", ht.Find(2))
}

/*
Hash Table contains :: (1=>10) (2=>20) (3=>30)
Find key 2: true
Value at key 2: 20
Find key 2: false
*/
