package main

import "fmt"
import "github.com/golang-collections/collections/queue"
import "github.com/golang-collections/collections/stack"
import "github.com/golang-collections/collections/set"

func main1() {
	st := set.New()
	st.Insert(1)
	st.Insert(2)
	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
}
func main2() {
	s := stack.New()
	s.Push(2)
	s.Push(3)
	s.Push(4)

	for s.Len() != 0 {
		val := s.Pop()
		fmt.Print(val, " ")
	}
}


func main() {
	q := queue.New()
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}
