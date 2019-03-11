package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
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

func main1() {
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

func main2() {
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

func main3() {
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

func main4() {
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

func main5() {
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

func main6() {
	stock := []int{4, 6, 8, 12, 2, 1, 7, 8}
	fmt.Println(stock)
	fmt.Println(GetMaxArea(stock))
	fmt.Println(GetMaxArea2(stock))
}

func sortedInsert(stk *Stack, element int) {
	var temp int
	if (stk.Length() == 0 || element > stk.Top().(int)) {
		stk.Push(element)
	} else {
		temp = stk.Pop().(int)
		sortedInsert(stk, element)
		stk.Push(temp)
	}
}

func sortStack(stk *Stack) {
	var temp int
	if (stk.Length() > 0) {
		temp = stk.Pop().(int)
		sortStack(stk)
		stk.Push(temp)
	}
}

func sortStack2(stk *Stack) {
	var temp int
	stk2 := new(Stack)
	for (stk.Length() > 0) {
		temp = stk.Pop().(int)
		for ((stk.Length() > 0) && (stk2.Top().(int) < temp)) {
			stk.Push(stk2.Pop().(int))
		}
		stk2.Push(temp)
	}
	for (stk2.Length() > 0) {
		stk.Push(stk2.Pop().(int))
	}
}

func bottomInsert(stk *Stack, element int) {
	var temp int
	if (stk.Length() == 0) {
		stk.Push(element)
	} else {
		temp = stk.Pop().(int)
		bottomInsert(stk, element)
		stk.Push(temp)
	}
}

func reverseStack2(stk *Stack) {
	que := new(Queue)
	for (stk.Length() > 0) {
		que.Enqueue(stk.Pop())
	}

	for (que.Length() != 0) {
		stk.Push(que.Dequeue())
	}
}

func reverseKElementInStack(stk *Stack, k int) {
	que := new(Queue)
	i := 0
	for (stk.Length() > 0 && i < k) {
		que.Enqueue(stk.Pop())
		i++
	}
	for (que.Length() != 0) {
		stk.Push(que.Dequeue())
	}
}

func reverseQueue(que *Queue) {
	stk := new(Stack)
	for (que.Length() != 0) {
		stk.Push(que.Dequeue())
	}

	for (stk.Length() > 0) {
		que.Enqueue(stk.Pop())
	}
}

func reverseKElementInQueue(que *Queue, k int) {
	stk := new(Stack)
	i := 0
	var diff, temp int
	for (que.Length() != 0 && i < k) {
		stk.Push(que.Dequeue())
		i++
	}
	for (stk.Length() > 0) {
		que.Enqueue(stk.Pop())
	}
	diff = que.Length() - k
	for (diff > 0) {
		temp = que.Dequeue().(int)
		que.Enqueue(temp)
		diff -= 1
	}
}

func main7() {
	stk := new(Stack)
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	stk.Push(4)
	stk.Push(5)
	stk.Print()
	fmt.Println()
}

func main8() {
	stk := new(Stack)
	stk.Push(-2)
	stk.Push(13)
	stk.Push(16)
	stk.Push(-6)
	stk.Push(40)
	stk.Print()
	fmt.Println()

	reverseStack2(stk)
	stk.Print()
	fmt.Println()

	reverseKElementInStack(stk, 2)
	stk.Print()
	fmt.Println()
	/*
		* System.out.println(stk) sortStack2(stk) System.out.println(stk)
		*/
	que := new(Queue)
	que.Enqueue(1)
	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(4)
	que.Enqueue(5)
	que.Enqueue(6)
	que.Print()
	fmt.Println()

	reverseQueue(que)
	que.Print()
	fmt.Println()

	reverseKElementInQueue(que, 2)
	que.Print()
	fmt.Println()
}

func maxDepthParenthesis(expn string, size int) int {
	stk := new(Stack)
	maxDepth := 0
	depth := 0
	var ch byte

	for i := 0; i < size; i++ {
		ch = expn[i]

		if (ch == '(') {
			stk.Push(ch)
			depth += 1
		} else  if (ch == ')') {
			stk.Pop()
			depth -= 1
		}
		if (depth > maxDepth) {
			maxDepth = depth
		}
	}
	return maxDepth
}

func maxDepthParenthesis2(expn string, size int) int {
	maxDepth := 0
	depth := 0
	var ch byte
	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == '(') {
			depth += 1
		} else  if (ch == ')') {
			depth -= 1
		}

		if (depth > maxDepth) {
			maxDepth = depth
		}
	}
	return maxDepth
}

func main9() {
	expn := "((((A)))((((BBB()))))()()()())"
	size := len(expn)
	value :=maxDepthParenthesis(expn, size)
	value2 :=maxDepthParenthesis2(expn, size)

	fmt.Println("Given expn " , expn)
	fmt.Println("Max depth parenthesis is " , value)
	fmt.Println("Max depth parenthesis is " , value2)
}

func longestContBalParen(str string, size int) int {
	stk := new(Stack)
	stk.Push(-1)
	length := 0

	for i := 0; i < size; i++ {

		if (str[i] == '(') {
			stk.Push(i)
		} else  // string[i] == ')'
		{
			stk.Pop()
			if (stk.Length() != 0) {
				if length < i - stk.Top().(int){
					length = i - stk.Top().(int)
				}
			} else {
				stk.Push(i)
			}
		}
	}
	return length
}

func main10() {
	expn := "())((()))(())()(()"
	size := len(expn)
	value :=longestContBalParen(expn, size)
	fmt.Println("longestContBalParen " , value)
}

func reverseParenthesis(expn string, size int) int {
	stk := new(Stack)
	openCount := 0
	closeCount := 0
	var ch byte

	if (size % 2 == 1) {
		fmt.Println("Invalid odd length " , size)
		return -1
	}
	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == '(') {
			stk.Push(ch)
		} else  if (ch == ')') {
			if (stk.Length() != 0 && stk.Top().(byte) == '(') {
				stk.Pop()
			} else {
				stk.Push(')')
			}
		}
	}
	for (stk.Length() != 0) {
		if (stk.Pop() == '(') {
			openCount += 1
		} else {
			closeCount += 1
		}
	}
	reversal := int(math.Ceil(float64(openCount) / 2.0)) + int(math.Ceil(float64(closeCount) / 2.0))
	return reversal
}

func main11() {
	//expn := "())((()))(())()(()()()()))"
	expn2 := ")(())((("
	size := len(expn2)
	value :=reverseParenthesis(expn2, size)
	fmt.Println("Given expn : " , expn2)
	fmt.Println("reverse Parenthesis is : " , value)
}

func findDuplicateParenthesis(expn string, size int) bool {
	stk := new(Stack)
	var ch byte
	var count int

	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == ')') {
			count = 0
			for (stk.Length() != 0 && stk.Top().(byte) != '(') {
				stk.Pop()
				count += 1
			}
			if (count <= 1) {
				return true
			}
		} else {
			stk.Push(ch)
		}
	}
	return false
}

