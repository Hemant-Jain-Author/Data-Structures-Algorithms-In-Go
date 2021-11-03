package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stk *list.List
}

func NewStack() *Stack {
	st := new(Stack)
	st.stk = list.New()
	return st
}

func (st *Stack) Push(data int) {
	st.stk.PushBack(data)
}

func (st *Stack) Pop() int {
	back := st.stk.Back()
	val := back.Value.(int)
	st.stk.Remove(back)
	return val
}
func (st *Stack) Top() int {
	back := st.stk.Back()
	return back.Value.(int)
}

func (st Stack) Len() int {
	return st.stk.Len()
}

func (st Stack) IsEmpty() bool {
	return st.stk.Len() == 0
}

func (st Stack) Print() {
	for e := st.stk.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func main() {

	stk := list.New()
	for i := 0; i < 5; i++ {
		stk.PushBack(i)
	}
	for stk.Len() > 0 {
		back := stk.Back()
		fmt.Print(back.Value, " ")
		stk.Remove(back)
	}


	fmt.Println()
	st := NewStack()
	for i := 0; i < 5; i++ {
		st.Push(i)
	}
	st.Print()
	for st.Len() > 0 {
		fmt.Print(st.Pop(), " ")
	}
}

/*
4 3 2 1 0
*/
