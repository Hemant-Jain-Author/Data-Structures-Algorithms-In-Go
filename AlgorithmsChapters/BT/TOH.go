package main

import "fmt"

func TowerOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	tohUtil(num, "A", "C", "B")
}

func tohUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	tohUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	tohUtil(num-1, temp, to, from)
}

// Testing code.
func main() {
	TowerOfHanoi(3)
}

/*
The sequence of moves involved in the Tower of Hanoi are :
Move disk  1  from peg  A  to peg  C
Move disk  2  from peg  A  to peg  B
Move disk  1  from peg  C  to peg  B
Move disk  3  from peg  A  to peg  C
Move disk  1  from peg  B  to peg  A
Move disk  2  from peg  B  to peg  C
Move disk  1  from peg  A  to peg  C
*/