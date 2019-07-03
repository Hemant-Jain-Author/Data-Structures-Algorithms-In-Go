package main

import (
	"fmt"
	"math"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"container/heap"
)

type Edge struct {
	destination int
	cost        int
	next        *Edge
}

type Graph struct {
	count      int
	Edges [](*Edge)
}

func (gph *Graph) Init(cnt int) {
	gph.count = cnt
	gph.Edges = make([]*Edge, cnt)
}

func (gph *Graph) AddDirectedEdge(source int, destination int, cost int) {
	edge := &Edge{destination, cost, gph.Edges[source]}
	gph.Edges[source] = edge
}

func (gph *Graph) AddDirectedEdgeUnweighted(source int, destination int) {
	gph.AddDirectedEdge(source, destination, 1)
}

func (gph *Graph) AddUndirectedEdge(source int, destination int, cost int) {
	gph.AddDirectedEdge(source, destination, cost)
	gph.AddDirectedEdge(destination, source, cost)
}

func (gph *Graph) AddUndirectedEdgeUnweghted(source int, destination int) {
	gph.AddUndirectedEdge(source, destination, 1)
}

func (gph *Graph) Print() {
	for i := 0; i < gph.count; i++ {
		ad := gph.Edges[i]
		fmt.Print("Vertex ", i, " is connected to : ")
		for ad != nil {
			fmt.Print( ad.destination, "(cost:", ad.cost, ") ")
			ad = ad.next
		}
		fmt.Println()
	}
}

//Testing code
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
Vertex 0 is connected to : 2(cost:1) 1(cost:1) 
Vertex 1 is connected to : 2(cost:1) 0(cost:1) 
Vertex 2 is connected to : 3(cost:1) 1(cost:1) 0(cost:1) 
Vertex 3 is connected to : 2(cost:1)
*/

func (gph *Graph) DFSStack(source int, target int) bool {
	count := gph.count
	visited := make([]bool, count)
	var curr int
	stk := stack.New()
	path := []int{}
	visited[source] = true
	stk.Push(source)

	for stk.Len() != 0 {
		curr = stk.Pop().(int)
		path = append(path,curr)
		head := gph.Edges[curr]
		for head != nil {
			if visited[head.destination] == false {
				visited[head.destination] = true
				stk.Push(head.destination)
			}
			head = head.next
		}
	}
	fmt.Println("DFS Path is : ", path)
    return visited[target]
}

func (gph *Graph) DFS(source int, target int) bool {
	count := gph.count
	visited := make([]bool, count)

	fmt.Print("DFS Path is : ")
	gph.DFSUtil(source, visited)
	fmt.Println()
	return visited[target]
}

func (gph *Graph) DFSUtil(curr int, visited []bool) {
	visited[curr] = true
	head := gph.Edges[curr]
	//fmt.Print(curr, " ")
	for head != nil {
		if visited[head.destination] == false {
			gph.DFSUtil(head.destination, visited)
		}
		head = head.next
	}
}

func (gph *Graph) DFSUtil2(curr int, visited []bool, stk *stack.Stack) {
	visited[curr] = true
	head := gph.Edges[curr]

	for head != nil {
		if visited[head.destination] == false {
			gph.DFSUtil2(head.destination, visited, stk)
		}
		head = head.next
	}
	stk.Push(curr)
}

func (gph *Graph) BFS(source int, target int) bool {
	count := gph.count
	visited := make([]bool, count)
	que := queue.New()
	path := []int{}
	visited[source] = true
	que.Enqueue(source)
	
	for que.Len() != 0 {
		curr := que.Dequeue().(int)
		path = append(path, curr)
	
		head := gph.Edges[curr]
		for head != nil {
			if visited[head.destination] == false {
				visited[head.destination] = true
				que.Enqueue(head.destination)
			}
			head = head.next
		}
	}
	fmt.Println("BFS Path is : ", path)
    return visited[target]
}


