package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for key, val := range m {
		fmt.Print("[ ", key, " @ ", val, " ]")
	}
	fmt.Println()

	v, ok := m["Apple"]
	if ok {
		fmt.Println("Apple available at price :", v)
	} else {
		fmt.Println("Apple unavailable.")
	}

	delete(m, "Apple")

	v, ok = m["Apple"]
	if ok {
		fmt.Println("Apple available at price :", v)
	} else {
		fmt.Println("Apple unavailable.")
	}
}

/*
[ Apple @ 40 ][ Banana @ 30 ][ Mango @ 50 ]
Apple available at price : 40
Apple unavailable.
*/