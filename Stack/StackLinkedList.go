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

func (st *Stack) Push(data interface{}) {
	st.stk.PushBack(data)
}

func (st *Stack) Pop() interface{} {
	back := st.stk.Back()
	value := back.Value
	st.stk.Remove(back)
	return value
}

func (st *Stack) Top() interface{} {
	return st.stk.Back().Value
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
	stk := NewStack()
	for i := 0; i < 5; i++ {
		stk.Push(i)
	}
	//stk.Print()
	for stk.IsEmpty() == false {
		fmt.Print(stk.Pop(), " ")
	}
}

/*
4 3 2 1 0
*/