//Testing code
func main2() {
	gph := new(Graph)
	gph.Init(8)
	gph.AddUndirectedEdge(0, 3, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(0, 1, 1)

	gph.AddUndirectedEdge(1, 4, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	gph.AddUndirectedEdge(3, 6, 1)

	gph.AddUndirectedEdge(6, 7, 1)
	gph.AddUndirectedEdge(5, 7, 1)
	gph.AddUndirectedEdge(4, 7, 1)
	
	fmt.Println("Path between 0 & 6 : ", gph.DFSStack(0, 6))
	fmt.Println("Path between 0 & 6 : ", gph.DFS(0, 6))
	fmt.Println("Path between 0 & 6 : ", gph.BFS(0, 6))
}

/*
FS Path is :  [0 3 6 7 5 4 2 1]
Path between 0 & 6 :  true
DFS Path is : 0 1 4 7 5 2 6 3 
Path between 0 & 6 :  true
BFS Path is :  [0 1 2 3 4 5 6 7]
Path between 0 & 6 :  true
*/

func (gph *Graph) TopologicalSort() {
	fmt.Print("Topological order of given graph is : ")
	var count = gph.count
	stk := stack.New()
	visited := make([]bool, count)

	for i := 0; i < count; i++ {
		if visited[i] == false {
			visited[i] = true
			gph.TopologicalSortDFS(i, visited, stk)
		}
	}
	for stk.Len() != 0 {
		fmt.Print(stk.Pop().(int), " ")
	}
	fmt.Println()
}

func (gph *Graph) TopologicalSortDFS(index int, visited []bool, stk *stack.Stack) {
	head := gph.Edges[index]
	for head != nil {
		if visited[head.destination] == false {
			visited[head.destination] = true
			gph.TopologicalSortDFS(head.destination, visited, stk)
		}
		head = head.next
	}
	stk.Push(index)
}


//Testing code
func main3() {
	gph := new(Graph)
	gph.Init(6)
	gph.AddDirectedEdge(5, 2, 1)
	gph.AddDirectedEdge(5, 0, 1)
	gph.AddDirectedEdge(4, 0, 1)
	gph.AddDirectedEdge(4, 1, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 1, 1)
	gph.TopologicalSort()
}
/*
Topological order of given graph is : 5 4 2 3 1 0 
*/

func (gph *Graph) PathExist(source int, destination int) bool {
	count := gph.count
	visited := make([]bool, count)
	visited[source] = true
	gph.DFSUtil(source, visited)
	return visited[destination]
}

func (gph *Graph)countAllPathDFS(visited[] bool, source int, dest int) int {
	if (source == dest){
		return 1
	}

	count := 0
	visited[source] = true
	head := gph.Edges[source]

	for head != nil {
		if visited[head.destination] == false {
			count += gph.countAllPathDFS(visited, head.destination, dest)
		}
		visited[head.destination] = false
		head = head.next
	}
	return count
}

func (gph *Graph)countAllPath(src int, dest int) int {
	count := gph.count
	visited := make([]bool, count)
	return gph.countAllPathDFS(visited, src, dest)
}

func (gph *Graph)printAllPath(src int, dest int){
	count := gph.count
	visited := make([]bool, count)
	path := stack.New()
	gph.printAllPathUtil(src, dest, visited, path)
}

func (gph *Graph)printAllPathUtil(source int, dest int, visited []bool, path *stack.Stack){
	path.Push(source)
	if (source == dest) {
		path.Print()
		fmt.Println()
		path.Pop()
		return
	}
	visited[source] = true
	head := gph.Edges[source]
	for head != nil {
		if visited[head.destination] == false {
			gph.printAllPathUtil(head.destination, dest, visited, path)
		}
		head = head.next
	}
	visited[source] = false
	path.Pop()
}

//Testing code
func main4() {
	gph := new(Graph)
	gph.Init(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.AddDirectedEdge(1, 4, 1)
	fmt.Println("PathExist :: " , gph.PathExist(0, 4))
	fmt.Println("Path Count :: " , gph.countAllPath(0, 4))
	gph.printAllPath(0, 4)
	fmt.Println("")
}
/*
PathExist ::  true
Path Count ::  3
4 3 2 0 
4 1 0 
4 3 1 0 
*/


func (gph *Graph) rootVertex() int {
	count := gph.count
	visited := make([]bool, count)
	retVal := -1

	for i := 0; i < count; i++ {
		if visited[i] == false {
			gph.DFSUtil(i, visited)
			retVal = i
		}
	}
	fmt.Println("Root vertex is :: ", retVal)
	return retVal
}

//Testing code
func main5() {
	gph := new(Graph)
	gph.Init(7)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(4, 1, 1)
	gph.AddDirectedEdge(6, 4, 1)
	gph.AddDirectedEdge(5, 6, 1)
	gph.AddDirectedEdge(5, 2, 1)
	gph.AddDirectedEdge(6, 0, 1)
	gph.rootVertex()
}

/*
Root vertex is ::  5
*/

func (gph *Graph) transitiveClosureUtil(source int, dest int, tc *[10][10]int) {
	tc[source][dest] = 1
	head := gph.Edges[dest]
	for head != nil {
		if (tc[source][head.destination] == 0) {
			gph.transitiveClosureUtil(source, head.destination, tc)
		}
		head = head.next
	}
}

func (gph *Graph) transitiveClosure() [10][10]int {
	count := gph.count
	tc := [10][10]int{}
	for i := 0; i < count; i++ {
		gph.transitiveClosureUtil(i, i, &tc)
	}
	return tc
}

//Testing code
func main6() {
	gph := new(Graph)
	gph.Init(4)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 3, 1)
	tc := gph.transitiveClosure()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(tc[i][j] , " ")
		}
		fmt.Println()
	}
}
/*
1 1 1 1 
1 1 1 1 
1 1 1 1 
0 0 0 1 
*/

