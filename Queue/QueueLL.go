package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type QueueLinkedList struct {
    tail *Node
    size int
}

func (q *QueueLinkedList) Size() int {
    return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
    return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }
    return q.tail.next.value
}

func (q *QueueLinkedList) Add(value int) {
    temp := &Node{value, nil}
    if q.tail == nil {
        q.tail = temp
		temp.next = temp
    } else {
        temp.next = q.tail.next
		q.tail.next = temp
        q.tail = temp
    }
    q.size++
}

func (q *QueueLinkedList) Remove() int {
    if q.IsEmpty() {
        fmt.Println("QueueEmptyException")
        return 0
    }

    value := q.tail.next.value
	if q.tail ==  q.tail.next {
		q.tail = nil
	} else {
		q.tail.next = q.tail.next.next
	}
    q.size--
    return value
}

func (q *QueueLinkedList) Print() {
    if q.IsEmpty() {
		return
	}
	temp := q.tail.next
	for temp != q.tail {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func main() {
	que := new(QueueLinkedList)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println("IsEmpty :", que.IsEmpty());
    fmt.Println("Size :", que.Size());
    fmt.Println("Queue remove :", que.Remove());
    fmt.Println("Queue remove :", que.Remove());

}

/*
IsEmpty : false
Size : 3
Queue remove : 1
Queue remove : 2
*/
