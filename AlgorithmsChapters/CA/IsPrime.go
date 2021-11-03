package main

import "fmt"

func IsPrime(n int) (bool) {
	answer := true
	if n < 2 {
		answer = false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			answer = false
			break
		}
	}
	return answer
}

func main() {
	for i := 5; i <= 10; i++ {
		fmt.Println(i, IsPrime(i))
	}
}

/*
5 true
6 false
7 true
8 false
9 false
10 false
*/