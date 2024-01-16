package main

import "fmt"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("7:", IsPrime(7))
	fmt.Println("8:", IsPrime(8))
}

/*
7: true
8: false
*/
