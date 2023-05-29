package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

type Edge struct {
	src  int
	dest int
	cost int
	next *Edge
}

type Graph struct {
	count int
	Edges []*Edge
}

func NewGraph(count int) *Graph {
	gph := &Graph{
		count: count,
		Edges: make([]*Edge, count),
	}
	return gph
}

func (gph *Graph) AddDirectedEdge(source, destination, cost int) {
	edge := &Edge{src: source, dest: destination, cost: cost, next: gph.Edges[source]}
	gph.Edges[source] = edge
}

func (gph *Graph) AddDirectedEdgeUnweighted(source, destination int) {
	gph.AddDirectedEdge(source, destination, 1)
}

func (gph *Graph) AddUndirectedEdge(source, destination, cost int) {
	gph.AddDirectedEdge(source, destination, cost)
	gph.AddDirectedEdge(destination, source, cost)
}

func (gph *Graph) AddUndirectedEdgeUnweighted(source, destination int) {
	gph.AddUndirectedEdge(source, destination, 1)
}

func (gph *Graph) Print() {
	for i := 0; i < gph.count; i++ {
		ad := gph.Edges[i]
		fmt.Print("Vertex ", i, " is connected to: ")
		for ad != nil {
			fmt.Print(ad.dest, "(cost:", ad.cost, ") ")
			ad = ad.next
		}
		fmt.Println()
	}
}

//Testing code
func main1() {
	gph := NewGraph(4)
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
	stk := new(Stack)
	path := []int{}
	visited[source] = true
	stk.Push(source)

	for stk.Len() != 0 {
		curr = stk.Pop().(int)
		path = append(path, curr)
		head := gph.Edges[curr]
		for head != nil {
			if !visited[head.dest] {
				visited[head.dest] = true
				stk.Push(head.dest)
			}
			head = head.next
		}
	}
	fmt.Println("DFS Path is: ", path)
	return visited[target]
}

func (gph *Graph) DFS(source, target int) bool {
	count := gph.count
	visited := make([]bool, count)

	fmt.Print("DFS Path is: ")
	gph.dfsUtil(source, visited)
	fmt.Println()
	return visited[target]
}

func (gph *Graph) dfsUtil(curr int, visited []bool) {
	visited[curr] = true
	head := gph.Edges[curr]
	for head != nil {
		if !visited[head.dest] {
			gph.dfsUtil(head.dest, visited)
		}
		head = head.next
	}
}

func (gph *Graph) dfsUtil2(curr int, visited []bool, stk *Stack) {
	visited[curr] = true
	head := gph.Edges[curr]

	for head != nil {
		if !visited[head.dest] {
			gph.dfsUtil2(head.dest, visited, stk)
		}
		head = head.next
	}
	stk.Push(curr)
}

func (gph *Graph) BFS(source int, target int) bool {
	count := gph.count
	visited := make([]bool, count)
	que := new(Queue)
	path := []int{}
	visited[source] = true
	que.Add(source)

	for que.Len() != 0 {
		curr := que.Remove().(int)
		path = append(path, curr)

		head := gph.Edges[curr]
		for head != nil {
			if !visited[head.dest] {
				visited[head.dest] = true
				que.Add(head.dest)
			}
			head = head.next
		}
	}
	fmt.Println("BFS Path is:", path)
	return visited[target]
}

