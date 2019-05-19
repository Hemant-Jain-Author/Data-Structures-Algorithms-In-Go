package main

import "fmt"
import "github.com/golang-collections/collections/queue"


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

func main(){
    fmt.Println("3 to ", 13 ," : ", convertXY(3, 13))
    
}
    