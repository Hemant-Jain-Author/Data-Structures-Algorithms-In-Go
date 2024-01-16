package main

import "fmt"

func GraphColouring(graph [][]bool, V int, m int) bool {
	colour := make([]int, V)
	if graphColouringUtil(graph, V, m, colour, 0) {
		fmt.Println("Assigned colours are:", colour)
		return true
	}
	fmt.Println("Solution does not exist")
	return false
}

func graphColouringUtil(graph [][]bool, V int, m int, colour []int, i int) bool {
	if i == V {
		return true
	}

	for j := 1; j <= m; j++ {
		if isGraphColouredProperly(graph, V, colour, i, j) {
			colour[i] = j
			if graphColouringUtil(graph, V, m, colour, i+1) {
				return true
			}
		}
	}

	return false
}

// Check if the whole graph is coloured properly.
func isGraphColouredProperly(graph [][]bool, V int, colour []int, v int, c int) bool {
	for i := 0; i < V; i++ {
		if graph[v][i] && c == colour[i] {
			return false
		}
	}
	return true
}

func GraphColouring2(graph [][]bool, V int, m int) bool {
	colour := make([]int, V)
	if graphColouringUtil2(graph, V, m, colour, 0) {
		fmt.Println("Assigned colours are:", colour)
		return true
	}
	fmt.Println("Solution does not exist")
	return false
}

func graphColouringUtil2(graph [][]bool, V int, m int, colour []int, i int) bool {
	if i == V {
		return isGraphColouredProperly2(graph, colour, V)
	}

	for j := 1; j <= m; j++ {
		colour[i] = j
		if graphColouringUtil2(graph, V, m, colour, i+1) {
			return true
		}
	}

	return false
}

// Check if the whole graph is coloured properly.
func isGraphColouredProperly2(graph [][]bool, colour []int, V int) bool {
	for i := 0; i < V; i++ {
		for j := i + 1; j < V; j++ {
			if graph[i][j] && colour[j] == colour[i] {
				return false
			}
		}
	}
	return true
}

// Testing code.
func main() {
	graph := [][]bool{
		{false, true, false, false, true},
		{true, false, true, false, true},
		{false, true, false, true, true},
		{false, false, true, false, true},
		{true, true, true, true, false},
	}
	V := 5
	m := 4

	GraphColouring(graph, V, m)
	GraphColouring2(graph, V, m)
}

/*
Assigned colours are: [1 2 1 2 3]
Assigned colours are: [1 2 1 2 3]
*/