// Testing code
func main2() {
	gph := NewGraph(8)
	gph.AddUndirectedEdge(0, 3, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(0, 1, 1)

	gph.AddUndirectedEdge(1, 4, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	gph.AddUndirectedEdge(3, 6, 1)

	gph.AddUndirectedEdge(6, 7, 1)
	gph.AddUndirectedEdge(5, 7, 1)
	gph.AddUndirectedEdge(4, 7, 1)

	fmt.Println("Path between 0 & 6:", gph.DFSStack(0, 6))
	fmt.Println("Path between 0 & 6:", gph.DFS(0, 6))
	fmt.Println("Path between 0 & 6:", gph.BFS(0, 6))
}

/*
DFS Path is: [0 3 6 7 5 4 2 1]
Path between 0 & 6: true
DFS Path is: 0 1 4 7 5 2 6 3
Path between 0 & 6: true
BFS Path is: [0 1 2 3 4 5 6 7]
Path between 0 & 6: true
*/

func (gph *Graph) TopologicalSort() {
	fmt.Print("Topological order of given graph is: ")
	var count = gph.count
	stk := new(Stack)
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

func (gph *Graph) TopologicalSortDFS(index int, visited []bool, stk *Stack) {
	head := gph.Edges[index]
	for head != nil {
		if !visited[head.dest] {
			visited[head.dest] = true
			gph.TopologicalSortDFS(head.dest, visited, stk)
		}
		head = head.next
	}
	stk.Push(index)
}

// Testing code
func main3() {
	gph := NewGraph(9)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(1, 4, 1)
	gph.AddDirectedEdge(3, 2, 1)
	gph.AddDirectedEdge(3, 5, 1)
	gph.AddDirectedEdge(4, 5, 1)
	gph.AddDirectedEdge(4, 6, 1)
	gph.AddDirectedEdge(5, 7, 1)
	gph.AddDirectedEdge(6, 7, 1)
	gph.AddDirectedEdge(7, 8, 1)
	gph.TopologicalSort()
}

/*
Topological order of given graph is: 1 3 4 5 6 7 8 0 2
*/

func (gph *Graph) PathExist(source int, destination int) bool {
	count := gph.count
	visited := make([]bool, count)
	visited[source] = true
	gph.dfsUtil(source, visited)
	return visited[destination]
}

func (gph *Graph) countAllPathDFS(visited []bool, source int, dest int) int {
	if source == dest {
		return 1
	}

	count := 0
	visited[source] = true
	head := gph.Edges[source]

	for head != nil {
		if !visited[head.dest] {
			count += gph.countAllPathDFS(visited, head.dest, dest)
		}
		visited[head.dest] = false
		head = head.next
	}
	return count
}

func (gph *Graph) countAllPath(src int, dest int) int {
	visited := make([]bool, gph.count)
	return gph.countAllPathDFS(visited, src, dest)
}

func (gph *Graph) printAllPath(src int, dest int) {
	visited := make([]bool, gph.count)
	path := new(Stack)
	gph.printAllPathDFS(src, dest, visited, path)
}

func (gph *Graph) printAllPathDFS(source int, dest int, visited []bool, path *Stack) {
	path.Push(source)
	if source == dest {
		path.Print()
		path.Pop()
		return
	}
	visited[source] = true
	head := gph.Edges[source]
	for head != nil {
		if !visited[head.dest] {
			gph.printAllPathDFS(head.dest, dest, visited, path)
		}
		head = head.next
	}
	visited[source] = false
	path.Pop()
}

// Testing code
func main4() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.AddDirectedEdge(1, 4, 1)
	fmt.Println("PathExist:", gph.PathExist(0, 4))
	fmt.Println("Path Count:", gph.countAllPath(0, 4))
	gph.printAllPath(0, 4)
	fmt.Println()
}

/*
PathExist: true
Path Count: 3
4 3 2 0
4 1 0
4 3 1 0
*/

func (gph *Graph) rootVertex() int {
	visited := make([]bool, gph.count)
	retVal := -1

	for i := 0; i < gph.count; i++ {
		if !visited[i] {
			gph.dfsUtil(i, visited)
			retVal = i
		}
	}
	fmt.Println("Root vertex is:", retVal)
	return retVal
}

// Testing code
func main5() {
	gph := NewGraph(7)
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
Root vertex is: 5
*/

func (gph *Graph) transitiveClosureUtil(source int, dest int, tc *[10][10]int) {
	tc[source][dest] = 1
	head := gph.Edges[dest]
	for head != nil {
		if tc[source][head.dest] == 0 {
			gph.transitiveClosureUtil(source, head.dest, tc)
		}
		head = head.next
	}
}

func (gph *Graph) transitiveClosure() [10][10]int {
	tc := [10][10]int{}
	for i := 0; i < gph.count; i++ {
		gph.transitiveClosureUtil(i, i, &tc)
	}
	return tc
}

// Testing code
func main6() {
	gph := NewGraph(4)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 3, 1)
	tc := gph.transitiveClosure()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(tc[i][j], " ")
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

func (gph *Graph) BfsLevelNode(source int) {
	count := gph.count
	visited := make([]bool, count)
	level := make([]int, count)
	visited[source] = true
	que := new(Queue)
	que.Add(source)
	level[source] = 0
	fmt.Println("\nNode  - Level")
	for que.Len() > 0 {
		curr := que.Remove().(int)
		depth := level[curr]
		head := gph.Edges[curr]
		fmt.Println(curr, " - ", depth)
		for head != nil {
			if !visited[head.dest] {
				visited[head.dest] = true
				que.Add(head.dest)
				level[head.dest] = depth + 1
			}
			head = head.next
		}
	}
}

func (gph *Graph) BfsDistance(source int, dest int) int {
	count := gph.count
	visited := make([]bool, count)
	que := new(Queue)
	que.Add(source)
	visited[source] = true
	level := make([]int, count)
	level[source] = 0
	for que.Len() > 0 {
		curr := que.Remove().(int)
		depth := level[curr]
		head := gph.Edges[curr]
		for head != nil {
			if head.dest == dest {
				return depth + 1
			}
			if !visited[head.dest] {
				visited[head.dest] = true
				que.Add(head.dest)
				level[head.dest] = depth + 1
			}
			head = head.next
		}
	}
	return -1
}

//Testing code
func main7() {
	gph := NewGraph(7)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(0, 2, 1)
	gph.AddUndirectedEdge(0, 4, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	gph.AddUndirectedEdge(3, 4, 1)
	gph.AddUndirectedEdge(4, 5, 1)
	gph.AddUndirectedEdge(4, 6, 1)
	gph.BfsLevelNode(1)
	fmt.Println("BfsDistance :", gph.BfsDistance(1, 6))
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
BfsDistance : 3
*/

func (gph *Graph) isCyclePresentUndirectedDFS(index int, parentIndex int, visited []bool) bool {
	visited[index] = true
	var dest int
	head := gph.Edges[index]
	for head != nil {
		dest = head.dest
		if !visited[dest] {
			if gph.isCyclePresentUndirectedDFS(dest, index, visited) {
				return true
			}
		} else if parentIndex != dest {
			return true
		}
		head = head.next
	}
	return false
}

func (gph *Graph) isCyclePresentUndirected() bool {
	count := gph.count
	visited := make([]bool, count)
	for i := 0; i < count; i++ {
		if visited[i] == false {
			if gph.isCyclePresentUndirectedDFS(i, -1, visited) {
				return true
			}
		}
	}
	return false
}

func (gph *Graph) find(parent []int, index int) int {
	var p int = parent[index]
	for p != -1 {
		index = p
		p = parent[index]
	}
	return index
}

func (gph *Graph) union(parent []int, x int, y int) {
	parent[y] = x
}

func (gph *Graph) isCyclePresentUndirected2() bool {
	count := gph.count
	parent := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = -1
	}
	var edge []*Edge
	flags := make([][]bool, count)
	for i := 0; i < count; i++ {
		flags[i] = make([]bool, count)
	}
	for i := 0; i < count; i++ {
		ad := gph.Edges[i]
		for ad != nil {
			// Using flags[][] array, if considered edge x to y,
			// then ignore edge y to x.
			if flags[ad.dest][ad.src] == false {
				edge = append(edge, ad)
				flags[ad.src][ad.dest] = true
			}
			ad = ad.next
		}
	}
	for _, e := range edge {
		var x int = gph.find(parent, e.src)
		var y int = gph.find(parent, e.dest)
		if x == y {
			return true
		}
		gph.union(parent, x, y)
	}
	return false
}

type Sets struct {
	parent int
	rank   int
}

func NewSet(p int, r int) (st *Sets) {
	st = new(Sets)
	st.parent = p
	st.rank = r
	return st
}

// root element of set
func find(sets []Sets, index int) int {
	p := sets[index].parent
	for p != index {
		index = p
		p = sets[index].parent
	}
	return index
}

// consider x and y are roots of sets.
func union(sets []Sets, x int, y int) {
	if sets[x].rank < sets[y].rank {
		sets[x].parent = y
	} else if sets[y].rank < sets[x].rank {
		sets[y].parent = x
	} else {
		sets[x].parent = y
		sets[y].rank++
	}
}

func (gph *Graph) isCyclePresentUndirected3() bool {
	count := gph.count
	//Different subsets are created.
	sets := make([]Sets, count)
	for i := 0; i < count; i++ {
		sets[i].parent = i
		sets[i].rank = 0
	}
	var edge []*Edge
	flags := make([][]bool, count)
	for i := 0; i < count; i++ {
		flags[i] = make([]bool, count)
	}
	for i := 0; i < count; i++ {
		ad := gph.Edges[i]
		for ad != nil {
			// Using flags[][] array, if considered edge x to y,
			// then ignore edge y to x.
			if flags[ad.dest][ad.src] == false {
				edge = append(edge, ad)
				flags[ad.src][ad.dest] = true
			}
			ad = ad.next
		}
	}
	for _, e := range edge {
		var x int = find(sets, e.src)
		var y int = find(sets, e.dest)
		if x == y {
			return true
		}
		union(sets, x, y)
	}
	return false
}

// Testing code
func main8() {
	gph := NewGraph(6)
	gph.AddUndirectedEdge(0, 1, 1)
	gph.AddUndirectedEdge(1, 2, 1)
	gph.AddUndirectedEdge(3, 4, 1)
	gph.AddUndirectedEdge(4, 2, 1)
	gph.AddUndirectedEdge(2, 5, 1)
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected())
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected2())
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected3())
	fmt.Println("IsConnectedUndirected:", gph.IsConnectedUndirected())
	gph.AddUndirectedEdge(4, 1, 1)
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected())
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected2())
	fmt.Println("Cycle Present:", gph.isCyclePresentUndirected3())
	fmt.Println("IsConnectedUndirected:", gph.IsConnectedUndirected())
}

