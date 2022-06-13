package main

import (
	"fmt"
)

type Set map[interface{}]bool

func (s *Set) Insert(key interface{}) {
    (*s)[key] = true
}

func (s *Set) Remove(key interface{}) {
    delete((*s), key)
}

func (s *Set) Has(key interface{}) bool {
    return (*s)[key]
}

func main() {
    st := make(Set)
	st.Insert("Banana")
	st.Insert("Apple")
    st.Insert("Mango")

    fmt.Println("Apple present :", st.Has("Apple"))
	fmt.Println("Grapes present :", st.Has("Grapes"))
    st.Remove("Apple")
    fmt.Println("Apple present :", st.Has("Apple"))
}

/*
Apple present : true
Grapes present : false
Apple present : false
*/