package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	
	for index, val := range numbers {
		sum += val
		fmt.Print("[", index, ",", val, "] ")
	}

	fmt.Println("\nSum is :: ", sum)

	kvs := map[int]string{1: "apple", 2: "banana"}
	
	for k, v := range kvs {
		fmt.Println(k, " -> ", v)
	}

	for i, c := range "Hello, World!" {
		fmt.Print("[", i, ",", string(c), "] ")
	}
}

/*
[0,1] [1,2] [2,3] [3,4] [4,5] [5,6] [6,7] [7,8] [8,9] [9,10]
Sum is ::  55
1  ->  apple
2  ->  banana
[0,H] [1,e] [2,l] [3,l] [4,o] [5,,] [6, ] [7,W] [8,o] [9,r] [10,l] [11,d] [12,!]
*/