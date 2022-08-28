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
		cm.Add(ch)
	}
	for _, ch := range str2 {
		cm.Remove(ch)
	}
	return (len(cm) == 0)
}

func main1() {
	var1 := "hello"
	var2 := "elloh"
	var3 := "world"

	fmt.Println("IsAnagram : ", IsAnagram(var1, var2))
	fmt.Println("IsAnagram : ", IsAnagram(var1, var3))
}

/*
IsAnagram :  true
IsAnagram :  false
*/

func RemoveDuplicate(str string) string {
	input := []rune(str)
	hs := make(Set)
	var output []rune
	for _, ch := range input {
		if hs.Has(ch) == false {
			output = append(output, ch)
			hs.Add(ch)
		}
	}
	return string(output)
}

func main2() {
	var1 := "hello"
	fmt.Println(RemoveDuplicate(var1))
}

/*
helo
*/

func FindMissing(arr []int, start int, end int) (int, bool) {
	hs := make(Set)
	for _, i := range arr {
		hs.Add(i)
	}
	for curr := start; curr <= end; curr++ {
		if hs.Has(curr) == false {
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
Missing number is ::  4
*/

func PrintRepeating(arr []int) {
	hs := make(Set)
	fmt.Print("Repeating elements are :: ")
	for _, val := range arr {
		if hs.Has(val) {
			fmt.Print(val, "  ")
		} else {
			hs.Add(val)
		}
	}
	fmt.Println()
}

func main4() {
	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	PrintRepeating(arr1)
}

/*
Repeating elements are :: 4  1  
*/

func PrintFirstRepeating(arr []int) {
	size := len(arr)
	hs := make(Counter)

	for i := 0; i < size; i++ {
		hs.Add(arr[i])
	}
	for i := 0; i < size; i++ {
		hs.Remove(arr[i])
		if hs.Has(arr[i]) {
			fmt.Println("First Repeating number is : ", arr[i])
			return
		}
	}
}

func main5() {
	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	PrintFirstRepeating(arr1)
}

/*
First Repeating number is :  1
*/

func hornerHash(key []rune, tableSize int) int {
	size := len(key)
	h := 0
	for i := 0; i < size; i++ {
		h = (32*h + (int)(key[i])) % tableSize
	}
	return h
}

func main(){
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
    // Create a hash set.
    hs := make(Set)
    // Add elements to the hash set.
    hs.Add("Banana");
    hs.Add("Apple");
    hs.Add("Mango");
    fmt.Println(hs);
    fmt.Println("Apple present :", hs.Has("Apple"));
    fmt.Println("Grapes present :", hs.Has("Grapes"));
    hs.Remove("Apple");
    fmt.Println(hs);
    fmt.Println("Apple present :", hs.Has("Apple"));
}

/*
map[Apple:true Banana:true Mango:true]
Apple present : true
Grapes present : false
map[Banana:true Mango:true]
Apple present : false
*/

func main7() {
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

func main8() {
	mp := make(Counter)
	mp.Add("a")
	mp.Add("b")
	mp.Add("a")

	fmt.Println(mp.Has("a"))
	fmt.Println(mp.Has("b"))
	fmt.Println(mp.Has("c"))

	fmt.Println(mp.Get("a"))
	fmt.Println(mp.Get("b"))
	fmt.Println(mp.Get("c"))
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

func (s *Counter) Add(key interface{}) {
	(*s)[key]++
}
func (s *Counter) Has(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Counter) Remove(key interface{}) {
	val, ok := (*s)[key]
	if ok == false {
		return
	} else if val == 1 {
		delete((*s), key)
		return
	}
	(*s)[key]--
}

func (s *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*s)[key]
	return val, ok
}

//*******************************************

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Has(key interface{}) bool {
	return (*s)[key]
}
//**************************************