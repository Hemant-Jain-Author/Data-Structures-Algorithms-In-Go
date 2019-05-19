package main

import "fmt"

type StackInt struct {
	s []int
}

func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

func (s *StackInt) Pop() int {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *StackInt) Top() int {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func main() {
	s := new(StackInt)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	
	length := s.Length()
	for i := 0; i < length; i++ {
		fmt.Print(s.Pop(), " ")
	}
}