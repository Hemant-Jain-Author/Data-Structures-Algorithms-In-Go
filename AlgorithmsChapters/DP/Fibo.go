package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 2 {
		return n - 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func FibonacciBU(n int) int {
	if n <= 2 {
		return n - 1
	}
	first := 0
	second := 1
	temp := 0
	for i := 2; i < n; i++ {
		temp = first + second
		first = second
		second = temp
	}
	return temp
}
func FibonacciSeries(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(Fibonacci(i), " ")
	}
}
func FibonacciSeriesBU(n int) {
	if n < 1 {
		return
	}
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	fmt.Print(dp)
}
func fibonacciSeriesTDUtil(n int, dp []int) int {
	if n <= 1 {
		dp[n] = n
		return n
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = fibonacciSeriesTDUtil(n-1, dp) + fibonacciSeriesTDUtil(n-2, dp)
	return dp[n]
}

func FibonacciSeriesTD(n int) {
	if n < 1 {
		return
	}
	dp := make([]int, n)
	fibonacciSeriesTDUtil(n-1, dp)
	fmt.Print(dp)
}
func main() {
	for i := 6; i < 7; i++ {
		FibonacciSeries(i)
		fmt.Println()
		FibonacciSeriesBU(i)
		fmt.Println()
		FibonacciSeriesTD(i)
		fmt.Println()
	}
	fmt.Println(Fibonacci(6))
	fmt.Println(FibonacciBU(6))
}

/*
0 1 1 2 3 5 
[0 1 1 2 3 5]
[0 1 1 2 3 5]
5
5
*/