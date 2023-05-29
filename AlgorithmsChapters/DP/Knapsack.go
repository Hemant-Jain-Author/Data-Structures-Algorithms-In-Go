package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	weight int
	cost   int
}

func NewItem(weight, cost int) *Item {
	return &Item{
		weight: weight,
		cost:   cost,
	}
}

func KS01UnboundedBU(items []*Item, capacity int) int {
	n := len(items)
	dp := make([]int, capacity+1)

	for w := 1; w <= capacity; w++ {
		for i := 0; i < n; i++ {
			itemWeight := items[i].weight
			itemCost := items[i].cost

			if itemWeight <= w {
				dp[w] = max(dp[w], dp[w-itemWeight]+itemCost)
			}
		}
	}

	return dp[capacity]
}

func GetMaxCost01(items []*Item, capacity int) int {
	n := len(items)
	return getMaxCost01Util(items, n, capacity)
}

func getMaxCost01Util(items []*Item, n, capacity int) int {
	// Base case
	if n == 0 || capacity == 0 {
		return 0
	}

	// Return the maximum of two cases:
	// (1) nth item is included
	// (2) nth item is not included
	first := 0
	itemWeight := items[n-1].weight
	itemCost := items[n-1].cost

	if itemWeight <= capacity {
		first = itemCost + getMaxCost01Util(items, n-1, capacity-itemWeight)
	}

	second := getMaxCost01Util(items, n-1, capacity)

	return max(first, second)
}

func GetMaxCost01TD(items []*Item, capacity int) int {
	n := len(items)
	dp := make([][]int, capacity+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	return getMaxCost01TDUtil(dp, items, n, capacity)
}

func getMaxCost01TDUtil(dp [][]int, items []*Item, n, capacity int) int {
	if capacity == 0 || n == 0 {
		return 0
	}

	if dp[capacity][n-1] != 0 {
		return dp[capacity][n-1]
	}

	// Their are two cases:
	// (1) ith item is included
	// (2) ith item is not included
	first := 0
	itemWeight := items[n-1].weight
	itemCost := items[n-1].cost

	if itemWeight <= capacity {
		first = itemCost + getMaxCost01TDUtil(dp, items, n-1, capacity-itemWeight)
	}

	second := getMaxCost01TDUtil(dp, items, n-1, capacity)
	dp[capacity][n-1] = max(first, second)

	return dp[capacity][n-1]
}

func GetMaxCost01BU(items []*Item, capacity int) int {
	n := len(items)
	dp := make([][]int, capacity+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Build table dp[][] in bottom up approach.
	// Weights considered against capacity.
	for w := 1; w <= capacity; w++ {
		for i := 1; i <= n; i++ {
			itemWeight := items[i-1].weight
			itemCost := items[i-1].cost

			// Their are two cases:
			// (1) ith item is included
			// (2) ith item is not included
			first := 0
			if itemWeight <= w {
				first = dp[w-itemWeight][i-1] + itemCost
			}

			second := dp[w][i-1]
			dp[w][i] = max(first, second)
		}
	}
	printSelectedItems(dp, items, n, capacity)
	// Number of weights considered and final capacity.
	return dp[capacity][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printSelectedItems(dp [][]int, items []*Item, n, capacity int) {
	totalCost := dp[capacity][n]
	output := ""

	for i := n - 1; i > 0; i-- {
		if totalCost != dp[capacity][i-1] {
			itemWeight := items[i].weight
			itemCost := items[i].cost

			output += " (wt:" + strconv.Itoa(itemWeight) + ", cost:" + strconv.Itoa(itemCost) + ")"
			capacity -= itemWeight
			totalCost -= itemCost
		}
	}

	fmt.Println("Selected items are:" + output)
}

func main() {
	items := []*Item{
		NewItem(10, 60),
		NewItem(40, 40),
		NewItem(20, 90),
		NewItem(30, 120),
	}
	capacity := 50

	maxCost := KS01UnboundedBU(items, capacity)
	fmt.Println("Maximum cost obtained =", maxCost)

	maxCost = GetMaxCost01(items, capacity)
	fmt.Println("Maximum cost obtained =", maxCost)

	maxCost = GetMaxCost01BU(items, capacity)
	fmt.Println("Maximum cost obtained =", maxCost)

	maxCost = GetMaxCost01TD(items, capacity)
	fmt.Println("Maximum cost obtained =", maxCost)
}

/*Maximum cost obtained = 300
Maximum cost obtained = 210
Selected items are: (wt:30, cost:120) (wt:20, cost:90)
Maximum cost obtained = 210
Maximum cost obtained = 210
*/
