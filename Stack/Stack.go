package main

import "fmt"

type Stack struct {
	stk []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[:n-1]
	return value
}

func (s *Stack) Top() interface{} {
	n := len(s.stk)
	return s.stk[n-1]
}

func (s Stack) Len() int {
	return len(s.stk)
}

func (s Stack) IsEmpty() bool {
	return len(s.stk) == 0
}

func (s Stack) Print() {
	fmt.Println(s.stk)
}

func main() {
	stk := &Stack{}
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Print()
	fmt.Print(stk.Pop(), " ")
	fmt.Print(stk.Pop(), " ")
}

/*
[1 2 3]
3 2 
*/