func main12() {
	// expn = "(((a+(b))+(c+d)))"
	// expn = "(b)"
	expn := "(((a+b))+c)"
	fmt.Println("Given expn : " , expn)
	size := len(expn)
	value := findDuplicateParenthesis(expn, size)
	fmt.Println("Duplicate Found : " , value)
}

func printParenthesisNumber(expn string, size int) {
	var ch byte
	stk := new(Stack)
	output := ""
	var count int = 1
	for i := 0; i < size; i++ {
		ch = expn[i]
		if (ch == '(') {
			stk.Push(count)
			output += fmt.Sprintf("%v",count)
			count += 1
		} else if (ch == ')') {
			output += fmt.Sprintf("%v",stk.Pop().(int))
		}
	}
	fmt.Println("Parenthesis Count : " , output)
}

func main13() {
	expn1 := "(((a+(b))+(c+d)))"
	expn2 := "(((a+b))+c)((("
	size := len(expn1)
	fmt.Println("Given expn " , expn1)
	printParenthesisNumber(expn1, size)
	size = len(expn2)
	fmt.Println("Given expn " , expn2)
	printParenthesisNumber(expn2, size)
}

func nextLargerElement(arr []int, size int) {
	output := make([]int, size)
	outIndex := 0
	var next int

	for i := 0; i < size; i++ {
		next = -1
		for j:= i + 1; j < size; j++ {
			if (arr[i] < arr[j]) {
				next = arr[j]
				break
			}
		}
		output[outIndex] = next
		outIndex++
	}
	for _, v := range output {
		fmt.Print(v , " ")
	}
	fmt.Println()
}

