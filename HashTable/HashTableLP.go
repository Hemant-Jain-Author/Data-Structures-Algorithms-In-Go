package main

import "fmt"

const (
	EmptyNode byte = iota
	LazyDeleted
	FillledNode
)

type HashTable struct {
	Arr       []int
	Flag      []byte
	tableSize int
}

func (ht *HashTable) Init(tSize int) {
	ht.tableSize = tSize
	ht.Arr = make([]int, (tSize + 1))
	ht.Flag = make([]byte, (tSize + 1))
}

func (ht *HashTable) ComputeHash(key int) int {
	return key % ht.tableSize
}

func (ht *HashTable) ResolverFun(index int) int {
	return index
}

func (ht *HashTable) Add(value int) bool {
	hashValue := ht.ComputeHash(value)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode || ht.Flag[hashValue] == LazyDeleted {
			ht.Arr[hashValue] = value
			ht.Flag[hashValue] = FillledNode
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return false
}

func (ht *HashTable) Find(value int) bool {
	hashValue := ht.ComputeHash(value)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode {
			return false
		}
		if ht.Flag[hashValue] == FillledNode && ht.Arr[hashValue] == value {
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return false
}

func (ht *HashTable) Remove(value int) bool {
	hashValue := ht.ComputeHash(value)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode {
			return false
		}
		if ht.Flag[hashValue] == FillledNode && ht.Arr[hashValue] == value {
			ht.Flag[hashValue] = LazyDeleted
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return false
}

func (ht *HashTable) Print() {
	fmt.Println("\nValues Stored in HashTable are::")
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[i] == FillledNode {
			fmt.Println("Node at index [", i, " ] :: ", ht.Arr[i])
		}
	}
}

func main() {
	ht := new(HashTable)
	ht.Init(1000)
	ht.Add(89)
	ht.Add(18)
	ht.Add(49)
	ht.Add(58)
	ht.Add(69)
	ht.Add(89)
	ht.Add(18)
	ht.Add(49)
	ht.Add(58)
	ht.Add(69)
	ht.Print()
	fmt.Println("89 found : ", ht.Find(89))
	ht.Remove(89)
	ht.Remove(18)
	ht.Remove(49)
	ht.Remove(58)
	ht.Remove(69)
	ht.Print()
}

/*
Values Stored in HashTable are::
Node at index [18 ] :: 18
Node at index [19 ] :: 18
Node at index [49 ] :: 49
Node at index [50 ] :: 49
Node at index [58 ] :: 58
Node at index [59 ] :: 58
Node at index [69 ] :: 69
Node at index [70 ] :: 69
Node at index [89 ] :: 89
Node at index [90 ] :: 89

Values Stored in HashTable are::
Node at index [19 ] :: 18
Node at index [50 ] :: 49
Node at index [59 ] :: 58
Node at index [70 ] :: 69
Node at index [90 ] :: 89
*/
