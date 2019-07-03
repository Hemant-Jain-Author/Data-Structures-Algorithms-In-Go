package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	first := 0
	second := 1
	temp := 0

	if n == 0 {
		return first
	} else if n == 1 {
		return second
	}

	for i := 2; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
	}
	return temp
}

func main1() {
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci2(6))
}
