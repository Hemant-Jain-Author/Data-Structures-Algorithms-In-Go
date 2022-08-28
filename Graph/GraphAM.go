package main

import (
	"fmt"
	"math"
	"strconv"
)

type Graph struct {
	count int
	adj   [][]int
}

func NewGraph(count int) (gph *Graph) {
	gph = new(Graph)
	gph.count = count
	gph.adj = make([][]int, count)
	for i := range gph.adj {
		gph.adj[i] = make([]int, count)
	}
	return
}

func (gph *Graph) AddDirectedEdge(src int, dst int, cost int) {
	gph.adj[src][dst] = cost
}

func (gph *Graph) AddUndirectedEdge(src int, dst int, cost int) {
	gph.AddDirectedEdge(src, dst, cost)
	gph.AddDirectedEdge(dst, src, cost)
}

func (gph *Graph) Print() {
	for i := 0; i < gph.count; i++ {
		fmt.Print("Node index ", i, " is connected to : ")
		for j := 0; j < gph.count; j++ {
			if gph.adj[i][j] != 0 {
				fmt.Print(j, "(cost:", gph.adj[i][j], ") ")
			}
		}
		fmt.Println("")
	}
}

// Testing code.
func main1() {
	gph := NewGraph(4)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(2, 3, 1)
	gph.Print()
}

/*
Node index 0 is connected to : 1(cost:1) 2(cost:1)
Node index 1 is connected to : 0(cost:1) 2(cost:1)
Node index 2 is connected to : 0(cost:1) 1(cost:1) 3(cost:1)
Node index 3 is connected to : 2(cost:1)
*/

func (gph *Graph) Dijkstra(source int) {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)

	for i := 0; i < gph.count; i++ {
		previous[i] = -1
		dist[i] = math.MaxInt32 // infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = source

	type Edge struct {
		index    int
		cost int
	}

	cmp := func(a, b interface{}) bool { // Greater function
		return a.(Edge).cost > b.(Edge).cost
	}
	hp := CreateHeap(cmp)
	hp.Add(Edge{source, 0})

	for hp.Size() != 0 {
		curr := hp.Remove().(Edge).index

		if visited[curr]  {
			continue
		}
		visited[curr] = true

		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[curr][dest]
			if cost != 0 {
				alt := cost + dist[curr]
				if dist[dest] > alt && !visited[dest] {
					dist[dest] = alt
					previous[dest] = curr
					hp.Add(Edge{dest, alt})
				}
			}
		}
	}

	// Printing result.
	gph.printPath(previous, dist, count, source)
}

func (gph *Graph) printPathUtil(previous []int, source int, dest int) string {
	var path = "";
	if (dest == source){
		path += strconv.Itoa(source);
	} else {
		path += gph.printPathUtil(previous, source, previous[dest]);
		path += ("->" + strconv.Itoa(dest));
	}
	return path;
}

func (gph *Graph) printPath(previous []int, dist []int, count int, source int) {
	output := "Shortest Paths: ";
	for i := 0; i < count; i++ {
		if (dist[i] == 99999) {
			output += ("(" + strconv.Itoa(source) + "->" + strconv.Itoa(i) + " @ Unreachable) ");
		} else if (i != previous[i]) {
			output += "(";
			output += gph.printPathUtil(previous, source, i);
			output += (" @ " + strconv.Itoa(dist[i]) + ") ");
		}
	}
	fmt.Println(output);
}

func (gph *Graph) PrimsMST() {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)

	for i := 0; i < gph.count; i++ {
		previous[i] = -1
		dist[i] = math.MaxInt32 // infinite
		visited[i] = false
	}

	source := 0
	dist[source] = 0
	previous[source] = source

	type Edge struct {
		index    int
		cost int
	}

	cmp := func(a, b interface{}) bool {
		return a.(Edge).cost > b.(Edge).cost
	}

	hp := CreateHeap(cmp)
	hp.Add(Edge{source, 0})
	for hp.Size() != 0 {
		source := hp.Remove().(Edge).index

		if visited[source]  {
			continue
		}
		visited[source] = true

		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[source][dest]
			if cost != 0 {
				if dist[dest] > cost && !visited[dest] {
					dist[dest] = cost
					previous[dest] = source
					hp.Add(Edge{dest, cost})
				}
			}
		}
	}

	// Printing result.
	total := 0
	output := "Edges are : "
	for i := 0; i < count; i++ {
		if dist[i] == math.MaxInt32 {
			output += "( " + strconv.Itoa(i) + ",  Unreachable)"
		} else if (previous[i] != i) {
			output += "( " + strconv.Itoa(previous[i]) + "->" + strconv.Itoa(i) + " @ " + strconv.Itoa(dist[i]) + ") " 
			total += dist[i]
		}   
	}
	fmt.Println(output)
    fmt.Println("Total MST cost :", total)

}

// Testing code.
func main2() {
	gph := NewGraph(9)
	gph.AddUndirectedEdge(0, 1, 4)
	gph.AddUndirectedEdge(0, 7, 8)
	gph.AddUndirectedEdge(1, 2, 8)
	gph.AddUndirectedEdge(1, 7, 11)
	gph.AddUndirectedEdge(2, 3, 7)
	gph.AddUndirectedEdge(2, 8, 2)
	gph.AddUndirectedEdge(2, 5, 4)
	gph.AddUndirectedEdge(3, 4, 9)
	gph.AddUndirectedEdge(3, 5, 14)
	gph.AddUndirectedEdge(4, 5, 10)
	gph.AddUndirectedEdge(5, 6, 2)
	gph.AddUndirectedEdge(6, 7, 1)
	gph.AddUndirectedEdge(6, 8, 6)
	gph.AddUndirectedEdge(7, 8, 7)
	gph.Dijkstra(1)
	gph.PrimsMST()
}

