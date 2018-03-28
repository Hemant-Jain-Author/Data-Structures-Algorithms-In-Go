package main

import "fmt"

func main7() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Println("Sum is :: ", sum)
}

func main8() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	fmt.Println("Sum is :: ", sum)
}

func main9() {
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

func main10() {
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

func main11() {
	main2()
	main3()
	main4()
	main5()
}
