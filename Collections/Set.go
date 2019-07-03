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
	mp.Add(1)
	mp.Add(2)
	fmt.Println(mp.Find(1))
	fmt.Println(mp.Find(3))
	mp.Remove(1)
	fmt.Println(mp.Find(1))
}
