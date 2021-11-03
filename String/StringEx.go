package main

import (
	"fmt"
)


func matchExpUtil(exp string, str string, i int, j int) bool {
	if i == len(exp) && j == len(str) {
		return true
	}
	if (i == len(exp) && j != len(str)) || (i != len(exp) && j == len(str)) {
		return false
	}
	if exp[i] == '?' || exp[i] == str[j] {
		return matchExpUtil(exp, str, i+1, j+1)
	}
	if exp[i] == '*' {
		return matchExpUtil(exp, str, i+1, j) || matchExpUtil(exp, str, i, j+1) || matchExpUtil(exp, str, i+1, j+1)
	}
	return false
}

func matchExp(exp string, str string) bool {
	return matchExpUtil(exp, str, 0, 0)
}

func main1() {
    fmt.Println(matchExp("hello*", "helloworld"))
    fmt.Println(matchExp("hello?d", "hellowd"))
    fmt.Println(matchExp("hello*hemant", "helloworldfsdfsdfdsfhemant"))
    fmt.Println(matchExp("*hemantj", "helloworldfsdfsdfdsfhemant"))
}
/*
true
true
true
false
*/
func match(source string, pattern string) bool {
	iSource := 0
	iPattern := 0
	sourceLen := len(source)
	patternLen := len(pattern)
	for iSource = 0; iSource < sourceLen; iSource++ {
		if source[iSource] == pattern[iPattern] {
			iPattern++
		}
		if iPattern == patternLen {
			return true
		}
	}
	return false
}

func main2() {
    fmt.Println(match("hellofskdlfjsdlfjsldjflksdworld", "helloworld"))
    fmt.Println(match("hellod", "hellowd"))
    fmt.Println(match("hello*xxxxxxxxxxhemantxxxxxxxxxxxx", "hellnt"))
    fmt.Println()
}

/*
true
false
true
*/

func isPrime(n int) bool {
	answer := false

	if n > 1 {
		answer = true
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			answer = false
			break
		}
	}
	return answer
}

func main3() {
    fmt.Print("Prime numbers under 10 :: ")
    for i := 0; i < 10; i++ {
        if (isPrime(i)){
            fmt.Print(i, " ")
        }
    }
    fmt.Println()
}

/*
Prime numbers under 10 :: 2 3 5 7 
*/

func isUniqueChar(str string) bool {
	mp := make(map[byte]int)
	size := len(str)
	for i := 0; i < size; i++ {
		c := str[i]
		if mp[c] != 0 {
			fmt.Println("Duplicate detected!")
			return false
		}
		mp[c] = 1
	}
	fmt.Println("No duplicate detected!")
	return true
}

func main4() {
    isUniqueChar("aple")
    isUniqueChar("apple")
}

/*
No duplicate detected!
Duplicate detected!
*/

func isPermutation(s1 string, s2 string) bool {
	count := make(map[byte]int)
	length := len(s1)
	if len(s2) != length {
		fmt.Println(s1, "&", s2, "are not permutation")
		return false
	}

	for i := 0; i < length; i++ {
		ch := s1[i]
		count[ch]++
		ch = s2[i]
		count[ch]--
	}
	for i := 0; i < length; i++ {
		ch := s1[i]
		if count[ch] != 0 {
			fmt.Println(s1, "&", s2, "are not permutation")
			return false
		}
	}
	fmt.Println(s1, "&", s2, "are permutation")
	return true
}

func main5() {
    fmt.Println("isPermutation :", isPermutation("apple", "plepa"))
    fmt.Println("isPermutation :", isPermutation("appleb", "plepaa"))
}

/*
apple & plepa are permutation
isPermutation : true
appleb & plepaa are not permutation
isPermutation : false
*/

func isPalindrome(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j && str[i] == str[j] {
		i++
		j--
	}
	if i < j {
		fmt.Println("String is not a Palindrome")
		return false
	}
	fmt.Println("String is a Palindrome")
	return true
}

func main6() {
    isPalindrome("hello")
    isPalindrome("eoloe")
}

/*
String is not a Palindrome
String is a Palindrome
*/

func pow(x int, n int) int {
	var value int
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		value = pow(x, n/2)
		return (value * value)
	} else {
		value = pow(x, n/2)
		return x * value * value
	}
}

func main7() {
    fmt.Println(pow(5, 2))
}

/*
25
*/

func strcmp(a string, b string) int {
	index := 0
	len1 := len(a)
	len2 := len(b)
	minlen := len1
	if len1 > len2 {
		minlen = len2
	}

	for index < minlen && a[index] == b[index] {
		index++
	}

	if index == len1 && index == len2 {
		return 0
	} else if len1 == index {
		return -1
	} else if len2 == index {
		return 1
	}
	return (int)(a[index]) - (int)(b[index])
}

func main8() {
    fmt.Println("StrCmp returns :", strcmp("aba", "aas"))
}

/*
StrCmp returns : 1
*/

func reverseString(a string) string {
	chars := []rune(a)
	reverseStringUtil(chars)
	return string(chars)
}

func reverseStringUtil(a []rune) {
	lower := 0
	upper := len(a) - 1
	for lower < upper {
		a[lower], a[upper] = a[upper], a[lower]
		lower++
		upper--
	}
}

func reverseStringRange(a []rune, lower int, upper int) {
	for lower < upper {
		a[lower], a[upper] = a[upper], a[lower]
		lower++
		upper--
	}
}

func reverseWords(str string) string {
	length := len(str)
	upper := -1
	lower := 0
	arr := []rune(str)
	for i := 0; i < length; i++ {
		if arr[i] == ' ' {
			reverseStringRange(arr, lower, upper)
			lower = i + 1
			upper = i
		} else {
			upper++
		}
	}
	reverseStringRange(arr, lower, upper)
	reverseStringRange(arr, 0, length-1)
	return string(arr)
}

func main9() {
    fmt.Println(reverseString("apple"))
    fmt.Println(reverseWords("hello world"))
}

/*
elppa
world hello
*/

func printAnagram(a string) {
	n := len(a)
	printAnagramUtil([]rune(a), n, n)
}

func printAnagramUtil(a []rune, max int, n int) {
	if max == 1 {
		fmt.Println(string(a))
	}
	for i := -1; i < max-1; i++ {
		if i != -1 {
			a[i], a[max-1] = a[max-1], a[i]
		}
		printAnagramUtil(a, max-1, n)
		if i != -1 {
			a[i], a[max-1] = a[max-1], a[i]
		}
	}
}

func main10() {
    printAnagram("123")
}

/*
123
213
321
231
132
312
*/

func shuffle(arr string) string {
	ar := []rune(arr)
	n := len(ar) / 2
	count := 0
	k := 1
	var temp rune
	for i := 1; i < n; i = i + 2 {
		k = i
		temp = ar[i]
		for true {
			k = (2 * k) % (2*n - 1)
			temp, ar[k] = ar[k], temp
			count++
			if i == k {
				break
			}
		}
		if count == (2*n - 2) {
			break
		}
	}
	return string(ar)
}

func main11() {
    fmt.Println(shuffle("ABCDE12345"))
}

/*
A1B2C3D4E5
*/

func main() {
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
}