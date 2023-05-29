package main

import "fmt"

func power(x int, n int) int {
	if n == 0 {
		return 1
	}

	value := power(x, n/2)
	result := value * value

	if n%2 == 1 {
		result *= x
	}

	return result
}

func main() {
	fmt.Println(power(5, 2))
}

/*
25
*/
