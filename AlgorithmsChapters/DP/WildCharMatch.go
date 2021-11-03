package main

import "fmt"

func main() {
	fmt.Println(WildCharMatchExp("*llo,?World?", "Hello, World!"))
	fmt.Println(WildCharMatchExpDP("*llo,?World?", "Hello, World!"))
}

func WildCharMatchExp(exp string, str string) bool {
	return wildCharMatchExpUtil(exp, str, 0, 0)
}

func wildCharMatchExpUtil(exp string, str string, m int, n int) bool {
	if m == len(exp) && (n == len(str) || exp[m-1] == '*') {
		return true
	}
	if m == len(exp) && n != len(str) || m != len(exp) && n == len(str) {
		return false
	}
	if exp[m] == '?' || exp[m] == str[n] {
		return wildCharMatchExpUtil(exp, str, m+1, n+1)
	}
	if exp[m] == '*' {
		return wildCharMatchExpUtil(exp, str, m+1, n) || wildCharMatchExpUtil(exp, str, m, n+1)
	}
	return false
}

func WildCharMatchExpDP(exp string, str string) bool {
	return wildCharMatchExpUtilDP(exp, str, len(exp), len(str))
}

func wildCharMatchExpUtilDP(exp string, str string, m int, n int) bool {
	lookup := make([][]bool, m+1)
	for i := range lookup {
		lookup[i] = make([]bool, n+1)
	}

	// empty exp and empty str match.
	lookup[0][0] = true

	// 0 row will remain all false. empty exp can't match any str.
	// '*' can match with empty string, column 0 update.
	for i := 1; i <= m; i++ {
		if exp[i-1] == '*' {
			lookup[i][0] = lookup[i-1][0]
		} else {
			break
		}
	}

	// Fill the table in bottom-up fashion
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if exp[i-1] == '*' {
				// If we see a '*' in pattern:
				// 1) We ignore '*' character and consider
				// next character in the pattern.
				// 2) We ignore one character in the input str
				// and consider next character.
				lookup[i][j] = lookup[i-1][j] || lookup[i][j-1]
			} else if exp[i-1] == '?' || str[j-1] == exp[i-1] {
				// Condition when both the pattern and input string
				// have same character. Also '?' match with all the
				// characters.
				lookup[i][j] = lookup[i-1][j-1]
			} else {
				// If characters don't match
				lookup[i][j] = false
			}
		}
	}
	return lookup[m][n]
}
