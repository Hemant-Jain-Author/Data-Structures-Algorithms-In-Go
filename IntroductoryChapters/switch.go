package main

import "fmt"

func main35() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something else")
	}
}

func main36() {
	i := 2
	switch i {
	case 1, 2, 3:
		fmt.Println("one,two or three")
	default:
		fmt.Println("something else")
	}
}

func isEven(value int) {
	switch {
	case value%2 == 0:
		fmt.Println("I is even")
	default:
		fmt.Println("I is odd")
	}
}

// switch with precondition.
func limitCheck(first, second, limit int) bool {
	switch value := max(first, second); {
	case value < limit:
		return true
	case value > limit:
		return false
	}
	return true
}
