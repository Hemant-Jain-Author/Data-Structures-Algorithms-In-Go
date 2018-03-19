package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50
	for key, val := range m {
		fmt.Print("[ ",key," -> ", val," ]")
	}
	fmt.Println()

	fmt.Println("Apple price:", m["Apple"])	
	delete(m, "Apple")
	fmt.Println("Apple price:", m["Apple"])
	
	v, ok := m["Apple"]
	fmt.Println("Apple price:", v, "Present:", ok)
	
	v2, ok2 := m["Banana"]
	fmt.Println("Banana price:", v2, "Present:", ok2)
}
/*
[ Apple -> 40 ][ Banana -> 30 ][ Mango -> 50 ]
Apple price: 40
Apple price: 0
Apple price: 0 Present: false
Banana price: 30 Present: true
*/

func main22() {
	m := make(map[string]int) // [1]
	m["k1"] = 7               // [2]
	m["k2"] = 13
	fmt.Println("map:", m) // [3]
	v1 := m["k1"]          // [4]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m)) // [5]
	delete(m, "k2")             // [6]
	fmt.Println("map:", m)
	_, prs := m["k2"] // [7]
	fmt.Println("prs:", prs)
	n := map[string]int{"foo": 1, "bar": 2} // [8]
	fmt.Println("map:", n)
}

// You are creating an empty map using the format make(map[key-type]value-type
