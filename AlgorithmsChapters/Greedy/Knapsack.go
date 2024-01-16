package main

import (
	"fmt"
	"sort"
)

type Item struct {
	wt, cost int
	density  float64
}

func NewItem(w, c int) *Item {
	itm := new(Item)
	itm.wt = w
	itm.cost = c
	itm.density = float64(c) / float64(w)
	return itm
}

// Approximate solution.
func GetMaxCostGreedy(wt []int, cost []int, capacity int) int {
	totalCost := 0
	n := len(wt)
	itemList := make([]*Item, n)
	for i := 0; i < n; i++ {
		itemList[i] = NewItem(wt[i], cost[i])
	}
	sort.Slice(itemList, func(i, j int) bool {
		return itemList[i].density > itemList[j].density
	})

	for i := 0; i < n && capacity > 0; i++ {
		if capacity-itemList[i].wt >= 0 {
			capacity -= itemList[i].wt
			totalCost += itemList[i].cost
		}
	}
	return totalCost
}

func main() {
	wt := []int{10, 40, 20, 30}
	cost := []int{60, 40, 90, 120}
	capacity := 50
	maxCost := GetMaxCostGreedy(wt, cost, capacity)
	fmt.Println("Maximum cost obtained:", maxCost)
}

/*
Maximum cost obtained: 150
*/
