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

	mp2 := make(SetInt)
	mp2.Add(4)
	mp2.Add(5)
	mp3 := make(SetInt)
	mp3.Add(5)
	mp3.Add(6)
	fmt.Println(mp2.Union(mp3))
	fmt.Println(mp2.Intersection(mp3))
}


type SetInt map[int]bool

func (s *SetInt) Add(key int) {
	(*s)[key] = true
}
func (s *SetInt) Remove(key int) {
	delete((*s), key)
}

func (s SetInt) Find(key int) bool {
	return s[key]
}

func (s SetInt) Intersection(s2 SetInt) []int {
	ans :=  make([]int,0)
	for val := range s2 {
		if s[val] == true{
			ans = append(ans, val)
		}
	}
	return ans
}

func (s SetInt) Union(s2 SetInt) []int {
	uni := make(SetInt)
	ans :=  make([]int,0)
	for val := range s {
		uni[val] = true
	}
	for val := range s2 {
		uni[val] = true
	}
	for val := range uni {
		ans = append(ans, val)
	}
	return ans
}

