package main

import "fmt"

const (
	EmptyNode byte = iota
	DeletedNode
	FillledNode
)

type HashTable struct {
	Key       []int
	Value     []int
	Flag      []byte
	tableSize int
}

func (ht *HashTable) Init(tSize int) {
	ht.tableSize = tSize
	ht.Key = make([]int, (tSize + 1))
	ht.Value = make([]int, (tSize + 1))
	ht.Flag = make([]byte, (tSize + 1))
}

func (ht *HashTable) ComputeHash(key int) int {
	return key % ht.tableSize
}

func (ht *HashTable) ResolverFun(index int) int {
	return index
}

func (ht *HashTable) Add(key int, args ...int) bool {
	value := key
	if len(args) > 0 {
		value = args[0]
	}

	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode || 
		ht.Flag[key] == DeletedNode {
			ht.Key[key] = key
			ht.Value[hashValue] = value
			ht.Flag[hashValue] = FillledNode
			return true
		} else if (ht.Flag[hashValue] == FillledNode && 
			ht.Key[hashValue] == key) {
            ht.Value[hashValue] = value
            return true
        }

		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return false
}

func (ht *HashTable) Find(key int) bool {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode {
			return false
		}
		if ht.Flag[hashValue] == FillledNode && 
		ht.Key[hashValue] == key {
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return false
}


func (ht *HashTable) Get(key int) int {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode {
			return 0
		}
		if ht.Flag[hashValue] == FillledNode && 
		ht.Key[hashValue] == key {
			return ht.Value[hashValue]
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.tableSize
	}
	return 0
}

func (ht *HashTable) Remove(key int) bool {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.tableSize; i++ {
		if ht.Flag[hashValue] == EmptyNode {
			return false
		}
		if ht.Flag[hashValue] == FillledNode && 
		ht.Key[hashValue] == key {
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
			fmt.Println("Node at index", i, "::", ht.Value[i])
		}
	}
	fmt.Println()
}

func main() {
	ht := new(HashTable)
	ht.Init(1000)
	ht.Add(1, 10)
	ht.Add(2, 10)
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
Node at index 1 :: 10
Node at index 2 :: 10
Node at index 3 :: 30

Find key 2 :  true
Value at key 2 :  10

After deleting node with key 2..
Find key 2 :  false

Values Stored in HashTable are::
Node at index 1 :: 10
Node at index 3 :: 30
*/