/*
Shortest Paths: (1->0 @ 4) (1->2 @ 8) (1->2->3 @ 15) (1->2->5->4 @ 22) (1->2->5 @ 12) (1->7->6 @ 12) (1->7 @ 11) (1->2->8 @ 10) 

Edges are : ( 0->1 @ 4) ( 5->2 @ 4) ( 2->3 @ 7) ( 3->4 @ 9) ( 6->5 @ 2) ( 7->6 @ 1) ( 0->7 @ 8) ( 2->8 @ 2) 
Total MST cost : 37
*/

func (gph *Graph) hamiltonianPath() bool {
	count := gph.count
	path := make([]int, count)
	added := make([]int, count)

	if gph.hamiltonianPathUtil(path, 0, added) {
		fmt.Print("Hamiltonian Path found :: ")
		for i := 0; i < count; i++ {
			fmt.Print(" ", path[i])
		}
		return true
	}
	fmt.Print("Hamiltonian Path not found")
	return false
}

func (gph *Graph) hamiltonianPathUtil(path []int, pSize int, added []int) bool {
	// Base case full length path is found
	if pSize == gph.count {
		return true
	}
	for vertex := 0; vertex < gph.count; vertex++ {
		// there is a path from last element and next vertex
		// and next vertex is not already included in path.
		if pSize == 0 || (gph.adj[path[pSize-1]][vertex] == 1 && added[vertex] == 0) {
			path[pSize] = vertex
			pSize++
			added[vertex] = 1
			if gph.hamiltonianPathUtil(path, pSize, added) {
				return true
			}
			// backtracking
			pSize--
			added[vertex] = 0
		}
	}
	return false
}

// Testing code.
func main3() {
	count := 5
	gph := NewGraph(count)

	adj := make([][]int, 5)
	adj[0] = []int{0, 1, 0, 1, 0}
	adj[1] = []int{1, 0, 1, 1, 0}
	adj[2] = []int{0, 1, 0, 0, 1}
	adj[3] = []int{1, 1, 0, 0, 1}
	adj[4] = []int{0, 1, 1, 1, 0}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if adj[i][j] == 1 {
				gph.AddDirectedEdge(i, j, 1)
			}
		}
	}

	fmt.Println("\nhamiltonianPath : ", gph.hamiltonianPath())

	gph2 := NewGraph(count)

	adj2 := make([][]int, 5)
	adj2[0] = []int{0, 1, 0, 1, 0}
	adj2[1] = []int{1, 0, 1, 1, 0}
	adj2[2] = []int{0, 1, 0, 0, 1}
	adj2[3] = []int{1, 1, 0, 0, 0}
	adj2[4] = []int{0, 1, 1, 0, 0}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if adj2[i][j] == 1 {
				gph2.AddDirectedEdge(i, j, 1)
			}
		}
	}

	fmt.Println("\nhamiltonianPath :  ", gph2.hamiltonianPath())
}

/*
Hamiltonian Path found ::  0 1 2 4 3
hamiltonianPath :  true
Hamiltonian Path found ::  0 3 1 2 4
hamiltonianPath :   true
*/

func (gph *Graph) hamiltonianCycle() bool {
	count := gph.count
	path := make([]int, count+1)
	added := make([]int, count)

	if gph.hamiltonianCycleUtil(path, 0, added) {
		fmt.Print("Hamiltonian Cycle found :: ")
		for i := 0; i <= gph.count; i++ {
			fmt.Print(" ", path[i])
		}
		return true
	}
	fmt.Print("Hamiltonian Cycle not found")
	return false
}

func (gph *Graph) hamiltonianCycleUtil(path []int, pSize int, added []int) bool {
	// Base case full length path is found this last check can be modified to make it a path.
	count := gph.count
	if pSize == count {
		if gph.adj[path[pSize-1]][path[0]] == 1 {
			path[pSize] = path[0]
			return true
		} else {
			return false
		}
	}
	for vertex := 0; vertex < gph.count; vertex++ {
		// there is a path from last element and next vertex
		if pSize == 0 || (gph.adj[path[pSize-1]][vertex] == 1 && added[vertex] == 0) {
			path[pSize] = vertex
			pSize++
			added[vertex] = 1
			if gph.hamiltonianCycleUtil(path, pSize, added) {
				return true
			}
			// backtracking
			pSize--
			added[vertex] = 0
		}
	}
	return false
}

// Testing code.
func main4() {
	count := 5
	gph := NewGraph(count)

	adj := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{0, 1, 1, 1, 0}}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if adj[i][j] == 1 {
				gph.AddDirectedEdge(i, j, 1)
			}
		}
	}

	fmt.Println("\nhamiltonianCycle : ", gph.hamiltonianCycle())

	gph2 := NewGraph(count)

	adj2 := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0}}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if adj2[i][j] == 1 {
				gph2.AddDirectedEdge(i, j, 1)
			}
		}
	}

	fmt.Println("\nhamiltonianCycle :  ", gph2.hamiltonianCycle())
}

/*
Hamiltonian Cycle found ::  0 1 2 4 3 0
hamiltonianCycle :  true
Hamiltonian Cycle not found
hamiltonianCycle :   false
*/

func main() {
	main1()
	main2()
	main3()
	main4()
}

type Heap struct {
	size  int
	arr   []interface{}
	comp func(x interface{}, y interface{}) bool
}

func CreateHeap(comp func(x interface{}, y interface{}) bool) *Heap {
	var arr []interface{}
	h := &Heap{comp: comp, arr : arr, size : 0}
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
