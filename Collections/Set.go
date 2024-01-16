package main

import (
	"fmt"

	"github.com/golang-collections/collections/set"
)

func main() {
    st := set.New()
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