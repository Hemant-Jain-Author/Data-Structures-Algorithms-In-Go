package main

import "fmt"

func Fibonacci(n int) int {
	if n < 2 {
		return n	
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciBU2(n int) int {
	if n < 2 {
		return n
	}
	first := 0
	second := 1
	temp := 0
	for i := 2; i <= n; i++ {
		temp = first + second
		first = second
		second = temp
	}
	return temp
}

func FibonacciBU(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n]
}

func fibonacciTDUtil(n int, dp []int) int {
	if n < 2 {
		dp[n] = n
		return n
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = fibonacciTDUtil(n-1, dp) + fibonacciTDUtil(n-2, dp)
	return dp[n]
}

func FibonacciTD(n int) int {
	dp := make([]int, n+1)
	fibonacciTDUtil(n, dp)
	return dp[n]
}
func main() {
	fmt.Println(Fibonacci(10))
	fmt.Println(FibonacciBU(10))
	fmt.Println(FibonacciBU2(10))
	fmt.Println(FibonacciTD(10))
}

/*
55
55
55
55
*/