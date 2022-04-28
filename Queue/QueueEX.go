package main

import (
	"fmt"
	"math"
)

func main1() {
    q := new(Queue)
    for i := 1; i <= 5; i++ {
        q.Add(i)
    }
    for i := 1; i <= 5; i++ {
        fmt.Print(q.Remove(), " ")
    }
    fmt.Println()
}

func  CircularTour(arr [][2]int, n int)  int {
    que := new(Queue)
    nextPump := 0
    prevPump := 0
    count := 0
    petrol := 0

    for que.Len() != n {
        for petrol >= 0 && que.Len() != n {
            que.Add(nextPump)
            petrol += (arr[nextPump][0] - arr[nextPump][1])
            nextPump = (nextPump + 1) % n
        }
    
        for petrol < 0 && que.Len() > 0 {
            prevPump = que.Remove().(int)
            petrol -= (arr[prevPump][0] - arr[prevPump][1])
        }

        count += 1
        if count == n{
            return -1
        }
    }

    if petrol >= 0 {

        return que.Remove().(int)
    } else {

        return -1
    }
}

func main2(){
    tour := [][2]int{{8, 6}, {1, 4}, {7, 6}}
    fmt.Println("Circular Tour :" , CircularTour(tour, 3))    
}

/*
Circular Tour : 2
*/

func convertXY(src int, dst int) int{
    que := new(Queue)
    visited := make(map[int]int)
    a := []int{src, 0}
    que.Add(a)
    ok := false

    for que.Len() != 0{
        node := que.Remove().([]int)
        visited[node[0]] = 1
        value := node[0]
        steps := node[1]
        if value == dst {
            return steps
        }
        _, ok =  visited[value*2]
        if value < dst && !ok{
            que.Add([]int{value*2, steps+1})
        }
        _, ok =  visited[value-1]
        if value > 0 && !ok{
            que.Add([]int{value-1, steps+1})
        }
    }
    return -1
}

func main3(){
    fmt.Println("Steps to convert 2 to 7 :", convertXY(2, 7))    
}

/*
Steps to convert 2 to 7 : 3
*/

func maxSlidingWindows(arr []int, k int){
    size := len(arr)
    que := new(Queue)
    i := 0
    for i < size {
       // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] <= arr[i] {
            que.RemoveBack()
        }

        que.Add(i)
        // window of size k
        if i >= (k - 1) {
            fmt.Print(arr[que.Front().(int)], " ")
        }

        i += 1
    }
}

func main4() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    fmt.Print("Max Sliding Windows : ")
    maxSlidingWindows(arr, k)
    fmt.Println()
}

/*
Max Sliding Windows : 75 92 92 92 90 
*/

func minOfMaxSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    minVal := math.MaxInt32
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] <= arr[i] {
            que.RemoveBack()
        }
        que.Add(i)
        // window of size k
        if i >= (k - 1) && minVal > arr[que.Front().(int)] {
            minVal = arr[que.Front().(int)]
        }
        i += 1
    }
    fmt.Println("Min of max is : ", minVal)
}

func main5() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    minOfMaxSlidingWindows(arr, k)
}

//Min of max is :  75

func maxOfMinSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    maxVal := math.MinInt32
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k {
            que.Remove()
        }
        // Remove smaller values at left.
        for que.Len() > 0 && arr[que.Back().(int)] >= arr[i] {
            que.RemoveBack()
        }
        que.Add(i)
        // window of size k
        if i >= (k - 1) && maxVal < arr[que.Front().(int)] {
            maxVal = arr[que.Front().(int)]
        }
        i += 1
    }
    fmt.Println("Max of min is : ", maxVal)
}

func main6() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    maxOfMinSlidingWindows(arr, k)
}
/*
Max of min is :: 59
*/

func firstNegSlidingWindows(arr []int, k int) {
    size := len(arr)
    que := new(Queue)
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Len() > 0 && que.Front().(int) <= i - k{
            que.Remove()
        }
        if arr[i] < 0{
            que.Add(i)
        }
        // window of size k
        if i >= (k - 1) {
            if que.Len() > 0{
                fmt.Print(arr[que.Front().(int)], " ")
            } else {
                fmt.Print("NAN ") 
            }

        }
        i += 1
    }        
}

func main7() {
    arr3 := []int{3, -2, -6, 10, -14, 50, 14, 21, 11, -2, -11, 2, 3}
    firstNegSlidingWindows(arr3, 3)
}
/*
-2 -2 -6 -14 -14 NAN NAN -2 -2 -2 -11 
*/

func main(){
	main1()
	main2()
	main3()
	main4()
	main5()
	main6()
	main7()
}

type Queue struct {
	que []interface{}
}

func (q *Queue) Add(value interface{}) {
	q.que = append(q.que, value)
}

func (q *Queue) Remove() interface{} {
	n := len(q.que)
	value := q.que[0]
	q.que = q.que[1:n]
	return value
}

func (q *Queue) RemoveBack() interface{} {
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func (q *Queue) Front() interface{} {
	return q.que[0]
}

func (q *Queue) Back() interface{} {
	n := len(q.que)
	return q.que[n-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.que) == 0
}

func (q *Queue) Len() int {
	return len(q.que)
}

func (q Queue) Print() {
	fmt.Println(q.que)
}
