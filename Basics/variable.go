package main

import "fmt"

func main1() {
	var v1 int
	var v2 int
	v1 = 100
	v2 = 200
	fmt.Println("Value stored in variable v1 :: ", v1)
	fmt.Println("Value stored in variable v2 :: ", v2)
}

func main2() {
	var v1, v2 int
	v1 = 100
	fmt.Println("Value stored in variable v1 :: ", v1)
	fmt.Println("Value stored in variable v2 :: ", v2)
}

func main3() {
	v1 := 100
	v2 := 200
	fmt.Println("Value stored in variable v1 :: ", v1)
	fmt.Println("Value stored in variable v2 :: ", v2)
}

const PI = 3.14

// Declare global variables
var g1 int = 100

func main4() {
	// Declaring a variable in a function
	v1 := 10
	v2 := 20
	fmt.Println("Value stored in variable v1 :: ", v1)
	fmt.Println("Value stored in variable v2 :: ", v2)
	fmt.Println("Value stored in variable g1 :: ", g1)
}

func main(){
	main1()
	main2()
	main3()
	main4()
}