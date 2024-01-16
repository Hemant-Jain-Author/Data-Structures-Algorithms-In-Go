package main

import "fmt"

func StockBuySellMaxProfit(arr []int) int {
	buyProfit := -arr[0] // Buy stock profit
	sellProfit := 0  // Sell stock profit
	n := len(arr)
	for i := 1; i < n; i++ {
		newBuyProfit := buyProfit
		if sellProfit-arr[i] > buyProfit {
			newBuyProfit = sellProfit - arr[i]
		}
		newSellProfit := sellProfit
		if buyProfit+arr[i] > sellProfit {
			newSellProfit = buyProfit + arr[i]
		}

		buyProfit = newBuyProfit
		sellProfit = newSellProfit
	}
	return sellProfit
}
func StockBuySellMaxProfit2(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -arr[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		if dp[i-1][1]-arr[i] > dp[i-1][0] {
			dp[i][0] = dp[i-1][1] - arr[i]
		}
		dp[i][1] = dp[i-1][1]
		if dp[i-1][0]+arr[i] > dp[i-1][1] {
			dp[i][1] = dp[i-1][0] + arr[i]
		}
	}
	return dp[n-1][1]
}

func StockBuySellMaxProfitTC(arr []int, t int) int {
	buyProfit := -arr[0]  // Buy stock profit
	sellProfit := 0 // Sell stock profit
	n := len(arr)
	for i := 1; i < n; i++ {
		newBuyProfit := buyProfit
		if sellProfit-arr[i] > buyProfit {
			newBuyProfit = sellProfit - arr[i]
		}
		newSellProfit := sellProfit
		if buyProfit+arr[i]-t > sellProfit {
			newSellProfit = buyProfit + arr[i] - t
		}
		buyProfit = newBuyProfit
		sellProfit = newSellProfit

	}
	return sellProfit
}

func StockBuySellMaxProfitTC2(arr []int, t int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -arr[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		if dp[i-1][1]-arr[i] > dp[i-1][0] {
			dp[i][0] = dp[i-1][1] - arr[i]
		}
		dp[i][1] = dp[i-1][1]
		if dp[i-1][0]+arr[i]-t > dp[i-1][1] {
			dp[i][1] = dp[i-1][0] + arr[i] - t
		}
	}
	return dp[n-1][1]
}

func main() {
	arr := []int{10, 12, 9, 23, 25, 55, 49, 70}
	fmt.Println("Total profit :", StockBuySellMaxProfit(arr))
	fmt.Println("Total profit :", StockBuySellMaxProfit2(arr))
	fmt.Println("Total profit :", StockBuySellMaxProfitTC(arr, 2))
	fmt.Println("Total profit :", StockBuySellMaxProfitTC2(arr, 2))
}

/*
Total profit : 69
Total profit : 69
Total profit : 63
Total profit : 63
*/