/*
Cycle Present: false
Cycle Present: false
Cycle Present: false
IsConnectedUndirected: true
Cycle Present: true
Cycle Present: true
Cycle Present: true
IsConnectedUndirected: true
*/

/*
* Given a directed graph find if there is a cycle in it.
 */

func (gph *Graph) isCyclePresentDFS(index int, visited []int, marked []int) bool {
	visited[index] = 1
	marked[index] = 1
	head := gph.Edges[index]
	for head != nil {
		dest := head.dest
		if marked[dest] == 1 {
			return true
		}

		if visited[dest] == 0 {
			if gph.isCyclePresentDFS(dest, visited, marked) {
				return true
			}
		}
		head = head.next
	}
	marked[index] = 0
	return false
}

func (gph *Graph) isCyclePresent() bool {
	count := gph.count
	visited := make([]int, count)
	marked := make([]int, count)
	for index := 0; index < count; index++ {
		if visited[index] == 0 {
			if gph.isCyclePresentDFS(index, visited, marked) {
				return true
			}
		}
	}
	return false
}

func (gph *Graph) isCyclePresentDFSColor(index int, visited []int) bool {
	visited[index] = 1 // 1 = grey
	var dest int
	head := gph.Edges[index]
	for head != nil {
		dest = head.dest
		if visited[dest] == 1 { // "Grey":
			return true
		}

		if visited[dest] == 0 { // "White":
			if gph.isCyclePresentDFSColor(dest, visited) {
				return true
			}
		}
		head = head.next
	}
	visited[index] = 2 // "Black"
	return false
}

