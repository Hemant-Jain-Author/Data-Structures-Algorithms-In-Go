package main

import "fmt"

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

func main() {
	mp := make(Set)
	mp.Add("a")
	mp.Add("b")
	mp.Add("a")
	fmt.Println(mp.Find("a"))
	fmt.Println(mp.Find("b"))
	fmt.Println(mp.Find("c"))
}
