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
	hp := &Heap{}

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
	huffTree.root = (*hp)[0]
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

type Heap []*Node

func (hp Heap) Len() int {
	return len(hp)
}

func (hp Heap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}

func (hp Heap) Less(p, c int) bool {
	return hp[p].freq < hp[c].freq
}

func (hp *Heap) Empty() bool {
	return (hp.Len() == 0)
}

func (hp *Heap) Peek() interface{} {
	return (*hp)[0]
}

func (hp *Heap) Print() {
	n := hp.Len()
	for i := 0; i < n; i++ {
		fmt.Print((*hp)[i].freq, " ")
	}
	fmt.Println()
}

func (hp *Heap) Push(x interface{}) {
	*hp = append(*hp, x.(*Node))
}

func (hp *Heap) Pop() interface{} {
	old := *hp
	n := len(old)
	value := old[n-1]
	*hp = old[0 : hp.Len()-1]
	return value
}
