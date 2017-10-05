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

func RemoveDuplicate(str string) string {
	input := []rune(str)
	hs := make(Set)
	var output []rune
	for _, ch := range input {
		if hs.Find(ch) == false {
			output = append(output, ch)
			hs.Add(ch)
		}
	}
	return string(output)
}

func FindMissing(arr []int, start int, end int) (int, bool) {
	hs := make(Set)
	for _, i := range arr {
		hs.Add(i)
	}
	for curr := start; curr <= end; curr++ {
		if hs.Find(curr) == false {
			return curr, true
		}
	}
	return 0, false
}

func PrintRepeating(arr []int) {
	hs := make(Set)
	fmt.Print("Repeating elements are :: ")
	for _, val := range arr {
		if hs.Find(val) {
			fmt.Print(val, "  ")
		} else {
			hs.Add(val)
		}
	}
	fmt.Println()
}

func PrintFirstRepeating(arr []int) {
	size := len(arr)
	hs := make(Counter)

	for i := 0; i < size; i++ {
		hs.Add(arr[i])
	}
	for i := 0; i < size; i++ {
		hs.Remove(arr[i])
		if hs.Find(arr[i]) {
			fmt.Println("First Repeating number is : ", arr[i])
			return
		}
	}
}

func hornerHash(key []rune, tableSize int) int {
	size := len(key)
	h := 0
	for i := 0; i < size; i++ {
		h = (32*h + (int)(key[i])) % tableSize
	}
	return h
}

func main() {
	var1 := "hello"
	var2 := "elloh"
	var3 := "world"

	fmt.Println("isAnagram : ", isAnagram(var1, var2))
	fmt.Println("isAnagram : ", isAnagram(var1, var3))

	fmt.Println(removeDuplicate(var1))

	arr := []int{1, 2, 3, 5, 6, 7, 9, 8, 10}
	fmt.Print("Missing number is :: ")
	fmt.Println(findMissing(arr, 1, 10))

	arr1 := []int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 1}
	printRepeating(arr1)

	printFirstRepeating(arr1)
}
