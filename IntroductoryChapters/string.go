package main

import "fmt"

func main34() {
	s := "hello, World!"
	r := []rune(s)
	r[0] = 'H'
	s2 := string(r)
	fmt.Println(s2)
}
