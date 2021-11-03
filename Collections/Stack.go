package main

import "fmt"

type Stack []int

func (stk Stack) Len() int {
	return len(stk)
}

func (stk *Stack) Push(x int) {
	*stk = append(*stk, x)
}

func (stk *Stack) Pop() int {
	n := len(*stk)
	value := (*stk)[n-1]
	*stk = (*stk)[: n-1]
	return value
}

func (stk *Stack) Top() int {
	n := len(*stk)
	return (*stk)[n-1]
}

func (stk Stack) IsEmpty() bool {
	return len(stk) == 0
}

func (stk Stack) Print() {
	fmt.Println(stk)
}

func main() {
	stk := &Stack{}
	for i := 0; i < 5; i++ {
		stk.Push(i)
	}
	for stk.IsEmpty() == false {
		fmt.Print(stk.Pop(), " ")
	}
}

/*
4 3 2 1 0 
*/