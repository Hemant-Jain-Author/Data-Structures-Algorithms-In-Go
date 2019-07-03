package main

import "fmt"
import "github.com/golang-collections/collections/queue"


type Node struct {
    value int
    next  *Node
    prev *Node
}

type Dequeue struct {
    head *Node
    tail *Node
    size int
}

func (q *Dequeue) Size() int {
    return q.size
}

func (q *Dequeue) IsEmpty() bool {
    return q.size == 0
}

func (q *Dequeue) PeekFront() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    return q.head.value
}

func (q *Dequeue) PeekBack() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    return q.tail.value
}

func (q *Dequeue) AddBack(value int) {
    temp := &Node{value, nil, nil}
    if q.head == nil {
        q.head = temp
        q.tail = temp
    } else {
        q.tail.next = temp
        temp.prev = q.tail
        q.tail = temp
    }
    q.size++
}

func (q *Dequeue) AddFront(value int) {
    temp := &Node{value, nil, nil}
    if q.head == nil {
        q.head = temp
        q.tail = temp
    } else {
        temp.next = q.head
        q.head.prev = temp
        q.head = temp
    }
    q.size++
}

func (q *Dequeue) RemoveFront() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    q.size--
    value := q.head.value
    q.head = q.head.next
    
    if q.head != nil {
        q.head.prev = nil
    } else {
        q.tail = nil
    }
    
    return value
}

func (q *Dequeue) RemoveBack() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    q.size--
    value := q.tail.value
    q.tail = q.tail.prev
    if q.tail != nil {
        q.tail.next = nil
    } else {
        q.head = nil
    }
    return value
}

func (q *Dequeue) Print() {
    temp := q.head
    for temp != nil {
        fmt.Print(temp.value, " ")
        temp = temp.next
    }
    fmt.Println()
}

func main1() {
    q := new(Dequeue)
    for i := 1; i <= 5; i++ {
        q.AddFront(i)
    }
    for i := 1; i <= 5; i++ {
        fmt.Print(q.RemoveBack(), " ")
    }
    fmt.Println()
}

func  CircularTour(arr [][2]int, n int)  int {
    que := queue.New()
    nextPump := 0
    prevPump := 0
    count := 0
    petrol := 0

    for que.Len() != n {
        for petrol >= 0 && que.Len() != n {
            que.Enqueue(nextPump)
            petrol += (arr[nextPump][0] - arr[nextPump][1])
            nextPump = (nextPump + 1) % n
        }
    
        for petrol < 0 && que.Len() > 0 {
            prevPump = que.Dequeue().(int)
            petrol -= (arr[prevPump][0] - arr[prevPump][1])
        }

        count += 1
        if count == n{
            return -1
        }
    }

    if petrol >= 0 {

        return que.Dequeue().(int)
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
    que := queue.New()
    visited := make(map[int]int)
    a := []int{src, 0}
    que.Enqueue(a)
    ok := false

    for que.Len() != 0{
        node := que.Dequeue().([]int)
        visited[node[0]] = 1
        value := node[0]
        steps := node[1]
        if value == dst {
            return steps
        }
        _, ok =  visited[value*2]
        if value < dst && !ok{
            que.Enqueue([]int{value*2, steps+1})
        }
        _, ok =  visited[value-1]
        if value > 0 && !ok{
            que.Enqueue([]int{value-1, steps+1})
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
    que := new(Dequeue)
    i := 0
    for i < size {
       // Remove out of range elements
        if que.Size() > 0 && que.PeekFront() <= i - k {
            que.RemoveFront()
        }
        // Remove smaller values at left.
        for que.Size() > 0 && arr[que.PeekBack()] <= arr[i] {
            que.RemoveBack()
        }

        que.AddBack(i)
        // window of size k
        if i >= (k - 1) {
            fmt.Print(arr[que.PeekFront()], " ")
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
    que := new(Dequeue)
    minVal := 999999
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Size() > 0 && que.PeekFront() <= i - k {
            que.RemoveFront()
        }
        // Remove smaller values at left.
        for que.Size() > 0 && arr[que.PeekBack()] <= arr[i] {
            que.RemoveBack()
        }
        que.AddBack(i)
        // window of size k
        if i >= (k - 1) && minVal > arr[que.PeekFront()] {
            minVal = arr[que.PeekFront()]
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
    que := new(Dequeue)
    maxVal := -999999
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Size() > 0 && que.PeekFront() <= i - k {
            que.RemoveFront()
        }
        // Remove smaller values at left.
        for que.Size() > 0 && arr[que.PeekBack()] >= arr[i] {
            que.RemoveBack()
        }
        que.AddBack(i)
        // window of size k
        if i >= (k - 1) && maxVal < arr[que.PeekFront()] {
            maxVal = arr[que.PeekFront()]
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
    que := new(Dequeue)
    i := 0
    for i < size {
        // Remove out of range elements
        if que.Size() > 0 && que.PeekFront() <= i - k{
            que.RemoveFront()
        }
        if arr[i] < 0{
            que.AddBack(i)
        }
        // window of size k
        if i >= (k - 1) {
            if que.Size() > 0{
                fmt.Print(arr[que.PeekFront()], " ")
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