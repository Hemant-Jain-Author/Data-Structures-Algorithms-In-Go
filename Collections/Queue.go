package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func main() {
    que := queue.New()
    que.Enqueue(1)
    que.Enqueue(2)
    que.Enqueue(3)

	fmt.Println("Queue size :", que.Len());
	fmt.Println("Queue peek :", que.Peek().(int));
	fmt.Println("Queue remove :", que.Dequeue().(int));
	fmt.Println("Queue isEmpty :", que.Len() == 0);
}

/*
Queue size : 3
Queue peek : 1
Queue remove : 1
Queue isEmpty : false
*/