func nextLargerElement2(arr []int, size int) {
	stk := new(Stack)
	// output = [-1] * size
	output := make([]int, size)
	index := 0
	var curr int

	for i := 0; i < size; i++ {
		curr = arr[i]
		// stack always have values in decreasing order.
		for (stk.Length() > 0 && arr[stk.Top().(int)] <= curr) {
			index = stk.Pop().(int)
			output[index] = curr
		}
		stk.Push(i)
	}
	// index which dont have any next Larger.
	for (stk.Length() > 0) {
		index = stk.Pop().(int)
		output[index] = -1
	}
	for _, v := range output {
		fmt.Print(v , " ")
	}
	fmt.Println()
}

func nextSmallerElement(arr []int, size int) {
	stk := new(Stack)
	output := make([]int, size)
	var curr, index int
	for i := 0; i < size; i++ {
		curr = arr[i]
		// stack always have values in increasing order.
		for (stk.Length() > 0 && arr[stk.Top().(int)] > curr) {
			index = stk.Pop().(int)
			output[index] = curr
		}
		stk.Push(i)
	}
	// index which dont have any next Smaller.
	for (stk.Length() > 0) {
		index = stk.Pop().(int)
		output[index] = -1
	}
	for _, v := range output {
		fmt.Print(v , " ")
	}
	fmt.Println()
}

func main14() {
	arr := []int { 13, 21, 3, 6, 20, 3 }
	size := len(arr)
	nextLargerElement(arr, size)
	nextLargerElement2(arr, size)
	nextSmallerElement(arr, size)
}

func nextLargerElementCircular(arr []int, size int) {
	stk := new(Stack)
	var curr, index int
	output := make([]int, size)
	for i := 0; i < (2 * size - 1); i++ {
		curr = arr[i % size]
		// stack always have values in decreasing order.
		for (stk.Length() > 0 && arr[stk.Top().(int)] <= curr) {
			index = stk.Pop().(int)
			output[index] = curr
		}
		stk.Push(i % size)
	}
	// index which dont have any next Larger.
	for (stk.Length() > 0) {
		index = stk.Pop().(int)
		output[index] = -1
	}
	for _, v := range output {
		fmt.Print(v , " ")
	}
	fmt.Println()
}

func main15() {
	arr := []int { 6, 3, 9, 8, 10, 2, 1, 15, 7 }
	size := len(arr)
	nextLargerElementCircular(arr, size)
}

func RottenFruitUtil(arr [][]int, maxCol int, maxRow int, currCol int, currRow int, traversed [][]int, day int) { // Range check
	if (currCol < 0 || currCol >= maxCol || currRow < 0 || currRow >= maxRow) {
		return
	}
	// Traversable and rot if not already rotten.
	if (traversed[currCol][currRow] <= day || arr[currCol][currRow] == 0) {
		return
	}
	// Update rot time.
	traversed[currCol][currRow] = day
	// each line corresponding to 4 direction.
	RottenFruitUtil(arr, maxCol, maxRow, currCol - 1, currRow, traversed, day + 1)
	RottenFruitUtil(arr, maxCol, maxRow, currCol + 1, currRow, traversed, day + 1)
	RottenFruitUtil(arr, maxCol, maxRow, currCol, currRow + 1, traversed, day + 1)
	RottenFruitUtil(arr, maxCol, maxRow, currCol, currRow - 1, traversed, day + 1)
}

func RottenFruit(arr [][]int, maxCol int, maxRow int) int {
	traversed := make([][]int, maxRow)
    for i := range traversed {
        traversed[i] = make([]int, maxCol)
    }
	

	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			traversed[i][j] = 999999
		}
	}

	for i := 0; i < maxCol - 1; i++ {
		for j:= 0; j < maxRow - 1; j++ {
			if (arr[i][j] == 2) {
				RottenFruitUtil(arr, maxCol, maxRow, i, j, traversed, 0)
			}
		}
	}

	maxDay := 0
	for i := 0; i < maxCol - 1; i++ {
		for j:= 0; j < maxRow - 1; j++ {
			if (arr[i][j] == 1) {
				if (traversed[i][j] == 999999) {
					return -1
				}
				if (maxDay < traversed[i][j]) {
					maxDay = traversed[i][j]
				}
			}
		}
	}
	return maxDay
}

func main16() {
	arr := make([][]int, 5)
    arr[0] = []int{1, 0, 1, 1, 0} 
    arr[1] = []int{2, 1, 0, 1, 0}
    arr[2] = []int{0, 0, 0, 2, 1}
    arr[3] = []int{0, 2, 0, 0, 1}
    arr[4] = []int{1, 1, 0, 0, 1}
    //make([]int, col)
	fmt.Println(RottenFruit(arr, 5, 5))
}

