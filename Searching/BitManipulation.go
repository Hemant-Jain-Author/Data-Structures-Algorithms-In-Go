package main

import "fmt"

func AndEx(a int, b int) int {
	return a & b
}

func BitReversalEx(a int) int {
	return -a - 1
}

func CountBits(a int) int {
	count := 0
	for a > 0 {
		count += 1
		a = a & (a - 1)
	}
	return count
}

func IsPowerOf2(a int) bool {
	if a&(a-1) == 0 {
		return true
	} else {
		return false
	}
}

func KthBitCheck(a int, k int) bool {
	return a&(1<<(k-1)) > 0
}

func KthBitReset(a int, k int) int {
	return a & ^(1 << (k - 1))
}

func KthBitSet(a int, k int) int {
	return a | 1<<(k-1)
}

func KthBitToggle(a int, k int) int {
	return a ^ 1<<(k-1)
}

func LeftShiftEx(a int) int {
	return a << 1
}

func OrEx(a int, b int) int {
	return a | b
}

func ResetRightMostBit(a int) int {
	return a & (a - 1)
}

func RightMostBit(a int) int {
	return a & -a
}

func RightShiftEx(a int) int {
	return a >> 1
}

func TwoComplementEx(a int) int {
	return -a
}

func XorEx(a int, b int) int {
	return a ^ b
}

func main() {
	a := 4
	b := 8
	fmt.Println(AndEx(a, b))
	fmt.Println(OrEx(a, b))
	fmt.Println(XorEx(a, b))
	fmt.Println(LeftShiftEx(a))  // multiply by 2
	fmt.Println(RightShiftEx(a)) // divide by 2
	fmt.Println(BitReversalEx(a))
	fmt.Println(TwoComplementEx(a))
	fmt.Println(KthBitCheck(a, 3))
	fmt.Println(KthBitSet(a, 2))
	fmt.Println(KthBitReset(a, 3))
	fmt.Println(KthBitToggle(a, 3))
	fmt.Println(RightMostBit(a))
	fmt.Println(ResetRightMostBit(a))
	fmt.Println(IsPowerOf2(a))
	for i := 0; i < 10; i++ {
		fmt.Println(i, " bit count : ", CountBits(i))
	}
}