func (gph *Graph)bfsLevelNode(source int) {
	count := gph.count
	visited := make([]bool, count)
	level := make([]int, count)
	visited[source] = true
	que := queue.New()

	que.Enqueue(source)
	level[source] = 0
	fmt.Println("\nNode  - Level")

	for que.Len() > 0 {
		curr := que.Dequeue().(int)
		depth := level[curr]
		head := gph.Edges[curr]
		fmt.Println(curr ," - " , depth)
		for head != nil {
			if (visited[head.destination] == false) {
				visited[head.destination] = true
				que.Enqueue(head.destination)
				level[head.destination] = depth + 1
			}
			head = head.next
		}
	}
}

func (gph *Graph)bfsDistance(source int, dest int) int {
	count := gph.count
	visited := make([]bool, count)
	que := queue.New()
	que.Enqueue(source)
	visited[source] = true
	level := make([]int, count)
	level[source] = 0

	for que.Len() > 0 {
		curr := que.Dequeue().(int)
		depth := level[curr]
		head := gph.Edges[curr]
		for head != nil {
			if (head.destination == dest) {
				return depth+1
			}
			if (visited[head.destination] == false) {
				visited[head.destination] = true
				que.Enqueue(head.destination)
				level[head.destination] = depth + 1
			}
			head = head.next
		}
	}
	return -1
}

