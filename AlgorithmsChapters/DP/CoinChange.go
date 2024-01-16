package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// Greedy
func MinCoins(coins []int, n int, val int) int {
	if val <= 0 {
		return 0
	}
	count := 0
	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })
	for i := n - 1; i >= 0 && val > 0; {
		if coins[i] <= val {
			count++
			val -= coins[i]
		} else {
			i--
		}
	}
	if val == 0 {
		return count
	}
	return -1
}

// Brute Force
func MinCoins2(coins []int, n int, val int) int {
	if val == 0 {
		return 0
	}

	count := math.MaxInt32
	for i := 0; i < n; i++ {
		if coins[i] <= val {
			subCount := MinCoins2(coins, n, val-coins[i])
			if subCount >= 0 {
				count = min(count, subCount+1)
			}
		}
	}
	if count != math.MaxInt32 {
		return count
	}
	return -1
}

func MinCoinsBU(coins []int, n int, val int) int {
	count := make([]int, val+1)
	for i := range count {
		count[i] = math.MaxInt32
	}
	count[0] = 0
	for i := 1; i <= val; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				if count[i-coins[j]] != math.MaxInt32 {
					count[i] = min(count[i], count[i-coins[j]]+1)
				}
			}
		}
	}
	if count[val] != math.MaxInt32 {
		return count[val]
	}
	return -1
}

func MinCoinsBU2(coins []int, n int, val int) int {
	count := make([]int, val+1)
	cvalue := make([]int, val+1)
	for i := range count {
		count[i] = math.MaxInt32
	}
	count[0] = 0
	for i := 1; i <= val; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				if count[i-coins[j]] != math.MaxInt32 {
					count[i] = min(count[i], count[i-coins[j]]+1)
					cvalue[i] = coins[j];
				}
			}
		}
	}
	if count[val] == math.MaxInt32 {
		return -1
	}
	
	printCoins(cvalue, val)
	return count[val]
}

func printCoinsUtil(cvalue []int, val int) string {
	if (val > 0) {
		ret := printCoinsUtil(cvalue, val - cvalue[val]);
		return ret + strconv.Itoa(cvalue[val]) + " ";
	}
	return "";
}

func printCoins(cvalue []int, val int) {
	ret := printCoinsUtil(cvalue, val);
	fmt.Println("Coins are :", ret);
}

func MinCoinsTD(coins []int, n int, val int) int {
	count := make([]int, val+1)
	for i := range count {
		count[i] = math.MaxInt32
	}
	return MinCoinsTDUtil(count, coins, n, val)
}

func MinCoinsTDUtil(count []int, coins []int, n int, val int) int {
	// Base Case
	if val == 0 {
		return 0
	}
	if count[val] != math.MaxInt32 {
		return count[val]
	}

	// Recursion
	for i := 0; i < n; i++ {
		// check validity of a sub-problem
		if coins[i] <= val {
			subCount := MinCoinsTDUtil(count, coins, n, val-coins[i])
			if subCount != math.MaxInt32 {
				count[val] = min(count[val], subCount+1)
			}
		}
	}
	return count[val]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{5, 6}
	value := 16
	n := len(coins)
	fmt.Println("Count is :", MinCoins(coins, n, value))
	fmt.Println("Count is :", MinCoins2(coins, n, value))
	fmt.Println("Count is :", MinCoinsBU(coins, n, value))
	fmt.Println("Count is :", MinCoinsBU2(coins, n, value))
	fmt.Println("Count is :", MinCoinsTD(coins, n, value))
}

/*
Count is : -1
Count is : 3
Count is : 3
Coins are : 5 5 6 
Count is : 3
Count is : 3
*/
