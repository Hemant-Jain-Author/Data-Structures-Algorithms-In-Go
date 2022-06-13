package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
    stk := stack.New()
    stk.Push(1)
    stk.Push(2)
    stk.Push(3)

    fmt.Println("Stack size :", stk.Len());
    fmt.Println("Stack pop :", stk.Pop());
    fmt.Println("Stack top :", stk.Peek());
    fmt.Println("Stack isEmpty :", stk.Len() == 0);
}

/*
Stack size : 3
Stack pop : 3
Stack top : 2
Stack isEmpty : false
*/