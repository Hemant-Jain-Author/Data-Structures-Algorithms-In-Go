package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func FastestWayBU(a [][]int, t [][]int, e []int, x []int, n int) int {
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}

	// Time taken to leave first station.
	f[0][0] = e[0] + a[0][0]
	f[1][0] = e[1] + a[1][0]

	// Fill the tables f1[] and f2[] using bottom up approach.
	for i := 1; i < n; i++ {
		f[0][i] = min(f[0][i-1]+a[0][i], f[1][i-1]+t[1][i-1]+a[0][i])
		f[1][i] = min(f[1][i-1]+a[1][i], f[0][i-1]+t[0][i-1]+a[1][i])
	}

	// Consider exit times and return minimum.
	return min(f[0][n-1]+x[0], f[1][n-1]+x[1])
}

func FastestWayBU2(a [][]int, t [][]int, e []int, x []int, n int) int {
	f1 := make([]int, n)
	f2 := make([]int, n)

	// Time taken to leave first station.
	f1[0] = e[0] + a[0][0]
	f2[0] = e[1] + a[1][0]

	// Fill the tables f1[] and f2[] using bottom up approach.
	for i := 1; i < n; i++ {
		f1[i] = min(f1[i-1]+a[0][i], f2[i-1]+t[1][i-1]+a[0][i])
		f2[i] = min(f2[i-1]+a[1][i], f1[i-1]+t[0][i-1]+a[1][i])
	}
	return min(f1[n-1]+x[0], f2[n-1]+x[1])
}

func FastestWayTD(a [][]int, t [][]int, e []int, x []int, n int) int {
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}

	// Time taken to leave first station.
	f[0][0] = e[0] + a[0][0]
	f[1][0] = e[1] + a[1][0]

	fastestWayUtilTD(f, a, t, n-1)
	return min(f[0][n-1]+x[0], f[1][n-1]+x[1])
}

func fastestWayUtilTD(f [][]int, a [][]int, t [][]int, i int) {
	if i == 0 {
		return
	}
	fastestWayUtilTD(f, a, t, i-1)
	// Fill the tables f1[] and f2[] using top-down approach.
	f[0][i] = min(f[0][i-1]+a[0][i], f[1][i-1]+t[1][i-1]+a[0][i])
	f[1][i] = min(f[1][i-1]+a[1][i], f[0][i-1]+t[0][i-1]+a[1][i])
}

func main() {
	a := [][]int{{7, 9, 3, 4, 8, 4}, {8, 5, 6, 4, 5, 7}}
	t := [][]int{{2, 3, 1, 3, 4}, {2, 1, 2, 2, 1}}
	e := []int{2, 4}
	x := []int{3, 2}
	n := 6
	fmt.Println(FastestWayBU2(a, t, e, x, n))
	fmt.Println(FastestWayBU(a, t, e, x, n))
	fmt.Println(FastestWayTD(a, t, e, x, n))
}

/*
59
59
59
*/