func (gph *Graph) isCyclePresentColor() bool {
	count := gph.count
	visited := make([]int, count)
	for i := 0; i < count; i++ {
		if visited[i] == 0 { // "White"
			if gph.isCyclePresentDFSColor(i, visited) {
				return true
			}
		}
	}
	return false
}

//Testing code
func main9() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(1, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	fmt.Println("isCyclePresent :", gph.isCyclePresent())
	fmt.Println("isCyclePresent :", gph.isCyclePresentColor())
	gph.AddDirectedEdge(4, 1, 1)
	fmt.Println("isCyclePresent :", gph.isCyclePresent())
	fmt.Println("isCyclePresent :", gph.isCyclePresentColor())
}

/*
 isCyclePresent : false
 isCyclePresent : false
 isCyclePresent : true
 isCyclePresent : true
*/

func (gph *Graph) transposeGraph() *Graph {
	count := gph.count
	g := NewGraph(count)
	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		for head != nil {
			dest := head.dest
			g.AddDirectedEdge(dest, i, head.cost)
			head = head.next
		}
	}
	return g
}

//Testing code
func main10() {
	gph := NewGraph(4)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	fmt.Println("Graph is ::")
	gph.Print()
	fmt.Println("Transpose Graph is ::")
	g := gph.transposeGraph()
	g.Print()
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

func (gph *Graph) IsConnectedUndirected() bool {
	count := gph.count
	visited := make([]bool, count)

	gph.dfsUtil(0, visited)
	for i := 0; i < count; i++ {
		if visited[i] == false {
			return false
		}
	}
	return true
}

func (gph *Graph) IsStronglyConnected() bool {
	count := gph.count
	visited := make([]bool, count)
	gph.dfsUtil(0, visited)
	for i := 0; i < count; i++ {
		if visited[i] == false {
			return false
		}
	}

	gReversed := gph.transposeGraph()
	visited = make([]bool, count)
	gReversed.dfsUtil(0, visited)
	for i := 0; i < count; i++ {
		if visited[i] == false {
			return false
		}
	}
	return true
}

//Testing code
func main11() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 0, 1)
	gph.AddDirectedEdge(2, 4, 1)
	gph.AddDirectedEdge(4, 2, 1)
	fmt.Println("IsStronglyConnected :", gph.IsStronglyConnected())
}

/*
 IsStronglyConnected : true
*/

func (gph *Graph) StronglyConnectedComponent() {
	count := gph.count
	visited := make([]bool, count)

	stk := new(Stack)
	for i := 0; i < count; i++ {
		if visited[i] == false {
			gph.dfsUtil2(i, visited, stk)
		}
	}
	gReversed := gph.transposeGraph()
	visited = make([]bool, count)

	for stk.Len() > 0 {
		index := stk.Pop().(int)
		if visited[index] == false {
			stk2 := new(Stack)
			gReversed.dfsUtil2(index, visited, stk2)
			stk2.Print()
		}
	}
}

