package main

import (
	"fmt"
	"strconv"
)

type item struct {
	value    interface{}
	priority int
}

type PQueue struct {
	items []*item
	Count int
	isMin bool
}

func newItem(value interface{}, priority int) *item {
	return &item{
		value:    value,
		priority: priority,
	}
}

func CreatePQueue(isMin bool) *PQueue {
	items := make([]*item, 1)
	items[0] = nil // Heap queue first element should always be nil

	return &PQueue{
		items: items,
		Count: 0,
		isMin: isMin,
	}
}

func (pq *PQueue) comp(i, j int) bool { // always i < j in use
	if pq.isMin == true {
		return pq.items[i].priority > pq.items[j].priority // swaps for min heap
	}
	return pq.items[i].priority < pq.items[j].priority // swap for max heap.
}

func (pq *PQueue) proclateDown(position int) {
	lChild := 2 * position
	rChild := lChild + 1
	small := -1

	if lChild <= pq.Count {
		small = lChild
	}

	if rChild <= pq.Count && pq.comp(lChild, rChild) {
		small = rChild
	}

	if small != -1 && pq.comp(position, small) {
		pq.items[position], pq.items[small] = pq.items[small], pq.items[position]
		pq.proclateDown(small)
	}
}

func (pq *PQueue) proclateUp(position int) {
	parent := position / 2
	if parent == 0 {
		return
	}

	if pq.comp(parent, position) {
		pq.items[position], pq.items[parent] = pq.items[parent], pq.items[position]
		pq.proclateUp(parent)
	}
}

func (pq *PQueue) Add(value interface{}, priority int) {
	item := newItem(value, priority)
	pq.items = append(pq.items, item)
	pq.Count++
	pq.proclateUp(pq.Count)
}

func (pq *PQueue) Print() {
	n := pq.Count
	for i := 1; i <= n; i++ {
		fmt.Print(*(pq.items[i]), " ")
	}
	fmt.Println()
}

func (pq *PQueue) Remove() (interface{}, bool) {
	if pq.Empty() {
		fmt.Println("Heap Empty Error.")
		return 0, false
	}

	value := pq.items[1].value
	pq.items[1] = pq.items[pq.Count]
	pq.items = pq.items[0:pq.Count]
	pq.Count--
	pq.proclateDown(1)
	return value, true
}

func (pq *PQueue) Empty() bool {
	return (pq.Count == 0)
}

func (pq *PQueue) Size() int {
	return pq.Count
}

func (pq *PQueue) Peek() (interface{}, bool) {
	if pq.Empty() {
		fmt.Println("Heap empty exception.")
		return 0, false
	}
	return pq.items[1].value, true
}

func main() {
	a := []int{1, 9, 6, 7, 8, -1, 2, 4, 5, 3}
	pq := CreatePQueue(false)
	for _, val := range a {
		pq.Add(strconv.Itoa(val), val)
	}

	for _, _ = range a {
		value, _ := pq.Remove()
		fmt.Print(value, " ")
	}
	fmt.Println()
	pq2 := CreatePQueue(false)
	for _, val := range a {
		pq2.Add(strconv.Itoa(val), val)
	}
	for _, _ = range a {
		value, _ := pq2.Remove()
		fmt.Print(value, " ")
	}
	fmt.Println()
}
