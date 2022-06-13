package main

import (
	"fmt"
)

func main(){
    var arr [10]int
    for  i  := 0; i < 10; i++ {
        arr[i] = i;
    }
    for  i  := 0; i < 10; i++ {
        fmt.Print(arr[i], " ");
    }
}

/*
0 1 2 3 4 5 6 7 8 9 
*/