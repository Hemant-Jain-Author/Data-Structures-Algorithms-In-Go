package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IsBalancedParenthesis(expn string) bool {
	stk := new(Stack)
	for _, ch := range expn {
		switch ch {
		case '{', '[', '(':
			stk.Push(ch)
		case '}':
			val := stk.Pop()
			if val != '{' {
				return false
			}
		case ']':
			val := stk.Pop()
			if val != '[' {
				return false
			}
		case ')':
			val := stk.Pop()
			if val != '(' {
				return false
			}
		}
	}
	return stk.IsEmpty()
}

func main111() {
	expn := "{()}[]"
	value := IsBalancedParenthesis(expn)
	fmt.Println("Given Expn:", expn)
	fmt.Println("Result aft1er isParenthesisMatched:", value)
}

func insertAtBottom(stk *Stack, value interface{}) {
	if stk.IsEmpty() {
		stk.Push(value)
	} else {
		out := stk.Pop()
		insertAtBottom(stk, value)
		stk.Push(out)
	}
}

func reverseStack(stk *Stack) {
	if stk.IsEmpty() {
		return
	}
	value := stk.Pop()
	reverseStack(stk)
	insertAtBottom(stk, value)
}

func postfixEvaluate(expn string) int {
	stk := new(Stack)
	str := strings.Split(expn, " ")
	for _, tkn := range str {
		value, err := strconv.Atoi(tkn)
		if err == nil {
			stk.Push(value)
		} else {
			num1 := stk.Pop().(int)
			num2 := stk.Pop().(int)

			switch tkn {
			case "+":
				stk.Push(num1 + num2)
			case "-":
				stk.Push(num1 - num2)
			case "*":
				stk.Push(num1 * num2)
			case "/":
				stk.Push(num1 / num2)
			}
		}
	}
	return stk.Pop().(int)
}

func main222() {
	expn := "6 5 2 3 + 8 * + 3 + *"
	value := postfixEvaluate(expn)
	fmt.Println("Given Postfix Expn: ", expn)
	fmt.Println("Result after Evaluation: ", value)
}

func precedence(x rune) int {
	if x == '(' {
		return (0)
	}
	if x == '+' || x == '-' {
		return (1)
	}
	if x == '*' || x == '/' || x == '%' {
		return (2)
	}
	if x == '^' {
		return (3)
	}
	return (4)
}

func InfixToPostfix(expn string) string {
	fmt.Println(expn)
	stk := new(Stack)
	output := ""

	for _, ch := range expn {
		if ch <= '9' && ch >= '0' {
			output = output + string(ch)
		} else {
			switch ch {
			case '+', '-', '*', '/', '%', '^':
				for stk.IsEmpty() == false && precedence(ch) <= precedence(stk.Top().(rune)) {
					out := stk.Pop().(rune)
					output = output + " " + string(out)
				}
				stk.Push(ch)
				output = output + " "
			case '(':
				stk.Push(ch)
			case ')':
				for stk.IsEmpty() == false && stk.Top() != '(' {
					out := stk.Pop().(rune)
					output = output + " " + string(out) + " "
				}
				if stk.IsEmpty() == false && stk.Top() == '(' {
					stk.Pop()
				}
			}
		}
	}

	for stk.IsEmpty() == false {
		out := stk.Pop().(rune)
		output = output + string(out) + " "
	}
	return output
}

func main333() {
	expn := "10+((3))*5/(16-4)"
	value := InfixToPostfix(expn)
	fmt.Println("Infix Expn: ", expn)
	fmt.Println("Postfix Expn: ", value)
}

func InfixToPrefix(expn string) string {
	expn = reverseString(expn)
	expn = replaceParanthesis(expn)
	expn = InfixToPostfix(expn)
	expn = reverseString(expn)
	return expn
}
func reverseString(in string) string {
	expn := []rune(in)
	lower := 0
	upper := len(expn) - 1
	for lower < upper {
		expn[lower], expn[upper] = expn[upper], expn[lower]
		lower++
		upper--
	}
	return string(expn)
}

func replaceParanthesis(str string) string {
	a := []rune(str)
	lower := 0
	upper := len(a) - 1
	for lower <= upper {
		if a[lower] == '(' {
			a[lower] = ')'
		} else if a[lower] == ')' {
			a[lower] = '('
		}
		lower++
	}
	return string(a)
}

func main4444() {
	expn := "10+((3))*5/(16-4)"
	value := InfixToPrefix(expn)
	fmt.Println("Infix Expn: ", expn)
	fmt.Println("Prefix Expn: ", value)
}

func StockSpanRange(arr []int) []int {
	n := len(arr)
	SR := make([]int, n)

	SR[0] = 1
	for i := 1; i < len(arr); i++ {
		SR[i] = 1
		for j := i - 1; (j >= 0) && (arr[i] >= arr[j]); j-- {
			SR[i]++
		}
	}
	return SR
}

func StockSpanRange2(arr []int) []int {
	stk := new(Stack)
	n := len(arr)
	SR := make([]int, n)
	stk.Push(0)
	SR[0] = 1
	for i := 1; i < len(arr); i++ {
		for !stk.IsEmpty() && arr[stk.Top().(int)] <= arr[i] {
			stk.Pop()
		}
		if stk.IsEmpty() {
			SR[i] = (i + 1)
		} else {
			SR[i] = (i - stk.Top().(int))
		}
		stk.Push(i)
	}
	return SR
}

func main666() {
	stock := []int{4, 6, 8, 12, 2, 1, 7, 8}
	fmt.Println(stock)
	fmt.Println(StockSpanRange(stock))
	fmt.Println(StockSpanRange2(stock))
}

func GetMaxArea(arr []int) int {
	size := len(arr)
	maxArea := -1
	currArea := 0
	minHeight := 0
	for i := 1; i < size; i++ {
		minHeight = arr[i]
		for j := i - 1; j >= 0; j-- {
			if minHeight > arr[j] {
				minHeight = arr[j]
			}
			currArea = minHeight * (i - j + 1)
			if maxArea < currArea {
				maxArea = currArea
			}
		}
	}
	return maxArea
}

func GetMaxArea2(arr []int) int {
	size := len(arr)
	stk := new(Stack)
	maxArea := 0
	Top := 0
	TopArea := 0
	i := 0
	for i < size {
		for (i < size) && (stk.IsEmpty() || arr[stk.Top().(int)] <= arr[i]) {
			stk.Push(i)
			i++
		}
		for !stk.IsEmpty() && (i == size || arr[stk.Top().(int)] > arr[i]) {
			Top = stk.Top().(int)
			stk.Pop()
			if stk.IsEmpty() {
				TopArea = arr[Top] * i
			} else {
				TopArea = arr[Top] * (i - stk.Top().(int) - 1)
			}

			if maxArea < TopArea {
				maxArea = TopArea
			}
		}
	}
	return maxArea
}

func main77() {
	stock := []int{4, 6, 8, 12, 2, 1, 7, 8}
	fmt.Println(stock)
	fmt.Println(GetMaxArea(stock))
	fmt.Println(GetMaxArea2(stock))
}
