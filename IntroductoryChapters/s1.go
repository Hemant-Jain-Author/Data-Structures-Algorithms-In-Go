package main

import "fmt"
/*
func main2() {
	s := make([]string, 3) // [1]
	fmt.Println("emp:", s)
	s[0] = "a" // [2]
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s)) // [3]
	s = append(s, "d")          // [4]
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	c := make([]string, len(s)) // [5]
	copy(c, s)
	fmt.Println("cpy:", c)
	l := s[2:5] // [6]
	fmt.Println("sl1:", l)
	l = s[:5] // [7]
	fmt.Println("sl2:", l)
	l = s[2:] // [8]
	fmt.Println("sl3:", l)
	t := []string{"g", "h", "i"} // [9]
	fmt.Println("dcl:", t)
	towD := make([][]int, 3) // [10]
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Prinln("2d:", twoD)
}
*/
/*
func main() {
	var s []int
	for i := 1; i <= 17; i++ {
		s = append(s, i)
		PrintSlice(s)		
	}
}
/*
[1] :: len=1 cap=1
[1 2] :: len=2 cap=2
[1 2 3] :: len=3 cap=4
[1 2 3 4] :: len=4 cap=4
[1 2 3 4 5] :: len=5 cap=8
[1 2 3 4 5 6] :: len=6 cap=8
[1 2 3 4 5 6 7] :: len=7 cap=8
[1 2 3 4 5 6 7 8] :: len=8 cap=8
[1 2 3 4 5 6 7 8 9] :: len=9 cap=16
[1 2 3 4 5 6 7 8 9 10] :: len=10 cap=16
[1 2 3 4 5 6 7 8 9 10 11] :: len=11 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12] :: len=12 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12 13] :: len=13 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12 13 14] :: len=14 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15] :: len=15 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16] :: len=16 cap=16
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] :: len=17 cap=32
*/
func main() {	
	s := []int{1,2,3,4,5,6,7,8,9,10}
	PrintSlice(s)
	a := make([]int, 10)
	PrintSlice(a)
	b := make([]int, 0, 10)
	PrintSlice(b)
	c := s[0:4]
	PrintSlice(c)
	d := c[2:5]
	PrintSlice(d)
}

[1 2 3 4 5 6 7 8 9 10] :: len=10 cap=10
[0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10
[] :: len=0 cap=10
[1 2 3 4] :: len=4 cap=10
[3 4 5] :: len=3 cap=8

func PrintSlice(data []int) {
	fmt.Printf("%v :: len=%d cap=%d \n", data, len(data), cap(data))	
}
/*
[0 1 2 3 4 5 6 7 8 9] :: len=10 cap=16
[0 0 0 0 0 0 0 0 0 0] :: len=10 cap=10
[] :: len=0 cap=10
[0 1 2 3] :: len=4 cap=16
[2 3 4] :: len=3 cap=14
*/
