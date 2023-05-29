package main

import "fmt"

const (
	EmptyNode byte = iota
	DeletedNode
	FilledNode
)

type HashTable struct {
	Keys      []int
	Values    []int
	Flags     []byte
	TableSize int
}

func NewHashTable(tableSize int) *HashTable {
	ht := &HashTable{
		Keys:      make([]int, tableSize+1),
		Values:    make([]int, tableSize+1),
		Flags:     make([]byte, tableSize+1),
		TableSize: tableSize,
	}
	return ht
}

func (ht *HashTable) ComputeHash(key int) int {
	return key % ht.TableSize
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
	for i := 0; i < ht.TableSize; i++ {
		if ht.Flags[hashValue] == EmptyNode || ht.Flags[key] == DeletedNode {
			ht.Keys[key] = key
			ht.Values[hashValue] = value
			ht.Flags[hashValue] = FilledNode
			return true
		} else if ht.Flags[hashValue] == FilledNode && ht.Keys[hashValue] == key {
			ht.Values[hashValue] = value
			return true
		}

		hashValue += ht.ResolverFun(i)
		hashValue %= ht.TableSize
	}
	return false
}

func (ht *HashTable) Find(key int) bool {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.TableSize; i++ {
		if ht.Flags[hashValue] == EmptyNode {
			return false
		}
		if ht.Flags[hashValue] == FilledNode && ht.Keys[hashValue] == key {
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.TableSize
	}
	return false
}

func (ht *HashTable) Get(key int) int {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.TableSize; i++ {
		if ht.Flags[hashValue] == EmptyNode {
			return 0
		}
		if ht.Flags[hashValue] == FilledNode && ht.Keys[hashValue] == key {
			return ht.Values[hashValue]
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.TableSize
	}
	return 0
}

func (ht *HashTable) Remove(key int) bool {
	hashValue := ht.ComputeHash(key)
	for i := 0; i < ht.TableSize; i++ {
		if ht.Flags[hashValue] == EmptyNode {
			return false
		}
		if ht.Flags[hashValue] == FilledNode && ht.Keys[hashValue] == key {
			ht.Flags[hashValue] = DeletedNode
			return true
		}
		hashValue += ht.ResolverFun(i)
		hashValue %= ht.TableSize
	}
	return false
}

func (ht *HashTable) Print() {
	fmt.Print("Hash Table contains :: ")
	for i := 0; i < ht.TableSize; i++ {
		if ht.Flags[i] == FilledNode {
			fmt.Print("(", i, "=>", ht.Values[i], ") ")
		}
	}
	fmt.Println()
}

func main() {
	ht := NewHashTable(1000)
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
