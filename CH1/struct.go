package main

import "fmt"

type student struct {
	rollNo int
	name   string
}

func main() {
	stud := student{1, "johny"} 
	fmt.Println(stud) 
	fmt.Println("Student name ::",stud.name) // Accessing inner fields.

	pstud := &stud
    fmt.Println("Student name ::",pstud.name) // Accessing inner fields.

	fmt.Println(student{rollNo: 2, name: "Ann"}) // Named initialization.
	fmt.Println(student{name: "Ann", rollNo: 2}) // Order does not matter.
	fmt.Println(student{name: "Alice"}) // Default initialization of rollNo.	
}

/*
{1 johny}
Student name :: johny
Student name :: johny
{2 Ann}
{2 Ann}
{0 Alice}
*/