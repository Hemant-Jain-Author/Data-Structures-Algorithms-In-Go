package main

import (
	"fmt"
	// "math"
	// "github.com/golang-collections/collections/queue"
	// "github.com/golang-collections/collections/stack"
	"container/heap"
)

type Graph struct {
	count int
	adj [][]int
}

func (gph *Graph) Init(count int) {
	gph.count = count
	gph.adj = make([][]int, count)
	for i := range gph.adj {
        gph.adj[i] = make([]int, count)
    }
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
		fmt.Print("Node index [ " , i , " ] is connected with : ")
		for j := 0; j < gph.count; j++ {
			if (gph.adj[i][j] != 0) {
				fmt.Print(j , " ")
			}
		}
		fmt.Println("")
	}
}

func main1() {
	gph := new(Graph)
	gph.Init(4)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(2, 3, 1)
	gph.Print()
}

func (gph *Graph) dijkstra(source int) {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)

	for i := 0; i < gph.count; i++ {
		previous[i] = -1
		dist[i] = 999999 // infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = -1

	que := new(PQueue)
	que.Init()
	que.Add(source, 0)

	for (que.IsEmpty() != true) {
		source := que.Remove().(int)
		visited[source] = true
		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[source][dest]
			if (cost != 0) {
				alt := cost + dist[source]
				if (dist[dest] > alt && visited[dest] == false) {
					dist[dest] = alt
					previous[dest] = source
					que.Add(dest, alt)
				}
			}
		}
	}

	for i := 0; i < count; i++ {
		if (dist[i] == 999999) {
			fmt.Println(" node id " , i , "  prev " , previous[i] , " distance : Unreachable")
		} else {
			fmt.Println(" node id " , i , "  prev " , previous[i] , " distance : " , dist[i])
		}
	}
}

func (gph *Graph) prims() {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)
	source := 0

	for i := 0; i < gph.count; i++ {
		previous[i] = -1
		dist[i] = 999999// infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = -1

	que := new(PQueue)
	que.Init()
	que.Add(source, 0)

	for (que.IsEmpty() != true) {
		source := que.Remove().(int)
		visited[source] = true
		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[source][dest]
			if (cost != 0) {
				alt := cost
				if (dist[dest] > alt && visited[dest] == false) {
					dist[dest] = alt
					previous[dest] = source
					que.Add(dest, alt)
				}
			}
		}
	}

	for i := 0; i < count; i++ {
		if (dist[i] == 999999) {
			fmt.Println(" node id " , i , "  prev " , previous[i] , " distance : Unreachable")
		} else {
			fmt.Println(" node id " , i , "  prev " , previous[i] , " distance : " , dist[i])
		}
	}
}

func main2() {
	gph := new(Graph)
	gph.Init(9)
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
	gph.Print()
	gph.prims()
	gph.dijkstra(0)
}

func main3() {
	gph := new(Graph)
	gph.Init(9)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(1, 2, 5)
	gph.AddUndirectedEdge(1, 3, 7)
	gph.AddUndirectedEdge(1, 4, 9)
	gph.AddUndirectedEdge(3, 2, 2)
	gph.AddUndirectedEdge(3, 5, 4)
	gph.AddUndirectedEdge(4, 5, 6)
	gph.AddUndirectedEdge(4, 6, 3)
	gph.AddUndirectedEdge(5, 7, 1)
	gph.AddUndirectedEdge(6, 7, 7)
	gph.AddUndirectedEdge(7, 8, 17)
	gph.Print()
	gph.prims()
	gph.dijkstra(1)
}

func (gph *Graph) hamiltonianPathUtil(path []int, pSize int, added []int) bool {
	// Base case full length path is found
	if (pSize == gph.count) {
		return true
	}
	for vertex := 0; vertex < gph.count; vertex++ {
		// there is a path from last element and next vertex
		// and next vertex is not already included in path.
		if (pSize == 0 || (gph.adj[path[pSize - 1]][vertex] == 1 && added[vertex] == 0)) {
			path[pSize] = vertex
			pSize++
			added[vertex] = 1
			if (gph.hamiltonianPathUtil(path, pSize, added)) {
				return true
			}
			// backtracking
			pSize--
			added[vertex] = 0
		}
	}
	return false
}

func (gph *Graph)hamiltonianPath() bool {
	count := gph.count
	path := make([]int, count)
	added := make([]int, count)

	if (gph.hamiltonianPathUtil(path, 0, added)) {
		fmt.Print("Hamiltonian Path found :: ")
		for i := 0; i < count; i++ {
			fmt.Print(" " , path[i])
		}
		return true
	}
	fmt.Println("Hamiltonian Path not found")
	return false
}

