package main

import "fmt"

func power(x int, n int) int{
    var value int 
    if (n == 0) {
        return 1;
    } else if n % 2 == 0 {
        value = power(x, n / 2)
        return (value * value)
    } else {
        value = power(x, n / 2)
        return (x * value * value)
    }
}

func main() {
	fmt.Println(power(5, 2))
}

/*
25
*/