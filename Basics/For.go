package main

import "fmt"

func main1() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is :: ", sum)
}

func main2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0size := len(arr
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is :: ", sum)
}

func main3() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	n := len(numbers)
	for i < n {
		sum += numbers[i]
		i++
	}
	fmt.Println("Sum is :: ", sum)
}

func main4() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	i := 0
	for {
		sum += numbers[i]
		i++
		if i >= len(numbers) {
			break
		}
	}
	fmt.Println("Sum is :: ", sum)
}

func main() {
	main1()
	main2()
	main3()
	main4()
}