//Testing code
func main12() {
	gph := NewGraph(7)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.AddDirectedEdge(4, 5, 1)
	gph.AddDirectedEdge(5, 3, 1)
	gph.AddDirectedEdge(5, 6, 1)
	gph.StronglyConnectedComponent()
}

/*
 [1 2 0]
 [4 5 3]
 [6]
*/

func (gph *Graph) IsConnected() bool {
	count := gph.count
	visited := make([]bool, count)
	visited[0] = true
	gph.dfsUtil(0, visited)

	for i := 0; i < count; i++ {
		if visited[i] == false {
			return false
		}
	}
	return true
}

func heightTreeParentArr(arr []int) int {
	count := len(arr)
	heightArr := make([]int, count)
	gph := NewGraph(count)
	source := 0
	for i := 0; i < count; i++ {
		if arr[i] != -1 {
			gph.AddDirectedEdge(arr[i], i, 1)
		} else {
			source = i
		}
	}
	visited := make([]bool, count)
	visited[source] = true
	que := new(Queue)
	que.Add(source)
	heightArr[source] = 0
	maxHeight := 0
	for que.Len() > 0 {
		curr := que.Remove().(int)
		height := heightArr[curr]
		if height > maxHeight {
			maxHeight = height
		}
		head := gph.Edges[curr]
		for head != nil {
			if !visited[head.dest] {
				visited[head.dest] = true
				que.Add(head.dest)
				heightArr[head.dest] = height + 1
			}
			head = head.next
		}
	}
	return maxHeight
}

func getHeight(arr []int, height []int, index int) int {
	if arr[index] == -1 {
		return 0
	} else {
		return getHeight(arr, height, arr[index]) + 1
	}
}

func heightTreeParentArr2(arr []int) int {
	count := len(arr)
	height := make([]int, count)
	maxHeight := -1
	for i := 0; i < count; i++ {
		height[i] = getHeight(arr, height, i)
		if maxHeight < height[i] {
			maxHeight = height[i]
		}
	}
	return maxHeight
}

func main13() {
	parentArray := []int{-1, 0, 1, 2, 3}
	fmt.Println(heightTreeParentArr(parentArray))
	fmt.Println(heightTreeParentArr2(parentArray))
}

/*
4
4
*/

func (gph *Graph) IsConnected2() bool {
	count := gph.count
	visited := make([]bool, count)

	// Find a vertex with non-zero degree
	// DFS traversal of graph from a vertex with non-zero degree

	for i := 0; i < count; i++ {
		if gph.Edges[i] != nil {
			gph.dfsUtil(i, visited)
			break
		}
	}
	// Check if all non-zero degree count are visited
	for i := 0; i < count; i++ {
		if gph.Edges[i] != nil {
			if visited[i] == false {
				return false
			}
		}
	}
	return true
}

func (gph *Graph) IsEulerian() int {
	count := gph.count
	var odd int
	// Check if all non-zero degree nodes are connected
	if gph.IsConnected() == false {
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
				inDegree[head.dest] += 1
				head = head.next
			}
		}
		for i := 0; i < count; i++ {
			if (inDegree[i]+outDegree[i])%2 != 0 {
				odd += 1
			}
		}
	}

	if odd == 0 {
		fmt.Println("graph is Eulerian")
		return 2
	} else if odd == 2 {
		fmt.Println("graph is Semi-Eulerian")
		return 1
	} else {
		fmt.Println("graph is not Eulerian")
		return 0
	}
}

func (gph *Graph) IsStronglyConnected2() bool {
	count := gph.count
	visited := make([]bool, count)
	// Find a vertex with non-zero degree
	var index int
	for index = 0; index < count; index++ {
		if gph.Edges[index] != nil {
			break
		}
	}
	// DFS traversal of graph from a vertex with non-zero degree
	gph.dfsUtil(index, visited)
	for i := 0; i < count; i++ {
		if visited[i] == false && gph.Edges[i] != nil {
			return false
		}
	}

	gReversed := gph.transposeGraph()
	for i := 0; i < count; i++ {
		visited[i] = false
	}
	gReversed.dfsUtil(index, visited)

	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		if visited[i] == false && head != nil {
			return false
		}
	}
	return true
}

func (gph *Graph) IsEulerianCycle() bool {
	// Check if all non-zero degree count are connected
	count := gph.count
	inDegree := make([]int, count)
	outDegree := make([]int, count)
	if !gph.IsStronglyConnected2() {
		return false
	}

	// Check if in degree and out degree of every vertex is same
	for i := 0; i < count; i++ {
		head := gph.Edges[i]
		for head != nil {
			outDegree[i] += 1
			inDegree[head.dest] += 1
			head = head.next
		}
	}
	for i := 0; i < count; i++ {
		if inDegree[i] != outDegree[i] {
			return false
		}
	}
	return true
}

