package main

import (
	"fmt"
	"math"
)

var INF = math.MaxInt32

// Returns shortest distance from 0 to N-1.
func GraphShortestDist(graph [][]int, n int) int {
	// dist[i] is going to store shortest
	// distance from node i to node n-1.
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	path := make([]int, n)
	var value int
	dist[0] = 0
	path[0] = -1

	// Calculating shortest path for the nodes
	for i := 0; i < n; i++ {
		// Check all nodes of next
		for j := i; j < n; j++ {
			// Reject if no edge exists
			if graph[i][j] == INF {
				continue
			}
			value = graph[i][j] + dist[i]
			if dist[j] > value {
				dist[j] = value
				path[j] = i
			}
		}
	}
	value = n - 1
	fmt.Print("Path: ")
	for value != -1 {
		fmt.Print(value, " ")
		value = path[value]
	}
	fmt.Println()
	return dist[n-1]
}

func main() {
	graph := [][]int{
		{INF, 1, 2, 5, INF, INF, INF, INF},
		{INF, INF, INF, INF, 4, 11, INF, INF},
		{INF, INF, INF, INF, 9, 5, 16, INF},
		{INF, INF, INF, INF, INF, INF, 2, INF},
		{INF, INF, INF, INF, INF, INF, INF, 18},
		{INF, INF, INF, INF, INF, INF, INF, 13},
		{INF, INF, INF, INF, INF, INF, INF, 2},
		{INF, INF, INF, INF, INF, INF, INF, INF}}

	fmt.Println("Total Cost:", GraphShortestDist(graph, 8))
}

/*
Path: 7 6 3 0
Total Cost: 9
*/
