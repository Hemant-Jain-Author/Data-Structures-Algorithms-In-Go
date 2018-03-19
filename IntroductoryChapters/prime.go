package main

import "fmt"

func isPrime(n int) bool {
	if( n < 1){
		return false
	} 
	i := 2
	for (i*i <= n) {
		if(n%i == 0){
			return false
		}
	i += 1
	}
	return true
}