func (gph *Graph)hamiltonianCycleUtil(path []int, pSize int, added []int) bool {
	// Base case full length path is found this last check can be modified to make it a path.
	count := gph.count
	if (pSize == count) {
		if (gph.adj[path[pSize - 1]][path[0]] == 1) {
			path[pSize] = path[0]
			return true
		} else {
			return false
		}
	}
	for vertex := 0; vertex < gph.count; vertex++ {
		// there is a path from last element and next vertex
		if (pSize == 0 || (gph.adj[path[pSize - 1]][vertex] == 1 && added[vertex] == 0)) {
			path[pSize] = vertex
			pSize++
			added[vertex] = 1
			if (gph.hamiltonianCycleUtil(path, pSize, added)) {
				return true
			}
			// backtracking
			pSize--
			added[vertex] = 0
		}
	}
	return false
}

func (gph *Graph)hamiltonianCycle() bool {
	count := gph.count
	path := make([]int, count+1)
	added := make([]int, count)
	
	if (gph.hamiltonianCycleUtil(path, 0, added)) {
		fmt.Print("Hamiltonian Cycle found :: ")
		for i := 0; i <= gph.count; i++ {
			fmt.Print(" " , path[i])
		}
		return true
	}
	fmt.Println("Hamiltonian Cycle not found")
	return false
}

func main4() {
	count := 5
	gph := new(Graph)
	gph.Init(count)
	
	adj := make([][]int, 5)
    adj[0] = []int{0, 1, 0, 1, 0} 
    adj[1] = []int{1, 0, 1, 1, 0}
    adj[2] = []int{0, 1, 0, 0, 1}
    adj[3] = []int{1, 1, 0, 0, 1}
    adj[4] = []int{0, 1, 1, 1, 0}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if (adj[i][j] == 1) {
				gph.AddDirectedEdge(i, j, 1)
			}
		}
	}
	
	fmt.Println("hamiltonianPath : " , gph.hamiltonianPath())
	fmt.Println("hamiltonianCycle : " , gph.hamiltonianCycle())

	gph2 := new(Graph)
	gph2.Init(count)

	adj2 := make([][]int, 5)
    adj2[0] = []int{0, 1, 0, 1, 0} 
    adj2[1] = []int{1, 0, 1, 1, 0}
    adj2[2] = []int{0, 1, 0, 0, 1}
    adj2[3] = []int{1, 1, 0, 0, 0}
    adj2[4] = []int{0, 1, 1, 0, 0}

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if (adj2[i][j] == 1) {
				gph2.AddDirectedEdge(i, j, 1)
			}
		}
	}

	fmt.Println("hamiltonianPath :  " , gph2.hamiltonianPath())
	fmt.Println("hamiltonianCycle :  " , gph2.hamiltonianCycle())
}

func main(){
	main1()
	main2()
	main3()
	main4()
}

// *********************
type Item struct {
	value    interface{}
	priority int
}

type ItemList []*Item

func (lst ItemList) Len() int {
	return len(lst)
}

func (lst ItemList) Less(i, j int) bool {
	return lst[i].priority < lst[j].priority
}

func (lst ItemList) Swap(i, j int) {
	lst[i], lst[j] = lst[j], lst[i]
}

func (lst *ItemList) Push(val interface{}) {
	item := val.(*Item)
	*lst = append(*lst, item)
}

func (lst *ItemList) Pop() interface{} {
	old := *lst
	n := len(old)
	item := old[n-1]
	*lst = old[0 : n-1]
	return item
}

type PQueue struct {
	pq ItemList
}

func NewPQueue() *PQueue {
	queue := new(PQueue)
	queue.pq = make(ItemList, 0)
	heap.Init(&queue.pq)
	return queue
}

func (queue *PQueue) Init() {
	queue.pq = make(ItemList, 0)
	heap.Init(&queue.pq)
}

func (queue *PQueue) Add(value interface{}, priority int) {
	heap.Push(&queue.pq, &Item{value: value, priority: priority})
}

func (queue *PQueue) Remove() interface{} {
	return heap.Pop(&queue.pq).(*Item).value
}

func (queue *PQueue) Len() int {
	return queue.pq.Len()
}

func (queue *PQueue) IsEmpty() bool {
	return queue.pq.Len() == 0
}

// ************************************