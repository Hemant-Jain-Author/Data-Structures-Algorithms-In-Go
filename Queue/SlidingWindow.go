package main

import "fmt"

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

func mainDQ() {
    q := new(Dequeue)
    for i := 1; i <= 5; i++ {
        q.AddFront(i)
    }
    for i := 1; i <= 5; i++ {
        fmt.Println(q.RemoveBack())
    }
    q.Print()
}

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
            fmt.Println(arr[que.PeekFront()])
        }

        i += 1
    }

}


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
                fmt.Println(arr[que.PeekFront()])
            } else {
                fmt.Println("NAN") 
            }

        }
        i += 1
    }        
}

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

func main() {
    arr := []int{11, 2, 75, 92, 59, 90, 55}
    k := 3
    maxSlidingWindows(arr, k)
    minOfMaxSlidingWindows(arr, k)

    arr2 := []int{10, 20, 30, 50, 10, 70, 30}
    maxOfMinSlidingWindows(arr2, 2)

    arr3 := []int{13, -2, -6, 10, -14, 50, 14, 21}
    firstNegSlidingWindows(arr3, 3)
}