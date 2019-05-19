package main

import ("fmt"
"sort")


func main1() {
	eleMap := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	var symbols []string
	var elements []string
	for sym,name := range eleMap {
		symbols = append(symbols, sym)
		elements = append(elements, name)
	}

	fmt.Println(symbols)
	fmt.Println(elements)
}

func main2(){
	ages := make(map[string]int) // mapping from strings to ints
	ages["bob"] = 28
	ages["alice"] = 31
	ages["charlie"] = 34
	
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	
	sort.Strings(names)
	
	for _, name := range names {
		fmt.Printf("[%s : %d] ", name, ages[name])
	}
	fmt.Println()
}

func main3() {
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

func main(){
	main3()
}