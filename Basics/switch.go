package main

import (
"fmt"
)

func main1() {
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

func main2() {
	i := 2
	switch i {
	case 1, 2, 3:
		fmt.Println("one,two or three")
	default:
		fmt.Println("something else")
	}
}

func isEven(value int) bool {
	switch {
	case value%2 == 0:
		fmt.Println("I is even")
		return true
	default:
		fmt.Println("I is odd")
		return false
	}
}

// switch with precondition.
func limitCheck(first, second int) bool {
	switch value := 100; {
	case value > first && value < second:
		return true
	case value > second || value < first:
		return false
	}
	return true
}

func main(){
	main1()
	main2()
	fmt.Println(isEven(100))
	fmt.Println(isEven(103))
	fmt.Println(limitCheck(10,40))
	fmt.Println(limitCheck(10,400))
}