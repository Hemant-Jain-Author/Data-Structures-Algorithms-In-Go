package main

import "fmt"

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

func (que *QueueUsingStack) Length() int {
    return (que.stk1.Len() + que.stk2.Len());
}

func (que *QueueUsingStack) IsEmpty() bool {
    return (que.stk1.Len() + que.stk2.Len()) == 0;
}


func main() {
	que := new(QueueUsingStack)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println(que.Length())
	fmt.Println(que.IsEmpty())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.IsEmpty())
}

/*
3
false
1
2
3
true
*/

type Stack struct {
	stk []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[: n-1]
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