func StepsOfKnightUtil(size int,currCol int, currRow int, traversed [][]int, dist int) {
	// Range check
	if (currCol < 0 || currCol >= size || currRow < 0 || currRow >= size) {
		return
	}

	// Traversable and rot if not already rotten.
	if (traversed[currCol][currRow] <= dist) {
		return
	}

	// Update rot time.
	traversed[currCol][currRow] = dist
	// each line corresponding to 4 direction.
	StepsOfKnightUtil(size, currCol - 2, currRow - 1, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol - 2, currRow + 1, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol + 2, currRow - 1, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol + 2, currRow + 1, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol - 1, currRow - 2, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol + 1, currRow - 2, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol - 1, currRow + 2, traversed, dist + 1)
	StepsOfKnightUtil(size, currCol + 1, currRow + 2, traversed, dist + 1)
}

func StepsOfKnight(size int, srcX int, srcY int, dstX int, dstY int) int {
	traversed := make([][]int, size)
    for i := range traversed {
        traversed[i] = make([]int, size)
    }

	for i := 0; i < size; i++ {
		for j:= 0; j < size; j++ {
			traversed[i][j] = 999999
		}
	}

	StepsOfKnightUtil(size, srcX - 1, srcY - 1, traversed, 0)
	retval := traversed[dstX - 1][dstY - 1]
	return retval
}

func main17() {
	fmt.Println(StepsOfKnight(20, 10, 10, 20, 20))
}

func DistNearestFillUtil(arr [][]int, maxCol int, maxRow int, currCol int, currRow int, traversed [][]int, dist int) { // Range check
	if (currCol < 0 || currCol >= maxCol || currRow < 0 || currRow >= maxRow) {
		return
	}
	// Traversable if their is a better distance.
	if (traversed[currCol][currRow] <= dist) {
		return
	}
	// Update distance.
	traversed[currCol][currRow] = dist
	// each line corresponding to 4 direction.
	DistNearestFillUtil(arr, maxCol, maxRow, currCol - 1, currRow, traversed, dist + 1)
	DistNearestFillUtil(arr, maxCol, maxRow, currCol + 1, currRow, traversed, dist + 1)
	DistNearestFillUtil(arr, maxCol, maxRow, currCol, currRow + 1, traversed, dist + 1)
	DistNearestFillUtil(arr, maxCol, maxRow, currCol, currRow - 1, traversed, dist + 1)
}

func DistNearestFill(arr [][]int, maxCol int, maxRow int) {
	traversed := make([][]int, maxRow)
    for i := range traversed {
        traversed[i] = make([]int, maxCol)
    }
	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			traversed[i][j] = 999999
		}
	}
	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			if (arr[i][j] == 1) {
				DistNearestFillUtil(arr, maxCol, maxRow, i, j, traversed, 0)
			}
		}
	}

	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			fmt.Print(" " , traversed[i][j])
		}
		fmt.Println()
	}
}

func main18() {
	arr := make([][]int, 5)
    arr[0] = []int{1, 0, 1, 1, 0} 
    arr[1] = []int{1, 1, 0, 1, 0}
    arr[2] = []int{0, 0, 0, 0, 1}
    arr[3] = []int{0, 0, 0, 0, 1}
    arr[4] = []int{0, 0, 0, 0, 1}    

	DistNearestFill(arr, 5, 5)
}

func findLargestIslandUtil(arr [][]int, maxCol int, maxRow int, currCol int, 
	currRow int, value int , traversed [][]int) int {
	if (currCol < 0 || currCol >= maxCol || currRow < 0 || currRow >= maxRow) {
		return 0
	}
	if (traversed[currCol][currRow] == 1 || arr[currCol][currRow] != value) {
		return 0
	}
	traversed[currCol][currRow] = 1
	// each call corresponding to 8 direction.
	return 1 + findLargestIslandUtil(arr, maxCol, maxRow, currCol - 1, currRow - 1, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol - 1, currRow, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol - 1, currRow + 1, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol, currRow - 1, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol, currRow + 1, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol + 1, currRow - 1, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol + 1, currRow, value, traversed) + findLargestIslandUtil(arr, maxCol, maxRow, currCol + 1, currRow + 1, value, traversed)
}

func findLargestIsland(arr [][]int, maxCol int, maxRow int) int {
	maxVal := 0
	currVal := 0

	traversed := make([][]int, maxRow)
    for i := range traversed {
        traversed[i] = make([]int, maxCol)
    }

	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			traversed[i][j] = 999999
		}
	}
	for i := 0; i < maxCol; i++ {
		for j:= 0; j < maxRow; j++ {
			currVal = findLargestIslandUtil(arr, maxCol, maxRow, i, j, arr[i][j], traversed)
			if (currVal > maxVal) {
				maxVal = currVal
		}
		}
	}
	return maxVal
}


