package main

import "fmt"

type MyInt int

func (data MyInt) increment1() {
	data = data + 1
}

func (data *MyInt) increment2() {
	*data = *data + 1
}

func main() {
	var data MyInt = 1
	fmt.Println("value before increment1() call :", data)
	data.increment1()
	fmt.Println("value after increment1() call :", data)
	data.increment2()
	fmt.Println("value after increment2() call :", data)
}
/*
value before increment1() call : 1
value after increment1() call : 1
value after increment2() call : 2
*/