//Testing code
func main14() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(1, 0, 1)
	gph.AddDirectedEdge(0, 2, 1)
	gph.AddDirectedEdge(2, 1, 1)
	gph.AddDirectedEdge(0, 3, 1)
	gph.AddDirectedEdge(3, 4, 1)
	gph.IsEulerian()
	gph.AddDirectedEdge(4, 0, 1)
	gph.IsEulerian()
}

/*
graph is Semi-Eulerian
graph is Eulerian
*/

//Testing code
func main15() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(0, 1, 1)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 0, 1)
	gph.AddDirectedEdge(0, 4, 1)
	gph.AddDirectedEdge(4, 3, 1)
	gph.AddDirectedEdge(3, 0, 1)
	fmt.Println(gph.IsEulerianCycle())
}

/*
true
*/

func (gph *Graph) ShortestPath(source int) {
	var curr int
	count := gph.count
	distance := make([]int, count)
	path := make([]int, count)
	que := new(Queue)
	for i := 0; i < count; i++ {
		distance[i] = -1
	}
	que.Add(source)
	distance[source] = 0
	path[source] = source
	for que.Len() != 0 {
		curr = que.Remove().(int)
		head := gph.Edges[curr]
		for head != nil {
			if distance[head.dest] == -1 {
				distance[head.dest] = distance[curr] + 1
				path[head.dest] = curr
				que.Add(head.dest)
			}
			head = head.next
		}
	}

	gph.printPath(path, distance, count, source)
}

func (gph *Graph) printPathUtil(previous []int, source int, dest int) string {
	var path = ""
	if dest == source {
		path += strconv.Itoa(source)
	} else {
		path += gph.printPathUtil(previous, source, previous[dest])
		path += ("->" + strconv.Itoa(dest))
	}
	return path
}

func (gph *Graph) printPath(previous []int, dist []int, count int, source int) {
	output := "Shortest Paths: "
	for i := 0; i < count; i++ {
		if dist[i] == 99999 {
			output += ("(" + strconv.Itoa(source) + "->" + strconv.Itoa(i) + " @ Unreachable) ")
		} else if i != previous[i] {
			output += "("
			output += gph.printPathUtil(previous, source, i)
			output += (" @ " + strconv.Itoa(dist[i]) + ") ")
		}
	}
	fmt.Println(output)
}

//Testing code
func main16() {
	gph := NewGraph(9)
	gph.AddUndirectedEdgeUnweighted(0, 1)
	gph.AddUndirectedEdgeUnweighted(0, 7)
	gph.AddUndirectedEdgeUnweighted(1, 2)
	gph.AddUndirectedEdgeUnweighted(1, 7)
	gph.AddUndirectedEdgeUnweighted(2, 3)
	gph.AddUndirectedEdgeUnweighted(2, 8)
	gph.AddUndirectedEdgeUnweighted(2, 5)
	gph.AddUndirectedEdgeUnweighted(3, 4)
	gph.AddUndirectedEdgeUnweighted(3, 5)
	gph.AddUndirectedEdgeUnweighted(4, 5)
	gph.AddUndirectedEdgeUnweighted(5, 6)
	gph.AddUndirectedEdgeUnweighted(6, 7)
	gph.AddUndirectedEdgeUnweighted(6, 8)
	gph.AddUndirectedEdgeUnweighted(7, 8)
	gph.ShortestPath(0)
}

/*
Shortest Paths: (0->1 @ 1) (0->1->2 @ 2) (0->1->2->3 @ 3) (0->7->6->5->4 @ 4) (0->7->6->5 @ 3) (0->7->6 @ 2) (0->7 @ 1) (0->7->8 @ 2)
*/

func (gph *Graph) Dijkstra(source int) {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)

	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = math.MaxInt32
		visited[i] = false
	}

	dist[source] = 0
	previous[source] = source

	type Node struct {
		dest int
		cost int
	}
	cmp := func(a, b interface{}) bool { // Greater function
		return a.(Node).cost > b.(Node).cost
	}
	hp := CreateHeap(cmp)
	hp.Add(Node{source, 0})

	for hp.Size() != 0 {
		curr := hp.Remove().(Node).dest // Pop from PQ
		if visited[curr] {
			continue
		}
		visited[curr] = true

		head := gph.Edges[curr]
		for head != nil {
			alt := head.cost + dist[curr]
			if alt < dist[head.dest] && !visited[head.dest] {
				hp.Add(Node{head.dest, alt})
				dist[head.dest] = alt
				previous[head.dest] = curr
			}
			head = head.next
		}
	}
	// Printing result.
	gph.printPath(previous, dist, count, source)
}

//Testing code
func main17() {
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
	gph.Dijkstra(0)
}