func main19() {
	arr := make([][]int, 5)
    arr[0] = []int{1, 0, 1, 1, 0} 
    arr[1] = []int{1, 0, 0, 1, 0}
    arr[2] = []int{0, 1, 1, 1, 1}
    arr[3] = []int{0, 1, 0, 0, 0}
    arr[4] = []int{1, 1, 0, 0, 1}    

	fmt.Println("Largest Island : " , findLargestIsland(arr, 5, 5))
}

func isKnown(relation [][]int, a int , b int) bool {
	if (relation[a][b] == 1) {
		return true
	}
	return false
}

func findCelebrity(relation [][]int, count int) int {
	stk := new(Stack)
	first := 0 
	second := 0
	for i := 0; i < count; i++ {
		stk.Push(i)
	}
	first = stk.Pop().(int)
	for (stk.Length() != 0) {
		second = stk.Pop().(int)
		if (isKnown(relation, first, second)) {
			first = second
		}
	}
	for i := 0; i < count; i++ {
		if (first != i && isKnown(relation, first, i)) {
			return -1
		}
		if (first != i && isKnown(relation, i, first) == false) {
			return -1
		}
	}
	return first
}

func findCelebrity2(relation [][]int, count int) int {
	first := 0
	second := 1

	for i := 0; i < (count - 1); i++ {
		if (isKnown(relation, first, second)) {
			first = second
		}
		second = second + 1
	}
	for i := 0; i < count; i++ {
		if (first != i && isKnown(relation, first, i)) {
			return -1
		}
		if (first != i && isKnown(relation, i, first) == false) {
			return -1
		}
	}
	return first
}

func main20() {
	arr := make([][]int, 5)
    arr[0] = []int{1, 0, 1, 1, 0} 
    arr[1] = []int{1, 0, 0, 1, 0}
    arr[2] = []int{0, 0, 1, 1, 1}
    arr[3] = []int{0, 0, 0, 0, 0}
    arr[4] = []int{1, 1, 0, 1, 1}    

	fmt.Println("Celebrity : " , findCelebrity(arr, 5))
	fmt.Println("Celebrity : " , findCelebrity2(arr, 5))
}

func IsMinHeap(arr []int, size int) int {
	for i := 0; i <= (size - 2) / 2; i++ {
		if (2 * i + 1 < size) {
			if (arr[i] > arr[2 * i + 1]) {
				return 0
			}
		}
		if (2 * i + 2 < size) {
			if (arr[i] > arr[2 * i + 2]) {
				return 0
			}
		}
	}
	return 1
}

func IsMaxHeap(arr []int, size int) int {
	for i := 0; i <= (size - 2) / 2; i++ {
		if (2 * i + 1 < size) {
			if (arr[i] < arr[2 * i + 1]) {
				return 0
			}
		}
		if (2 * i + 2 < size) {
			if (arr[i] < arr[2 * i + 2]) {
				return 0
			}
		}
	}
	return 1
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
	main9()
	main10()
	main11()
	main12()
	main13()
	main14()
	main15()
	main16()
	main17()
	main18()
	main19()
	main20()
}

//**************************

type Stack struct {
	s []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.s = append(s.s, value)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

func (s *Stack) Top() interface{} {
	length := len(s.s)
	res := s.s[length-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

func (s *Stack) Length() int {
	length := len(s.s)
	return length
}

func (s *Stack) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}
//********************************

const capacity = 100

type Queue struct {
	size  int
	data  [capacity]interface{}
	front int
	back  int
}

func (q *Queue) Enqueue(value interface{}) {
	if q.size >= capacity {
		return
	}
	q.size++
	q.data[q.back] = value
	q.back = (q.back + 1) % (capacity - 1)
}

func (q *Queue) Dequeue() interface{} {
	var value interface{}
	if q.size <= 0 {
		return 0
	}
	q.size--
	value = q.data[q.front]
	q.front = (q.front + 1) % (capacity - 1)
	return value
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Length() int {
	return q.size
}

func (q *Queue) Print() {
	length := q.size
	f := q.front
	for i := 0; i < length; i++ {
		fmt.Print(q.data[(f + i)%capacity], " ")
	}
	fmt.Println()
}

//********************************