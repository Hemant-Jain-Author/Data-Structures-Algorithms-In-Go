package main

import (
	"fmt"
)

type Node struct {
	c           rune
	freq        int
	left, right *Node
}

func NewNode(ch rune, fr int, l *Node, r *Node) *Node {
	p := new(Node)
	p.c = ch
	p.freq = fr
	p.left = l
	p.right = r
	return p
}

type HuffmanTree struct {
	root *Node
}

func NewHuffmanTree(arr []rune, freq []int) (huffTree *HuffmanTree) {
	huffTree = &HuffmanTree{}
	n := len(arr)

	cmp := func(a, b interface{}) bool { 
		return a.(*Node).freq > b.(*Node).freq 
	}
	hp := CreateHeap(cmp)

	for i := 0; i < n; i++ {
		node := NewNode(arr[i], freq[i], nil, nil)
		hp.Add(node)
	}

	for hp.Size() > 1 {
		lt := hp.Remove().(*Node)
		rt := hp.Remove().(*Node)
		nd := NewNode('+', lt.freq+rt.freq, lt, rt)
		hp.Add(nd)
	}
	huffTree.root = hp.Remove().(*Node)
	return
}

func (huffTree *HuffmanTree) printUtil(root *Node, s string) {
	if root.left == nil && root.right == nil {
		fmt.Println(string(root.c), " = "+s)
		return
	}
	huffTree.printUtil(root.left, s+"0")
	huffTree.printUtil(root.right, s+"1")
}

func (huffTree *HuffmanTree) Print() {
	fmt.Println("Char = Huffman code")
	huffTree.printUtil(huffTree.root, "")
}

func main() {
	ar := []rune{'A', 'B', 'C', 'D', 'E'}
	fr := []int{30, 25, 21, 14, 10}
	hf := NewHuffmanTree(ar, fr)
	hf.Print()
}

/*
Char = Huffman code
C  = 00
E  = 010
D  = 011
B  = 10
A  = 11
*/


type Heap struct {
	size  int
	arr   []interface{}
	comp func(x interface{}, y interface{}) bool
}

func CreateHeap(comp func(x interface{}, y interface{}) bool, args ...[]interface{}) *Heap {
	var arr []interface{}
	size := 0
	if len(args) > 0 {
		arrInput := args[0]
		arr = append(arr, arrInput...)
		size = len(arrInput)
	}

	h := &Heap{comp: comp, arr : arr, size : size}
	for i := (size / 2); i >= 0; i-- {
		h.percolateDown(i)
	}

	return h
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) percolateDown(parent int) {
	lChild := 2 * parent + 1
	rChild := lChild + 1
	child := -1
	if lChild < h.size {
		child = lChild
	}
	if rChild < h.size && h.comp(h.arr[lChild], h.arr[rChild]) {
		child = rChild
	}
	if child != -1 && h.comp(h.arr[parent], h.arr[child]) {
		h.swap(parent, child)
		h.percolateDown(child)
	}
}

func (h *Heap) percolateUp(child int) {
	parent := (child - 1) / 2
	if parent >= 0 && h.comp(h.arr[parent], h.arr[child]) {
		h.swap(child, parent)
		h.percolateUp(parent)
	}
}

func (h *Heap) Add(value interface{}) {
	h.arr = append(h.arr, value)
	h.size++
	h.percolateUp(h.size-1)
}

func (h *Heap) Remove() interface{} {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyException.")
		return 0
	}
	value := h.arr[0]
	h.arr[0] = h.arr[h.size - 1]
	h.size--
	h.percolateDown(0)
	h.arr = h.arr[0 : h.size]
	return value
}


func (h *Heap) Delete( value interface{}) bool {
    for i := 0; i < h.size; i++ {
        if (h.arr[i] == value) {
            h.arr[i] = h.arr[h.size - 1]
            h.size -= 1
            h.percolateUp(i)
            h.percolateDown(i)
            return true
        }
    }
    return false
}


func (h *Heap) IsEmpty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() interface{} {
	if h.IsEmpty() {
		fmt.Println("Heap empty exception.")
		return 0
	}
	return h.arr[0]
}

func (h *Heap) Print() {
	fmt.Println("Heap size :", h.size)
	fmt.Print("Heap Array :")
	for i := 0; i < h.size; i++ {
		fmt.Print(" ", h.arr[i])
	}
	fmt.Println()
}

func minComp (i, j interface{}) bool { // always i < j in use
	return i.(int) > j.(int) // swaps for min heap
}

func maxComp (i, j interface{}) bool { // always i < j in use
	return i.(int) < j.(int) // swap for max heap.
}
