package main

import "fmt"

const (
	EmptyNode byte = iota
	DeletedNode
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
		if ht.Flag[hashValue] == EmptyNode || ht.Flag[hashValue] == DeletedNode {
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
			ht.Flag[hashValue] = DeletedNode
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

/*
Values Stored in HashTable are::
Node at index [ 1  ] ::  1
Node at index [ 2  ] ::  2
Node at index [ 3  ] ::  3
1 found :  true
4 found :  false
1 remove :  true
4 remove :  false

Values Stored in HashTable are::
Node at index [ 2  ] ::  2
Node at index [ 3  ] ::  3
*/
