package main
/*
import (
	"container/heap"
	"fmt"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
	"math"
)

type Edge struct {
	source      int
	destination int
	cost        int
	next        *Edge
}

type AdjList struct {
	head *Edge
}

type Graph struct {
	count      int
	ListVector [](*AdjList)
}

func (g *Graph) Init(cnt int) {
	g.count = cnt
	g.ListVector = make([]*AdjList, cnt)
	for i := 0; i < cnt; i++ {
		g.ListVector[i] = new(AdjList)
		g.ListVector[i].head = nil
	}
}

func (g *Graph) AddEdge(source int, destination int, cost int) {
	edge := new(Edge)
	edge.source = source
	edge.destination = destination
	edge.cost = cost
	edge.next = g.ListVector[source].head
	g.ListVector[source].head = edge
}

func (g *Graph) AddEdgeUnweighted(source int, destination int) {
	g.AddEdge(source, destination, 1)
}

func (g *Graph) AddBiEdge(source int, destination int, cost int) {
	g.AddEdge(source, destination, cost)
	g.AddEdge(destination, source, cost)
}

func (g *Graph) AddBiEdgeUnweghted(source int, destination int) {
	g.AddBiEdge(source, destination, 1)
}

func (g *Graph) Print() {
	for i := 0; i < g.count; i++ {
		ad := g.ListVector[i].head
		if ad != nil {
			fmt.Print("Vertex ", i, " is connected to : ")
			for ad != nil {
				fmt.Print("[", ad.destination, ad.cost, "] ")
				ad = ad.next
			}
			fmt.Println()
		}
	}
}


func main() {
	g := new(Graph)
	g.Init(9)
	g.AddBiEdge(0, 2, 1)
	g.AddBiEdge(1, 2, 5)
	g.AddBiEdge(1, 3, 7)
	g.AddBiEdge(1, 4, 9)
	g.AddBiEdge(3, 2, 2)
	g.AddBiEdge(3, 5, 4)
	g.AddBiEdge(4, 5, 6)
	g.AddBiEdge(4, 6, 3)
	g.AddBiEdge(5, 7, 1)
	g.AddBiEdge(6, 7, 7)
	g.AddBiEdge(7, 8, 17)
	g.Print()
	g.Dijkstra(1)
	//g.Prims();
}


type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(val interface{}, priority int) {
	edge := val.(*Edge)
	*pq = append(*pq, edge)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	edge := old[n-1]
	*pq = old[0 : n-1]
	return edge
}

func (g *Graph) Dijkstra(source int) {
	count := g.count
	previous := make([]int, count)
	dist := make([]int, count)
	que := make(PriorityQueue, 0)
	heap.Init(&que)

	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = Infinite
	}

	dist[source] = 0
	edge := new(Edge)
	edge.cost = 0
	edge.source = source
	edge.destination = source
	heap.Push(&que, edge)

	for que.Len() != 0 {
		edge = heap.Pop(&que).(*Edge) // Pop from PQ

		if dist[edge.destination] < edge.cost {
			continue
		}

		dist[edge.destination] = edge.cost
		previous[edge.destination] = edge.source

		adl := g.ListVector[edge.destination]
		adn := adl.head
		for adn != nil {
			if previous[adn.destination] == -1 {
				edge = new(Edge)
				edge.source = adn.source
				edge.destination = adn.destination
				edge.cost = adn.cost + dist[adn.source]
				heap.Push(&que, edge)
			}
			adn = adn.next
		}
	}

	for i := 0; i < count; i++ {
		if dist[i] == Infinite {
			fmt.Println(" edge id ", i, "  prev ", previous[i], " distance : Unreachable")
		} else {
			fmt.Println(" edge id ", i, "  prev ", previous[i], " distance : ", dist[i])
		}
	}
}

func (g *Graph) Prims() {
	count := g.count
	previous := make([]int, count)
	dist := make([]int, count)
	que := make(PriorityQueue, 0)
	heap.Init(&que)

	source := 1

	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = Infinite
	}

	dist[source] = 0
	edge := new(Edge)
	edge.cost = 0
	edge.source = source
	edge.destination = source
	heap.Push(&que, edge)

	for que.Len() != 0 {
		edge = heap.Pop(&que).(*Edge) // Pop from PQ

		if dist[edge.destination] < edge.cost {
			continue
		}

		dist[edge.destination] = edge.cost
		previous[edge.destination] = edge.source

		adl := g.ListVector[edge.destination]
		adn := adl.head
		for adn != nil {
			if previous[adn.destination] == -1 {
				edge = new(Edge)
				edge.source = adn.source
				edge.destination = adn.destination
				edge.cost = adn.cost
				heap.Push(&que, edge)
			}
			adn = adn.next
		}
	}

	// Printing result.
	for i := 0; i < count; i++ {
		if dist[i] == Infinite {
			fmt.Println(" edge id ", i, "  prev ", previous[i], " distance : Unreachable")
		} else {
			fmt.Println(" edge id ", i, "  prev ", previous[i], " distance : ", dist[i])
		}
	}
}

func main5() {
	g := new(Graph)
	g.Init(6)
	g.AddEdge(5, 2, 1)
	g.AddEdge(5, 0, 1)
	g.AddEdge(4, 0, 1)
	g.AddEdge(4, 1, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 1, 1)
	g.TopologicalSort()
}

func (g *Graph) TopologicalSortDFS(index int, visited []int, stk *stack.Stack) {
	head := g.ListVector[index].head
	for head != nil {
		if visited[head.destination] == 0 {
			visited[head.destination] = 1
			g.TopologicalSortDFS(head.destination, visited, stk)
		}
		head = head.next
	}
	stk.Push(index)
}

func (g *Graph) TopologicalSort() {
	fmt.Print("Topological order of given graph is : ")
	var count = g.count
	stk := new(stack.Stack)
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			g.TopologicalSortDFS(i, visited, stk)
		}
	}

	for stk.Len() != 0 {
		fmt.Print(stk.Pop().(int)," ")
	}
	fmt.Println()
}


func main() {
	g := new(Graph)
	g.Init(9)
	g.AddEdge(0, 2, 1)
	g.AddBiEdge(1, 2, 5)
	g.AddEdge(1, 3, 7)
	g.AddEdge(1, 4, 9)
	g.AddEdge(3, 2, 2)
	g.AddEdge(3, 5, 4)
	g.AddEdge(4, 5, 6)
	g.AddEdge(4, 6, 3)
	g.AddEdge(5, 7, 1)
	g.AddEdge(6, 7, 7)
	g.AddEdge(7, 8, 17)
	g.Print()
	// fmt.Println(g.PathExist(1, 0))
	// fmt.Println(g.PathExist(0, 1))
	// fmt.Println(g.PathExist(1, 2))
 	// g.DFS()
	// g.DFSStack()
	// g.BFS()
	// fmt.Println(g.IsConnected())
	// g.ShortestPath(1)
	g.BellmanFordShortestPath(1)
}

func (g *Graph) PathExist(source int, destination int) bool {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[source] = 1
	g.DFSRec(source, visited)
	return visited[destination] != 0
}

func (g *Graph) DFSRec(index int, visited []int) {
	head := g.ListVector[index].head
	fmt.Println(index)
	for head != nil {
		if visited[head.destination] == 0 {
			//fmt.Println(index)
			visited[head.destination] = 1
			g.DFSRec(head.destination, visited)
		}
		head = head.next
	}
}

func (g *Graph) DFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			visited[i] = 1
			g.DFSRec(i, visited)
		}
	}
}

func (g *Graph) DFSStack() {
	count := g.count
	visited := make([]int, count)
	var curr int
	stk := new(stack.Stack)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}

	visited[1] = 1
	stk.Push(1)

	for stk.Len() != 0 {
		curr = stk.Pop().(int)
		// fmt.Println(curr)
		head := g.ListVector[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				stk.Push(head.destination)
			}
			head = head.next
		}
	}
}

func (g *Graph) BFSQueue(index int, visited []int) {
	var curr int
	que := new(queue.Queue)

	visited[index] = 1
	que.Enqueue(index)

	for que.Len() != 0 {
		curr = que.Dequeue().(int)
		head := g.ListVector[curr].head
		for head != nil {
			if visited[head.destination] == 0 {
				visited[head.destination] = 1
				que.Enqueue(head.destination)
			}
			head = head.next
		}
	}
}

func (g *Graph) BFS() {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			g.BFSQueue(i, visited)
		}
	}
}

func (g *Graph) IsConnected() bool {
	count := g.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		visited[i] = 0
	}
	visited[0] = 1
	g.DFSRec(0, visited)
	
	for i := 0; i < count; i++ {
		if visited[i] == 0 {
			return false
		}
	}
	return true
}

func (g *Graph) ShortestPath(source int) {
	var curr int
	count := g.count
	distance := make([]int, count)
	path := make([]int, count)

	que := new(queue.Queue)
	for i := 0; i < count; i++ {
		distance[i] = -1
	}
	que.Enqueue(source)
	distance[source] = 0
	path[source] = source
	for que.Len() != 0 {
		curr = que.Dequeue().(int)
		head := g.ListVector[curr].head
		for head != nil {
			if distance[head.destination] == -1 {
				distance[head.destination] = distance[curr] + 1
				path[head.destination] = curr
				que.Enqueue(head.destination)
			}
			head = head.next
		}
	}
	for i := 0; i < count; i++ {
		fmt.Println(path[i], " to ", i, " weight ", distance[i])
	}
}
const Infinite = math.MaxInt32

func (g *Graph) BellmanFordShortestPath(source int) {
	count := g.count
	distance := make([]int, count)
	path := make([]int, count)

	for i := 0; i < count; i++ {
		distance[i] = Infinite
	}
	distance[source] = 0
	path[source] = source
	for i := 0; i < count-1; i++ {
		for j := 0; j < count; j++ {
			head := g.ListVector[j].head
			for head != nil {
				newDistance := distance[j] + head.cost
				if distance[head.destination] > newDistance {
					distance[head.destination] = newDistance
					path[head.destination] = j
				}
				head = head.next
			}
		}
	}
	for i := 0; i < count; i++ {
		fmt.Println(path[i], " to ", i, " weight ", distance[i])
	}
}
*/