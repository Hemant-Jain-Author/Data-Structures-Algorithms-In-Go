package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func more(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

// If with precondition.
func maxAreaCheck(length int, width int, limit int) bool {
	if area := length * width; area < limit {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Max : ", max(35, 33))
	fmt.Println("More : ", more(35, 33))
	fmt.Println("Max area check : ", maxAreaCheck(30, 10, 200))
}