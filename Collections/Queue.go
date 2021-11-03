package main

import "fmt"

type Queue []int

func (que *Queue) Add(x int) {
	*que = append(*que, x)
}

/*
In the Remove() function , the memory allocated for
the first element in the array is never returned so there is memory leak..
*/
func (que *Queue) Remove() int {
	n := len(*que)
	value := (*que)[0]
	*que = (*que)[1:n]
	return value
}

func (que *Queue) Front() int {
	return (*que)[0]
}

func (que Queue) Len() int {
	return len(que)
}

func (que Queue) IsEmpty() bool {
	return len(que) == 0
}

func (que Queue) Print() {
	fmt.Println(que)
}

func main() {
	que := &Queue{}
	for i := 0; i < 5; i++ {
		que.Add(i)
	}
	for que.IsEmpty() == false {
		fmt.Print(que.Remove(), " ")
	}
}

/*
0 1 2 3 4
*/
