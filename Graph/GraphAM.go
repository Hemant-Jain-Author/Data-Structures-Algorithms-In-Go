package main

import (
	"fmt"
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
		fmt.Print("Node index " , i , " is connected to : ")
		for j := 0; j < gph.count; j++ {
			if (gph.adj[i][j] != 0) {
				fmt.Print(j , "(cost:", gph.adj[i][j], ") ")
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
		dist[i] = 999999 // infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = -1

	que := new(PQueue)
	que.Init()
	que.Add(source, 0)

	for que.Len() != 0 {
		source := que.Remove().(int)
		
		if (visited[source] == true) {
            continue
        }
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

func (gph *Graph) Prims() {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)
	source := 0

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

	for que.Len() != 0 {
		source := que.Remove().(int)
		
		if (visited[source] == true) {
            continue
        }
		visited[source] = true
		
		for dest := 0; dest < gph.count; dest++ {
			cost := gph.adj[source][dest]
			if (cost != 0) {
				if (dist[dest] > cost && visited[dest] == false) {
					dist[dest] = cost
					previous[dest] = source
					que.Add(dest, cost)
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
	//gph.Print()
	//gph.Dijkstra(1)
	gph.Prims()
}

/*
 node id  0   prev  1  distance :  4
 node id  1   prev  -1  distance :  0
 node id  2   prev  1  distance :  8
 node id  3   prev  2  distance :  15
 node id  4   prev  5  distance :  22
 node id  5   prev  2  distance :  12
 node id  6   prev  7  distance :  12
 node id  7   prev  1  distance :  11
 node id  8   prev  2  distance :  10

 node id  0   prev  -1  distance :  0
 node id  1   prev  0  distance :  4
 node id  2   prev  5  distance :  4
 node id  3   prev  2  distance :  7
 node id  4   prev  3  distance :  9
 node id  5   prev  6  distance :  2
 node id  6   prev  7  distance :  1
 node id  7   prev  0  distance :  8
 node id  8   prev  2  distance :  2
*/

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
	fmt.Print("Hamiltonian Path not found")
	return false
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


func main3() {
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
	
	fmt.Println("\nhamiltonianPath : " , gph.hamiltonianPath())

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

	fmt.Println("\nhamiltonianPath :  " , gph2.hamiltonianPath())
}

/*
Hamiltonian Path found ::  0 1 2 4 3
hamiltonianPath :  true
Hamiltonian Path found ::  0 3 1 2 4
hamiltonianPath :   true
*/

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
	fmt.Print("Hamiltonian Cycle not found")
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
	
	fmt.Println("\nhamiltonianCycle : " , gph.hamiltonianCycle())

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

	fmt.Println("\nhamiltonianCycle :  " , gph2.hamiltonianCycle())
}

/*
Hamiltonian Cycle found ::  0 1 2 4 3 0
hamiltonianCycle :  true
Hamiltonian Cycle not found
hamiltonianCycle :   false
*/

func main(){
	//main1()
	//main2()
	//main3()
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