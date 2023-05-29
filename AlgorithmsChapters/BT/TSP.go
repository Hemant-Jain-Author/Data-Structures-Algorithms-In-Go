package main

import (
	"fmt"
	"math"
)

// Function to find the minimum weight Hamiltonian Cycle
func TSP(graph [][]int, n int) int {
	visited := make([]bool, n)
	path := make([]int, n)
	ansPath := make([]int, n+1)
	path[0] = 0
	visited[0] = true
	ans := math.MaxInt64
	ans = tspUtil(graph, n, path, 1, 0, visited, ans, ansPath)
	fmt.Println("Path length:", ans)
	fmt.Print("Path: ")
	for i := 0; i <= n; i++ {
		fmt.Print(ansPath[i], " ")
	}
	return ans
}

func tspUtil(graph [][]int, n int, path []int, pSize int, pCost int, visited []bool, ans int, ansPath []int) int {
	if pCost > ans {
		return ans
	}

	curr := path[pSize-1]
	if pSize == n {
		if graph[curr][0] > 0 && ans > pCost+graph[curr][0] {
			ans = pCost + graph[curr][0]
			for i := 0; i <= n; i++ {
				ansPath[i] = path[i%n]
			}
		}
		return ans
	}

	for i := 0; i < n; i++ {
		if !visited[i] && graph[curr][i] > 0 {
			visited[i] = true
			path[pSize] = i
			ans = tspUtil(graph, n, path, pSize+1, pCost+graph[curr][i], visited, ans, ansPath)
			visited[i] = false
		}
	}
	return ans
}

// Testing code.
func main() {
	n := 4
	graph := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}
	TSP(graph, n)
}

/*
Path length: 80
Path: 0 1 3 2 0
*/