/*
Shortest Paths: (0->1 @ 4) (0->1->2 @ 12) (0->1->2->3 @ 19) (0->7->6->5->4 @ 21) (0->7->6->5 @ 11) (0->7->6 @ 9) (0->7 @ 8) (0->1->2->8 @ 14)
*/

func (gph *Graph) BellmanFordShortestPath(source int) {
	count := gph.count
	distance := make([]int, count)
	path := make([]int, count)

	for i := 0; i < count; i++ {
		distance[i] = math.MaxInt32
	}
	distance[source] = 0
	path[source] = source
	for i := 0; i < count-1; i++ {
		for j := 0; j < count; j++ {
			head := gph.Edges[j]
			for head != nil {
				newDistance := distance[j] + head.cost
				if distance[head.dest] > newDistance {
					distance[head.dest] = newDistance
					path[head.dest] = j
				}
				head = head.next
			}
		}
	}
	gph.printPath(path, distance, count, source)
}

//Testing code
func main18() {
	gph := NewGraph(5)
	gph.AddDirectedEdge(0, 1, 3)
	gph.AddDirectedEdge(0, 4, 2)
	gph.AddDirectedEdge(1, 2, 1)
	gph.AddDirectedEdge(2, 3, 1)
	gph.AddDirectedEdge(4, 1, -2)
	gph.AddDirectedEdge(4, 3, 1)
	gph.BellmanFordShortestPath(0)
}

/*
Shortest Paths: (0->4->1 @ 0) (0->4->1->2 @ 1) (0->4->1->2->3 @ 2) (0->4 @ 2)

*/

func (gph *Graph) PrimsMST() {
	count := gph.count
	previous := make([]int, count)
	dist := make([]int, count)
	visited := make([]bool, count)
	for i := 0; i < count; i++ {
		previous[i] = -1
		dist[i] = math.MaxInt32
		visited[i] = false
	}

	source := 0
	dist[source] = 0
	previous[source] = source

	type Edge struct {
		dest int
		cost int
	}

	cmp := func(a, b interface{}) bool { // Greater function
		return a.(Edge).cost > b.(Edge).cost
	}
	hp := CreateHeap(cmp)
	hp.Add(Edge{source, 0})

	for hp.Size() != 0 {
		curr := hp.Remove().(Edge).dest // Pop from PQ

		if visited[curr] {
			continue
		}
		visited[curr] = true

		head := gph.Edges[curr]
		for head != nil {
			alt := head.cost
			if alt < dist[head.dest] && (!visited[head.dest]) {
				hp.Add(Edge{head.dest, alt})
				dist[head.dest] = alt
				previous[head.dest] = curr
			}
			head = head.next
		}
	}
	// Printing result.
	total := 0
	output := "Edges are : "
	for i := 0; i < count; i++ {
		if dist[i] == math.MaxInt32 {
			output += "(" + strconv.Itoa(i) + ",  Unreachable)"
		} else if previous[i] != i {
			output += "(" + strconv.Itoa(previous[i]) + "->" + strconv.Itoa(i) + " @ " + strconv.Itoa(dist[i]) + ") "
			total += dist[i]
		}
	}
	fmt.Println(output)
	fmt.Println("Total MST cost :", total)
}

func (gph *Graph) kruskalMST() {
	count := gph.count
	//Different subsets are created.
	sets := make([]Sets, count)
	for i := 0; i < count; i++ {
		sets[i].parent = i
		sets[i].rank = 0
	}

	// Edges are added to array and sorted.
	E := 0
	edge := make([]*Edge, 100)
	for i := 0; i < count; i++ {
		ad := gph.Edges[i]
		for ad != nil {
			edge[E] = ad
			E++
			ad = ad.next
		}
	}

	sort.Slice(edge[0:E], func(i, j int) bool {
		return edge[i].cost < edge[j].cost
	})

	sum := 0
	output := "Edges are : "
	for i := 0; i < E; i++ {
		x := find(sets, edge[i].src)
		y := find(sets, edge[i].dest)
		if x != y {
			output += ("(" + strconv.Itoa(edge[i].src) + "->" + strconv.Itoa(edge[i].dest) + " @ " + strconv.Itoa(edge[i].cost) + ") ")
			sum += edge[i].cost
			union(sets, x, y)
		}
	}
	fmt.Println(output)
	fmt.Println("Total MST cost : " + strconv.Itoa(sum))
}

//Testing code
func main19() {
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
	gph.PrimsMST()
	gph.kruskalMST()
}

/*
Edges are : (0->1 @ 4) (5->2 @ 4) (2->3 @ 7) (3->4 @ 9) (6->5 @ 2) (7->6 @ 1) (0->7 @ 8) (2->8 @ 2)
Total MST cost : 37

Edges are : (7->6 @ 1) (2->8 @ 2) (6->5 @ 2) (1->0 @ 4) (5->2 @ 4) (3->2 @ 7) (7->0 @ 8) (3->4 @ 9)
Total MST cost : 37
*/
var INF = 99999

