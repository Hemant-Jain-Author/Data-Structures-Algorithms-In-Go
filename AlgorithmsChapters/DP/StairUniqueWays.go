package main

import "fmt"

func StairUniqueWaysBU(n int) int {
	if n <= 2 {
		return n
	}
	first := 1
	second := 2
	temp := 0
	for i := 3; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
	}
	return temp
}
func StairUniqueWaysBU2(n int) int {
	if n < 2 {
		return n
	}
	ways := make([]int, n)
	ways[0] = 1
	ways[1] = 2
	for i := 2; i < n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n-1]
}
func main() {
	fmt.Println("Unique way to reach top::", StairUniqueWaysBU(4))
	fmt.Println("Unique way to reach top::", StairUniqueWaysBU2(4))
}

/*
Unique way to reach top:: 5
Unique way to reach top:: 5
*/
