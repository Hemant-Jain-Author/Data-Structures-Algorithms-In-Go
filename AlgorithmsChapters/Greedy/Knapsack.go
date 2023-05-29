package main

import (
	"fmt"
	"sort"
)

type Item struct {
	weight  int
	cost    int
	density float64
}

func NewItem(weight, cost int) *Item {
	return &Item{
		weight:  weight,
		cost:    cost,
		density: float64(cost) / float64(weight),
	}
}

func GetMaxCostGreedy(weights []int, costs []int, capacity int) int {
	totalCost := 0
	n := len(weights)
	items := make([]*Item, n)
	for i := 0; i < n; i++ {
		items[i] = NewItem(weights[i], costs[i])
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].density > items[j].density
	})

	for i := 0; i < n && capacity > 0; i++ {
		if capacity-items[i].weight >= 0 {
			capacity -= items[i].weight
			totalCost += items[i].cost
		}
	}
	return totalCost
}

func main() {
	weights := []int{10, 40, 20, 30}
	costs := []int{60, 40, 90, 120}
	capacity := 50
	maxCost := GetMaxCostGreedy(weights, costs, capacity)
	fmt.Println("Maximum cost obtained:", maxCost)
}