func (gph *Graph) FloydWarshall() {
	V := gph.count
	dist := make([][]int, V)
	path := make([][]int, V)
	for i := range dist {
		dist[i] = make([]int, V)
		path[i] = make([]int, V)
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			path[i][j] = -1
			dist[i][j] = INF
		}
		path[i][i] = 0
	}

	for i := 0; i < V; i++ {
		ad := gph.Edges[i]
		for ad != nil {
			path[i][ad.dest] = i
			dist[i][ad.dest] = ad.cost
			ad = ad.next
		}
	}

	for k := 0; k < V; k++ { // Pick intermediate vertices.
		for i := 0; i < V; i++ { // Pick source vertices one by one.
			for j := 0; j < V; j++ { // Pick destination vertices.
				// If we have shorter path from i to j via k.
				// then update dist[i][j]
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					path[i][j] = path[k][j]
				}
			}
			// dist[i][i] is 0 in the start.
			// If there is a better path from i to i and is better path then we have -ve cycle.                //
			if dist[i][i] < 0 {
				fmt.Println("Negative-weight cycle found.")
				return
			}
		}
	}
	// Print the shortest distance matrix
	printSolution(dist, path, V)
}

func printSolution(cost [][]int, path [][]int, V int) {
	output := "Shortest Paths : "
	for u := 0; u < V; u++ {
		for v := 0; v < V; v++ {
			if u != v && path[u][v] != -1 {
				output += "("
				output += printPath(path, u, v)
				output += " @ " + strconv.Itoa(cost[u][v]) + ") "
			}
		}
	}
	fmt.Println(output)
}

func printPath(path [][]int, u int, v int) string {
	if path[u][v] == u {
		return strconv.Itoa(u) + "->" + strconv.Itoa(v)
	}
	output := printPath(path, u, path[u][v])
	output += "->" + strconv.Itoa(v)
	return output
}

func main20() {
	gph := NewGraph(4)
	gph.AddDirectedEdge(0, 0, 0)
	gph.AddDirectedEdge(1, 1, 0)
	gph.AddDirectedEdge(2, 2, 0)
	gph.AddDirectedEdge(3, 3, 0)
	gph.AddDirectedEdge(0, 1, 5)
	gph.AddDirectedEdge(0, 3, 10)
	gph.AddDirectedEdge(1, 2, 3)
	gph.AddDirectedEdge(2, 3, 1)
	gph.FloydWarshall()
}

/*
Shortest Paths : (0->1 @ 5) (0->1->2 @ 8) (0->1->2->3 @ 9) (1->2 @ 3) (1->2->3 @ 4) (2->3 @ 1)
*/

func main() {
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
	main19()
	main20()
}

// *********************
type Heap struct {
	size int
	arr  []interface{}
	comp func(x interface{}, y interface{}) bool
}

func CreateHeap(comp func(x interface{}, y interface{}) bool) *Heap {
	var arr []interface{}
	h := &Heap{comp: comp, arr: arr, size: 0}
	return h
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) percolateDown(parent int) {
	lChild := 2*parent + 1
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
	h.percolateUp(h.size - 1)
}

func (h *Heap) Remove() interface{} {
	if h.IsEmpty() {
		fmt.Println("HeapEmptyException.")
		return 0
	}
	value := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.size--
	h.percolateDown(0)
	h.arr = h.arr[0:h.size]
	return value
}

func (h *Heap) Delete(value interface{}) bool {
	for i := 0; i < h.size; i++ {
		if h.arr[i] == value {
			h.arr[i] = h.arr[h.size-1]
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

func minComp(i, j interface{}) bool { // always i < j in use
	return i.(int) > j.(int) // swaps for min heap
}

func maxComp(i, j interface{}) bool { // always i < j in use
	return i.(int) < j.(int) // swap for max heap.
}

// ************************************

type Stack struct {
	stk []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[:n-1]
	return value
}

func (s Stack) Top() interface{} {
	n := len(s.stk)
	return s.stk[n-1]
}

func (s Stack) Len() int {
	return len(s.stk)
}

func (s Stack) IsEmpty() bool {
	return len(s.stk) == 0
}

func (s Stack) Print() {
	fmt.Println(s.stk)
}

// ************************************

type Queue struct {
	que []interface{}
}

func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func (q Queue) Front() interface{} {
	return q.que[0]
}

func (q Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

func (q Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q Queue) Len() int {
	return len(q.que)
}

func (q Queue) Print() {
	fmt.Println(q.que)
}

// ************************************
