package main

/*
Their are N number of petrol pumps. Each petrol pump have 
some limited amount of petrol and they also their distance from each other is provided.
Find if their is a circular tour possible to visit all the petrol pumps.
*/

import "fmt"
import "github.com/golang-collections/collections/queue"

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

func main(){
    tour := [][2]int{{8, 6}, {1, 4}, {7, 6}}
    fmt.Println(CircularTour(tour, 3))    
}
