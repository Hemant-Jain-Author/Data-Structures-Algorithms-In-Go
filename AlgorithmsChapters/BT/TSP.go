package main

import (
	"fmt"
	"math"
)

// Function to find the minimum weight Hamiltonian Cycle 
func TSP(graph [][]int, n int) (int) {
	visited := make([]bool, n)
	path := make([]int, n)
	path[0] = 0
	visited[0] = true
	ans := math.MaxInt
	ans = tspUtil(graph, n, path, 1, 0, visited, ans)
	return ans
}

func tspUtil(graph [][]int, n int, path []int, pSize int, 
	pCost int, visited []bool, ans int) (int) {
	curr := path[n-1]
	if pSize == n && graph[curr][0] > 0 {
		ans = min(ans, pCost+graph[curr][0])
		return ans
	}
	for i := 0; i < n; i++ {
		if visited[i] == false && graph[curr][i] > 0 {
			visited[i] = true
			path[pSize] = i
			ans = tspUtil(graph, n, path, pSize+1, pCost+graph[curr][i], visited, ans)
			visited[i] = false
		}
	}
	return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	n := 4
	graph := [][]int{
		{0, 10, 15, 20}, 
		{10, 0, 35, 25}, 
		{15, 35, 0, 30}, 
		{20, 25, 30, 0}}
	fmt.Println("Shortest Path:", TSP(graph, n))
}

/*
Shortest Path: 65
*/