//Testing code
func main7() {
	gph := new(Graph)
	gph.Init(7)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(0, 4, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	gph.AddUndirectedEdge(3, 4, 1)
	gph.AddUndirectedEdge(4, 5, 1)
	gph.AddUndirectedEdge(4, 6, 1)
	gph.bfsLevelNode(1)
	fmt.Println(gph.bfsDistance(1, 6))
}
/*

Node  - Level
1  -  0
2  -  1
0  -  1
5  -  2
4  -  2
6  -  3
3  -  3
3
*/

func (gph *Graph)isCyclePresentUndirectedDFS(index int, parentIndex int, visited[] bool) bool {
	visited[index] = true
	var dest int 
	head := gph.Edges[index]
	for head != nil {
		dest = head.destination		
		if (visited[dest] == false) {
			if (gph.isCyclePresentUndirectedDFS(dest, index, visited)) {
				return true
			}
		} else if (parentIndex != dest) {
			return true
		}
		head = head.next
	}
	return false
}

func (gph *Graph)isCyclePresentUndirected() bool {
	count := gph.count
	visited := make([]bool, count)
	for i := 0; i < count; i++ {
		if (visited[i] == false) {
			if (gph.isCyclePresentUndirectedDFS(i, -1, visited)) {
				return true
			}
		}
	}
	return false
}

//Testing code
func main8() {
	gph := new(Graph)
	gph.Init(6)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(3, 4, 1)
	gph.AddUndirectedEdge(4, 2, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	// gph.AddUndirectedEdge(4, 1, 1)
	// gph.AddUndirectedEdge(3, 5, 1)

	fmt.Println(gph.isCyclePresentUndirected())
	fmt.Println("IsConnectedUndirected : ", gph.isConnectedUndirected())

}
/*
false
IsConnectedUndirected :  true
*/

/*
* Given a directed graph find if there is a cycle in it.
*/

func (gph *Graph)isCyclePresentDFS(index int, visited[] int , marked[] int ) bool {
	visited[index] = 1
	marked[index] = 1
	head := gph.Edges[index]
	for head != nil {
		dest := head.destination
		if (marked[dest] == 1) {
			return true
		}

		if (visited[dest] == 0) {
			if (gph.isCyclePresentDFS(dest, visited, marked)) {
				return true
			}
		}
		head = head.next
	}
	marked[index] = 0
	return false
}

func (gph *Graph)isCyclePresent() bool {
	count := gph.count
	visited := make([]int, count)
	marked := make([]int, count)
	for index := 0; index < count; index++ {
		if (visited[index] == 0) {
			if (gph.isCyclePresentDFS(index, visited, marked)) {
				return true
			}
		}
	}
	return false
}

func (gph *Graph)isCyclePresentDFSColor(index int, visited[] int) bool {
	visited[index] = 1 // 1 = grey
	var dest int
	head := gph.Edges[index]
	for head != nil {
		dest = head.destination		
		if (visited[dest] == 1) {// "Grey":
			return true
		}

		if (visited[dest] == 0) {// "White":
			if (gph.isCyclePresentDFSColor(dest, visited)) {
				return true
			}
		}
		head = head.next
	}
	visited[index] = 2 // "Black"
	return false
}

func (gph *Graph)isCyclePresentColor() bool {
	count := gph.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		if (visited[i] == 0) {// "White"
			if (gph.isCyclePresentDFSColor(i, visited)) {
				return true
			}
		}
	}
	return false
}

//Testing code
func main9() {
	gph := new(Graph)
	gph.Init(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.AddDirectedEdge(4, 1, 1)
	fmt.Println(gph.isCyclePresent())
	fmt.Println(gph.isCyclePresentColor())
}
/*
true
true
*/

func (gph *Graph)transposeGraph() *Graph {
	count := gph.count
	g := new(Graph)
	g.Init(count)
	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		for head != nil {
			dest := head.destination			
			g.AddDirectedEdge(dest, i, head.cost)
			head = head.next
		}
	}
	return g
}

//Testing code
func main10() {
	gph := new(Graph)
	gph.Init(4)
    gph.AddDirectedEdge(0, 1, 1)
    gph.AddDirectedEdge(0, 2, 1)
    gph.AddDirectedEdge(1, 2, 1)
    gph.AddDirectedEdge(2, 3, 1)
    fmt.Println("Graph is ::")
    gph.Print()
    fmt.Println("Transpose Graph is ::")
    g := gph.transposeGraph()
    g.Print();
}
/*
Graph is ::
Vertex 0 is connected to : 2(cost:1) 1(cost:1) 
Vertex 1 is connected to : 2(cost:1) 
Vertex 2 is connected to : 3(cost:1) 
Vertex 3 is connected to : 
Transpose Graph is ::
Vertex 0 is connected to : 
Vertex 1 is connected to : 0(cost:1) 
Vertex 2 is connected to : 1(cost:1) 0(cost:1) 
Vertex 3 is connected to : 2(cost:1) 
*/
func  (gph *Graph)isConnectedUndirected()bool {
	count := gph.count
	visited := make([]bool, count)

	gph.DFSUtil(0, visited)
	for i := 0; i < count; i++ {
		if (visited[i] == false) {
			return false
		}
	}
	return true
}

	/*
	* Kosaraju Algorithm
	* 
	* Kosaraju’s Algorithm to find strongly connected directed graph based on DFS :
	* 1) Create a visited array of size V, and Initialize all count in visited array as 0. 
	* 2) Choose any vertex and perform a DFS traversal of graph. For all visited count mark them visited as 1. 
	* 3) If DFS traversal does not mark all count as 1, then return 0. 
	* 4) Find transpose or reverse of graph 
	* 5) Repeat step 1, 2 and 3 for the reversed graph. 
	* 6) If DFS traversal mark all the count as 1, then return 1.
	*/
func (gph *Graph)isStronglyConnected() bool {
	count := gph.count
	visited := make([]bool, count)
	gph.DFSUtil(0, visited)
	for i := 0; i < count; i++ {
		if (visited[i] == false) {
			return false
		}
	}

	gReversed := gph.transposeGraph()
	visited = make([]bool, count)
	gReversed.DFSUtil(0, visited)
	for i := 0; i < count; i++ {
		if (visited[i] == false) {
			return false
		}
	}
	return true
}

//Testing code
func main11() {
	gph := new(Graph)
	gph.Init(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 0, 1)
	gph.AddDirectedEdge(2, 4, 1)
	gph.AddDirectedEdge(4, 2, 1)
	fmt.Println(" IsStronglyConnected:: " , gph.isStronglyConnected())
}
/*
 IsStronglyConnected::  true
*/
func (gph *Graph)stronglyConnectedComponent() {
	count := gph.count
	visited := make([]bool, count)

	stk := stack.New()
	for i := 0; i < count; i++ {
		if (visited[i] == false) {
			gph.DFSUtil2(i, visited, stk)
		}
	}
	gReversed := gph.transposeGraph()
	visited = make([]bool, count)


	for (stk.Len() > 0) {
		index := stk.Pop().(int)
		if (visited[index] == false) {
			stk2 := stack.New() 
			gReversed.DFSUtil2(index, visited, stk2)
			stk2.Print()
			fmt.Println("")
		}
	}
}

//Testing code
func main12() {
	gph := new(Graph)
	gph.Init(7)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.AddDirectedEdge(4, 5, 1)
	gph.AddDirectedEdge(5, 3, 1)
	gph.AddDirectedEdge(5, 6, 1)
	gph.stronglyConnectedComponent()
}
/*
0 2 1 
3 5 4 
6 
*/

func (gph *Graph) IsConnected() bool {
	count := gph.count
	visited := make([]bool, count)
	visited[0] = true
	gph.DFSUtil(0, visited)

	for i := 0; i < count; i++ {
		if visited[i] == false {
			return false
		}
	}
	return true
}

func (gph *Graph) Dijkstra(source int) {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)
	que := new(PQueue)
	que.Init()

	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = Infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = source
	que.Add(source, 0)
	
	for que.Len() != 0 {
		curr := que.Remove().(int) // Pop from PQ

	    if (visited[curr] == true) {
            continue
        }
        visited[curr] = true
	
		head := gph.Edges[curr]
		for head != nil {
			alt := head.cost + dist[curr]
			if alt < dist[head.destination] && (visited[head.destination] == false) {
				que.Add(head.destination, alt)
				dist[head.destination] = alt
				previous[head.destination] = curr
			}
			head = head.next
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

func (gph *Graph) Prims() {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)
	que := new(PQueue)
	que.Init()
	source := 0

	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = Infinite
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = source
	que.Add(source, 0)
	
	for que.Len() != 0 {
		curr := que.Remove().(int) // Pop from PQ

	    if (visited[curr] == true) {
            continue
        }
        visited[curr] = true
	
		head := gph.Edges[curr]
		for head != nil {
			alt := head.cost
			if alt < dist[head.destination] && (visited[head.destination] == false) {
				que.Add(head.destination, alt)
				dist[head.destination] = alt
				previous[head.destination] = curr
			}
			head = head.next
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

//Testing code
func main13() {
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
	gph.Dijkstra(1)
	gph.Prims()
}
/*
  edge id  0   prev  1  distance :  4
 edge id  1   prev  1  distance :  0
 edge id  2   prev  1  distance :  8
 edge id  3   prev  2  distance :  15
 edge id  4   prev  5  distance :  22
 edge id  5   prev  2  distance :  12
 edge id  6   prev  7  distance :  12
 edge id  7   prev  1  distance :  11
 edge id  8   prev  2  distance :  10


 edge id  0   prev  -1  distance :  0
 edge id  1   prev  0  distance :  4
 edge id  2   prev  5  distance :  4
 edge id  3   prev  2  distance :  7
 edge id  4   prev  3  distance :  9
 edge id  5   prev  6  distance :  2
 edge id  6   prev  7  distance :  1
 edge id  7   prev  0  distance :  8
 edge id  8   prev  2  distance :  2

 */


func (gph *Graph) ShortestPath(source int) {
	var curr int
	count := gph.count
	distance := make([]int, count)
	path := make([]int, count)
	que := queue.New()
	for i := 0; i < count; i++ {
		distance[i] = -1
	}
	que.Enqueue(source)
	distance[source] = 0
	path[source] = source
	for que.Len() != 0 {
		curr = que.Dequeue().(int)
		head := gph.Edges[curr]
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

func (gph *Graph) BellmanFordShortestPath(source int) {
	count := gph.count
	distance := make([]int, count)
	path := make([]int, count)

	for i := 0; i < count; i++ {
		distance[i] = Infinite
	}
	distance[source] = 0
	path[source] = source
	for i := 0; i < count-1; i++ {
		for j := 0; j < count; j++ {
			head := gph.Edges[j]
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

//Testing code
func main14() {
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
	gph.ShortestPath(0)
}
/*
0  to  0  weight  0
0  to  1  weight  1
1  to  2  weight  2
2  to  3  weight  3
5  to  4  weight  4
6  to  5  weight  3
7  to  6  weight  2
0  to  7  weight  1
7  to  8  weight  2
*/

//Testing code
func main15() {
	gph := new(Graph)
	gph.Init(5)
    gph.AddDirectedEdge(0, 1, 3)
    gph.AddDirectedEdge(0, 4, 2)
    gph.AddDirectedEdge(1, 2, 1)
    gph.AddDirectedEdge(2, 3, 1)
    gph.AddDirectedEdge(4, 1, -2)
    gph.AddDirectedEdge(4, 3, 1)
    gph.BellmanFordShortestPath(0)
}
/*
0  to  0  weight  0
4  to  1  weight  0
1  to  2  weight  1
2  to  3  weight  2
0  to  4  weight  2
*/

func heightTreeParentArr(arr[] int) int{
	count := len(arr)
	heightArr := make([]int, count)
	gph := new(Graph)
	gph.Init(count)
	source := 0
	for i := 0; i < count; i++ {
		if (arr[i] != -1) {
			gph.AddDirectedEdge(arr[i], i, 1)
		} else {
			source = i
		}
	}
	visited := make([]bool, count)
	visited[source] = true
	que := queue.New()
	que.Enqueue(source)
	heightArr[source] = 0
	maxHight := 0
	for (que.Len() > 0) {
		curr := que.Dequeue().(int)
		height := heightArr[curr]
		if (height > maxHight) {
			maxHight = height
		}
		head := gph.Edges[curr]
		for head != nil {
			if (visited[head.destination] == false) {
				visited[head.destination] = true
				que.Enqueue(head.destination)
				heightArr[head.destination] = height + 1
			}
			head = head.next
		}
	}
	return maxHight
}

func getHeight(arr[] int, height[] int, index int) int {
	if (arr[index] == -1) {
		return 0
	} else {
		return getHeight(arr, height, arr[index]) + 1
	}
}

func heightTreeParentArr2(arr[] int) int {
	count := len(arr)
	height := make ([]int, count)
	maxHeight := -1
	for i := 0; i < count; i++ {
		height[i] = getHeight(arr, height, i)
		if (maxHeight <  height[i]) { 
			maxHeight = height[i]
		}
	}
	return maxHeight
}

func main16() {
	parentArray := [] int{ -1, 0, 1, 2, 3 }
	fmt.Println(heightTreeParentArr(parentArray))
	fmt.Println(heightTreeParentArr2(parentArray))
}

func (gph *Graph)isConnected() bool{
	count := gph.count
	visited := make([]bool, count)

	// Find a vertex with non - zero degree
	// DFS traversal of graph from a vertex with non - zero degree
	
	for i := 0; i < count; i++ {
		if (gph.Edges[i] != nil) {
			gph.DFSUtil(i, visited)
			break
		}
	}
	// Check if all non - zero degree count are visited
	for i := 0; i < count; i++ {
		if (gph.Edges[i] != nil) {
			if (visited[i] == false) {
				return false
			}
		}
	}
	return true
}

/*
	* The function returns one of the following values Return 0 if graph is not
	* Eulerian Return 1 if graph has an Euler path (Semi-Eulerian) Return 2 if
	* graph has an Euler Circuit (Eulerian)
	*/
func (gph *Graph)isEulerian() int {
	count := gph.count
	var odd int
	// Check if all non - zero degree nodes are connected
	if (gph.isConnected() == false) {
		fmt.Println("graph is not Eulerian")
		return 0
	} else {
		// Count odd degree
		odd = 0
		inDegree := make([]int, count) 
		outDegree := make([]int, count)
			
		for i := 0; i < count; i++ {
			head := gph.Edges[i]
			for head != nil {
				outDegree[i] += 1
				inDegree[head.destination] += 1
				head = head.next
			}
		}
		for i := 0; i < count; i++ {
			if ((inDegree[i] + outDegree[i]) % 2 != 0) {
				odd += 1
			}
		}
	}

	if (odd == 0) {
		fmt.Println("graph is Eulerian")
		return 2
	} else if (odd == 2) {
		fmt.Println("graph is Semi-Eulerian")
		return 1
	} else {
		fmt.Println("graph is not Eulerian")
		return 0
	}
}


func (gph *Graph)isStronglyConnected2() bool {
	count := gph.count
	visited := make([]bool, count)
	// Find a vertex with non - zero degree
	var index int
	for index = 0; index < count; index++ {
		if (gph.Edges[index] != nil) {
			break
		}
	}
	// DFS traversal of graph from a vertex with non - zero degree
	gph.DFSUtil(index, visited)
	for i := 0; i < count; i++ {
		if (visited[i] == false && gph.Edges[i] != nil) {
			return false
		}
	}

	gReversed := gph.transposeGraph()
	for i := 0; i < count; i++ {
		visited[i] = false
	}
	gReversed.DFSUtil(index, visited)

	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		if (visited[i] == false && head != nil) {
			return false
		}
	}
	return true
}

func (gph *Graph)isEulerianCycle() bool {
	// Check if all non - zero degree count are connected
	count := gph.count
	inDegree := make([]int, count)  
	outDegree := make([]int, count)
	if (!gph.isStronglyConnected2()) {
		return false
	}

	// Check if in degree and out degree of every vertex is same
	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		for head != nil {
			outDegree[i] += 1
			inDegree[head.destination] += 1
			head = head.next
		}
	}
	for i := 0; i < count; i++ {
		if (inDegree[i] != outDegree[i]) {
			return false
		}
	}
	return true
}

//Testing code
func main17() {
	gph := new(Graph)
	gph.Init(5)
	gph.AddDirectedEdge(1, 0, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 1, 1)
	gph.AddDirectedEdge(0, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.isEulerian()
}
/*
graph is Semi-Eulerian
*/

//Testing code
func main18() {
	gph := new(Graph)
	gph.Init(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(0, 4, 1)
	gph.AddDirectedEdge(4, 3, 1)
	gph.AddDirectedEdge(3, 0, 1)
	fmt.Println(gph.isEulerianCycle())
}
/*
true
*/
func main(){
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
	main9()
	main10()
	main11()
	main12()
	main13()
	main14()
	main15()
	main16()
	main17()
	main18()
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