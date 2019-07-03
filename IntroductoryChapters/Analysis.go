package main

import (
	"fmt"
	"math"
)

func fun1(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		m += 1
	}
	return m
}

func fun2(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	return m
}

func fun3(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}

func fun4(n int) int {
	m := 0
	i := 1
	for i < n {
		m += 1
		i = i * 2
	}
	return m
}

func fun5(n int) int {
	m := 0
	i := n
	for i > 0 {
		m += 1
		i = i / 2
	}
	return m
}

func fun6(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

func fun7(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m += 1
		}
	}
	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			m += 1
		}
	}
	return m
}

func fun8(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		sq := math.Sqrt(float64(n))
		for j := 0; j < int(sq); j++ {
			m += 1
		}
	}
	return m
}

func fun9(n int) int {
	m := 0
	for i := n; i > 0; i /= 2 {
		for j := 0; j < i; j++ {
			m += 1
		}
	}
	return m
}

func fun10(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			m += 1
		}
	}
	return m
}

func fun11(n int) int {
	m := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				m += 1
			}
		}
	}
	return m
}

func fun12(n int) int {
	j := 0
	m := 0
	for i := 0; i < n; i++ {
		for ; j < n; j++ {
			m += 1
		}
	}
	return m
}

func fun13(n int) int {
	m := 0
	for i := 1; i <= n; i *= 2 {
		for j := 0; j <= i; j++ {
			m += 1
		}
	}
	return m
}

func main() {
	fmt.Println("N = 100, Number of instructions O(n) ::", fun1(100))
	fmt.Println("N = 100, Number of instructions O(n^2) ::", fun2(100))
	fmt.Println("N = 100, Number of instructions O(n^2) ::", fun3(100))
	fmt.Println("N = 100, Number of instructions O(log(n)) ::", fun4(100))
	fmt.Println("N = 100, Number of instructions O(log(n)) ::", fun5(100))
	fmt.Println("N = 100, Number of instructions O(n^3) ::", fun6(100))
	fmt.Println("N = 100, Number of instructions O(n^2) ::", fun7(100))
	fmt.Println("N = 100, Number of instructions O(n^(3/2)) ::", fun8(100))
	fmt.Println("N = 100, Number of instructions O(n) ::", fun9(100))
	fmt.Println("N = 100, Number of instructions O(n^2) ::", fun10(100))
	fmt.Println("N = 100, Number of instructions O(n^3) ::", fun11(100))
	fmt.Println("N = 100, Number of instructions O(n) ::", fun12(100))
	fmt.Println("N = 100, Number of instructions O(n) ::", fun13(100))

}

/*
N = 100, Number of instructions O(n) :: 100
N = 100, Number of instructions O(n^2) :: 10000
N = 100, Number of instructions O(n^2) :: 4950
N = 100, Number of instructions O(log(n)) :: 7
N = 100, Number of instructions O(log(n)) :: 7
N = 100, Number of instructions O(n^3) :: 1000000
N = 100, Number of instructions O(n^2) :: 20000
N = 100, Number of instructions O(n^(3/2)) :: 1000
N = 100, Number of instructions O(n) :: 197
N = 100, Number of instructions O(n^2) :: 4950
N = 100, Number of instructions O(n^3) :: 166650
N = 100, Number of instructions O(n) :: 100
N = 100, Number of instructions O(n) :: 134
*/
