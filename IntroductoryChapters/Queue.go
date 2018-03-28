package main

import "fmt"
import "github.com/golang-collections/go-datastructures/queue"
import "container/heap"

type HeapArr []int

func (h HeapArr) Len() int{
	return len(h)
}
func (h HeapArr) Less(i,j int) bool {
	return h[i] < h[j]
}
func (h HeapArr) Swap(i,j int){
	h[i], h[j] = h[j], h[i]
}
func (h *HeapArr) Push(val interface{}){
	*h = append(*h,val.(int))
}

func (h *HeapArr) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main31() {
	q := new(queue.Queue)
	q.Put(2)
	q.Put(3)
	q.Put(3, 4, 5, 6, 7)

	for q.Empty() == false {
		val, _ := q.Get(1)
		fmt.Println(val[0])
	}

	h:= &HeapArr{2,1,3}
	heap.Init(h)
	heap.Push(h,33)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}	
}
