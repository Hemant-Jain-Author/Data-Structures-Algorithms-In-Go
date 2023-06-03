package main

import (
	"fmt"
	"math"
)

var INF = math.MaxInt32

func FloydWarshall(graph [][]int, V int) {
	dist := make([][]int, V)
	for i := range dist {
		dist[i] = make([]int, V)
		copy(dist[i], graph[i])
	}

	// Pick intermediate vertices.
	for k := 0; k < V; k++ {
		// Pick source vertices one by one.
		for i := 0; i < V; i++ {
			// Pick destination vertices.
			for j := 0; j < V; j++ {
				// If we have a shorter path from i to j via k, then update dist[i][j].
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	// Print the shortest distance matrix.
	printSolution(dist, V)
}

func printSolution(dist [][]int, V int) {
	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if dist[i][j] == INF {
				fmt.Print("INF ")
			} else {
				fmt.Print(dist[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func main() {
	graph := [][]int{
		{0, 2, 4, INF, INF, INF, INF},
		{2, 0, 4, 1, INF, INF, INF},
		{4, 4, 0, 2, 8, 4, INF},
		{INF, 1, 2, 0, 3, INF, 6},
		{INF, INF, 6, 4, 0, 3, 1},
		{INF, INF, 4, INF, 4, 0, 2},
		{INF, INF, INF, 4, 2, 3, 0},
	}
	FloydWarshall(graph, 7)
}

/*
0 2 4 3 6 8 7
2 0 3 1 4 7 5
4 3 0 2 5 4 6
3 1 2 0 3 6 4
7 5 6 4 0 3 1
8 7 4 6 4 0 2
7 5 6 4 2 3 0
*/
