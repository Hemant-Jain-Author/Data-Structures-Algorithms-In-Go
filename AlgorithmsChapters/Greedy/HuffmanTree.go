package main

import (
	"container/heap"
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
		return a.(*Node).freq < b.(*Node).freq 
	}
	hp := NewHeap(cmp)

	for i := 0; i < n; i++ {
		node := NewNode(arr[i], freq[i], nil, nil)
		heap.Push(hp, node)
	}

	for hp.Len() > 1 {
		lt := heap.Pop(hp).(*Node)
		rt := heap.Pop(hp).(*Node)
		nd := NewNode('+', lt.freq+rt.freq, lt, rt)
		heap.Push(hp, nd)
	}
	huffTree.root = heap.Pop(hp).(*Node)
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

type Heap struct {
	heap []interface{}
	comp func(x interface{}, y interface{}) bool
}

func NewHeap(comp func(x interface{}, y interface{}) bool) *Heap {
	hp := new(Heap)
	hp.comp = comp
	return hp
}

func (hp Heap) Len() int {
	return len(hp.heap)
}

func (hp Heap) Less(i, j int) bool {
	return hp.comp(hp.heap[i], hp.heap[j])
}

func (hp Heap) Swap(i, j int) {
	hp.heap[i], hp.heap[j] = hp.heap[j], hp.heap[i]
}

func (hp *Heap) Push(x interface{}) {
	hp.heap = append(hp.heap, x)
}

func (hp *Heap) Pop() interface{} {
	n := len(hp.heap)
	value := hp.heap[n-1]
	hp.heap = hp.heap[0 : n-1]
	return value
}

func (hp Heap) Print() {
	fmt.Println(hp.heap)
}

func (hp Heap) Empty() bool {
	return len(hp.heap) == 0
}

func (hp Heap) Peek() interface{} {
	return hp.heap[0]
}