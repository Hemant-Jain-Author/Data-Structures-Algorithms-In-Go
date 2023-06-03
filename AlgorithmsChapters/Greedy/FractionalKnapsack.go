package main

import (
	"fmt"
	"sort"
)

type Items struct {
	wt, cost int
	density  float64
}

func NewItems(a, b int) *Items {
	p := new(Items)
	p.wt = a
	p.cost = b
	p.density = float64(b) / float64(a)
	return p
}

func GetMaxCostFractional(wt []int, cost []int, capacity int) float64 {
	totalCost := 0.0
	n := len(wt)
	itemList := make([]*Items, n)
	for i := 0; i < n; i++ {
		itemList[i] = NewItems(wt[i], cost[i])
	}

	// Sort items by density in decreasing order.
	sort.Slice(itemList, func(i, j int) bool {
		return itemList[i].density > itemList[j].density
	})

	for i := 0; i < n; i++ {
		if capacity-itemList[i].wt >= 0 {
			capacity -= itemList[i].wt
			totalCost += float64(itemList[i].cost)
		} else {
			totalCost += itemList[i].density * float64(capacity)
			break
		}
	}
	return totalCost
}

func main() {
	wt := []int{10, 40, 20, 30}
	cost := []int{60, 40, 90, 120}
	capacity := 50
	maxCost := GetMaxCostFractional(wt, cost, capacity)
	fmt.Println("Maximum cost obtained =", maxCost)
}

/*
Maximum cost obtained = 230
*/
