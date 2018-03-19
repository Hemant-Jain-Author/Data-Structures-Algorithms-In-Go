package main

import "fmt"
import "github.com/golang-collections/go-datastructures/queue"
import "container/heap"

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(val interface{}) {
	n := len(*pq)
	item := val.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return x
}

func main() {
	q := new(queue.Queue)
	q.Put(2)
	q.Put(3)
	q.Put(3, 4, 5, 6, 7)

	for q.Empty() == false {
		val, _ := q.Get(1)
		fmt.Println(val[0])
	}

	h := &PriorityQueue{2, 1, 3}
	heap.Init(h)
	heap.Push(h, 33)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}
