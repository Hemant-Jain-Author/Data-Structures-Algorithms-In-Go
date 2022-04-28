package main

import (
	"container/heap"
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

	type Item struct {
		index    int
		priority int
	}

	cmp := func(a, b interface{}) bool { // Less function
		return a.(Item).priority < b.(Item).priority
	}
	hp := NewHeap(cmp)
	heap.Push(hp, Item{source, 0})

	for hp.Len() != 0 {
		curr := heap.Pop(hp).(Item).index

		if visited[curr] == true {
			continue
		}
		visited[curr] = true

		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[curr][dest]
			if cost != 0 {
				alt := cost + dist[curr]
				if dist[dest] > alt && visited[dest] == false {
					dist[dest] = alt
					previous[dest] = curr
					heap.Push(hp, Item{dest, alt})
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

	type Item struct {
		index    int
		priority int
	}

	cmp := func(a, b interface{}) bool {
		return a.(Item).priority < b.(Item).priority
	}

	hp := NewHeap(cmp)
	heap.Push(hp, Item{source, 0})
	for hp.Len() != 0 {
		source := heap.Pop(hp).(Item).index

		if visited[source] == true {
			continue
		}
		visited[source] = true

		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[source][dest]
			if cost != 0 {
				if dist[dest] > cost && visited[dest] == false {
					dist[dest] = cost
					previous[dest] = source
					heap.Push(hp, Item{dest, cost})
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

// *********************
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

// ************************************
