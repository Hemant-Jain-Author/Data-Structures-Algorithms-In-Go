package main

import (
	"fmt"
)

func BruteForceSearch(text string, pattern string) int {
	j := 0
	n := len(text)
	m := len(pattern)
	for i := 0; i <= n-m; i++ {
		j = 0
		for j < m && pattern[j] == text[i+j] {
			j++
		}
		if j == m {
			return i
		}
	}
	return -1
}

func RobinKarp(text string, pattern string) int {
	n := len(text)
	m := len(pattern)
	prime := 101
	powm := 1
	TextHash := 0
	PatternHash := 0
	i, j := 0, 0
	if m == 0 || m > n {
		return -1
	}

	for i = 0; i < m-1; i++ {
		powm = (powm << 1) % prime
	}

	for i = 0; i < m; i++ {
		PatternHash = ((PatternHash << 1) + (int)(pattern[i])) % prime
		TextHash = ((TextHash << 1) + (int)(text[i])) % prime
	}

	for i = 0; i <= (n - m); i++ {
		if TextHash == PatternHash {
			for j = 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					break
				}
			}
			if j == m {
				return i
			}
		}
		TextHash = (((TextHash - (int)(text[i])*powm) << 1) + (int)(text[i+m])) % prime
		if TextHash < 0 {
			TextHash = (TextHash + prime)
		}
	}
	return -1
}

func KMPPreprocess(pattern string, ShiftArr []int) {
	m := len(pattern)
	i := 0
	j := -1
	ShiftArr[i] = -1
	for i < m {
		for j >= 0 && pattern[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		ShiftArr[i] = j
	}
}

func KMP(text string, pattern string) int {
	i, j := 0, 0
	n := len(text)
	m := len(pattern)
	ShiftArr := make([]int, m+1)
	KMPPreprocess(pattern, ShiftArr)
	for i < n {
		for j >= 0 && text[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		if j == m {
			return (i - m)
		}
	}
	return -1
}

func KMPFindCount(text string, pattern string) int {
	i, j := 0, 0
	count := 0
	n := len(text)
	m := len(pattern)
	ShiftArr := make([]int, m+1)
	KMPPreprocess(pattern, ShiftArr)
	for i < n {
		for j >= 0 && text[i] != pattern[j] {
			j = ShiftArr[j]
		}
		i++
		j++
		if j == m {
			count++
			j = ShiftArr[j]
		}
	}
	return count
}

func main() {
	st1 := "hello, world!"
	st2 := "world"
	fmt.Println("Using BruteForceSearch pattern found at index : ", BruteForceSearch(st1, st2))
	fmt.Println("Using RobinKarp  pattern found at index : ", RobinKarp(st1, st2))
	fmt.Println("Using KMP  pattern found at index : ", KMP(st1, st2))
	
	str3 := "Only time will tell if we stand the test of time"
	fmt.Println("Frequency of 'time' is ", KMPFindCount(str3, "time"))
}
/*
Using BruteForceSearch pattern found at index :  7
Using RobinKarp  pattern found at index :  7
Using KMP  pattern found at index :  7
Frequency of 'time' is  2
*/