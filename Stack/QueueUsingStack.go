package main

import "fmt"

func main() {
	que := new(QueueUsingStack)
	que.Add(1)
	que.Add(11)
	que.Add(111)
	fmt.Println(que.Remove())
	que.Add(2)
	que.Add(21)
	que.Add(211)
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
}

type QueueUsingStack struct {
	stk1 Stack
	stk2 Stack
}

func (que *QueueUsingStack) Add(value int) {
	que.stk1.Push(value)
}

func (que *QueueUsingStack) Remove() int {
	var value int
	if que.stk2.IsEmpty() == false {
		value = que.stk2.Pop().(int)
		return value
	}

	for que.stk1.IsEmpty() == false {
		value = que.stk1.Pop().(int)
		que.stk2.Push(value)
	}

	value = que.stk2.Pop().(int)
	return value
}
