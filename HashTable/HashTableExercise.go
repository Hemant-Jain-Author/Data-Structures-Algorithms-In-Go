package main

import "fmt"

func IsAnagram(str1 string, str2 string) bool {
	size1 := len(str1)
	size2 := len(str2)
	if size1 != size2 {
		return false
	}
	cm := make(Counter)
	for _, ch := range str1 {
		cm.add(ch)
	}
	for _, ch := range str2 {
		cm.remove(ch)
	}
	return len(cm) == 0
}

func main1() {
	var1 := "hello"
	var2 := "elloh"
	var3 := "world"

	fmt.Println("IsAnagram:", IsAnagram(var1, var2))
	fmt.Println("IsAnagram:", IsAnagram(var1, var3))
}

/*
IsAnagram: true
IsAnagram: false
*/

func RemoveDuplicate(str string) string {
	input := []rune(str)
	hs := make(Set)
	var output []rune
	for _, ch := range input {
		if !hs.has(ch) {
			output = append(output, ch)
			hs.add(ch)
		}
	}
	return string(output)
}

func main2() {
	var1 := "hello"
	fmt.Println("RemoveDuplicate:", RemoveDuplicate(var1))
}

/*
helo
*/

func FindMissing(arr []int, start int, end int) (int, bool) {
	hs := make(Set)
	for _, i := range arr {
		hs.add(i)
	}
	for curr := start; curr <= end; curr++ {
		if !hs.has(curr) {
			return curr, true
		}
	}
	return 0, false
}

func main3() {
	arr := []int{1, 2, 3, 5, 6, 7, 9, 8, 10}
	num, _ := FindMissing(arr, 1, 10)
	fmt.Println("Missing number is :: ", num)
}

/*
Missing number is: 4
*/

func PrintRepeating(arr []int) {
	hs := make(Set)
	fmt.Print("Repeating elements are :: ")
	for _, val := range arr {
		if hs.has(val) {
			fmt.Print(val, " ")
		} else {
			hs.add(val)
		}
	}
	fmt.Println()
}

func main4() {
	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	PrintRepeating(arr1)
}

/*
Repeating elements are: 4 1
*/

func PrintFirstRepeating(arr []int) {
	size := len(arr)
	hs := make(Counter)

	for i := 0; i < size; i++ {
		hs.add(arr[i])
	}
	for i := 0; i < size; i++ {
		hs.remove(arr[i])
		if hs.has(arr[i]) {
			fmt.Println("First Repeating number is:", arr[i])
			return
		}
	}
}

func main5() {
	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	PrintFirstRepeating(arr1)
}

/*
First Repeating number is: 1
*/

func hornerHash(key []rune, tableSize int) int {
	size := len(key)
	h := 0
	for i := 0; i < size; i++ {
		h = (32*h + int(key[i])) % tableSize
	}
	return h
}

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
	main8()
}

func main6() {
	hs := make(Set)
	hs.add("Banana")
	hs.add("Apple")
	hs.add("Mango")
	fmt.Println(hs)
	fmt.Println("Apple present:", hs.has("Apple"))
	fmt.Println("Grapes present:", hs.has("Grapes"))
	hs.remove("Apple")
	fmt.Println(hs)
	fmt.Println("Apple present:", hs.has("Apple"))
}

/*
map[Apple:true Banana:true Mango:true]
Apple present: true
Grapes present: false
map[Banana:true Mango:true]
Apple present: false
*/

func main7() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Banana"] = 30
	m["Mango"] = 50

	for key, val := range m {
		fmt.Print("[", key, " @ ", val, "]")
	}
	fmt.Println()

	v, ok := m["Apple"]
	if ok {
		fmt.Println("Apple available at price:", v)
	} else {
		fmt.Println("Apple unavailable.")
	}

	delete(m, "Apple")

	v, ok = m["Apple"]
	if ok {
		fmt.Println("Apple available at price:", v)
	} else {
		fmt.Println("Apple unavailable.")
	}
}

/*
[Apple @ 40][Banana @ 30][Mango @ 50]
Apple available at price: 40
Apple unavailable.
*/

func main8() {
	mp := make(Counter)
	mp.add("a")
	mp.add("b")
	mp.add("a")

	fmt.Println(mp.has("a"))
	fmt.Println(mp.has("b"))
	fmt.Println(mp.has("c"))

	fmt.Println(mp.get("a"))
	fmt.Println(mp.get("b"))
	fmt.Println(mp.get("c"))
}

/*
true
true
false
2 true
1 true
0 false
*/

// ***********************
type Counter map[interface{}]int

func (c *Counter) add(key interface{}) {
	(*c)[key]++
}

func (c *Counter) has(key interface{}) bool {
	_, ok := (*c)[key]
	return ok
}

func (c *Counter) remove(key interface{}) {
	val, ok := (*c)[key]
	if !ok {
		return
	} else if val == 1 {
		delete(*c, key)
		return
	}
	(*c)[key]--
}

func (c *Counter) get(key interface{}) (int, bool) {
	val, ok := (*c)[key]
	return val, ok
}

//*******************************************

type Set map[interface{}]bool

func (s *Set) add(key interface{}) {
	(*s)[key] = true
}

func (s *Set) remove(key interface{}) {
	delete(*s, key)
}

func (s *Set) has(key interface{}) bool {
	return (*s)[key]
}

//**************************************
