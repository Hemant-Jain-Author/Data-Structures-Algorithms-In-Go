package main

import "fmt"

func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	TOHUtil(num, "A", "C", "B")
}

func TOHUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}

	TOHUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	TOHUtil(num-1, temp, to, from)
}


func main4() {
	TowersOfHanoi(3)
}
