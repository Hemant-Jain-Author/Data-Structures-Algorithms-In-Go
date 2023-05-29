package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FastestWayBU(a, t [][]int, e, x []int, n int) int {
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}

	// Time taken to leave the first station.
	f[0][0] = e[0] + a[0][0]
	f[1][0] = e[1] + a[1][0]

	// Fill the tables f1[] and f2[] using a bottom-up approach.
	for i := 1; i < n; i++ {
		f[0][i] = min(f[0][i-1]+a[0][i], f[1][i-1]+t[1][i-1]+a[0][i])
		f[1][i] = min(f[1][i-1]+a[1][i], f[0][i-1]+t[0][i-1]+a[1][i])
	}

	return min(f[0][n-1]+x[0], f[1][n-1]+x[1])
}

func FastestWayBU2(a, t [][]int, e, x []int, n int) int {
	f1 := make([]int, n)
	f2 := make([]int, n)

	// Time taken to leave the first station.
	f1[0] = e[0] + a[0][0]
	f2[0] = e[1] + a[1][0]

	// Fill the tables f1[] and f2[] using a bottom-up approach.
	for i := 1; i < n; i++ {
		f1[i] = min(f1[i-1]+a[0][i], f2[i-1]+t[1][i-1]+a[0][i])
		f2[i] = min(f2[i-1]+a[1][i], f1[i-1]+t[0][i-1]+a[1][i])
	}

	return min(f1[n-1]+x[0], f2[n-1]+x[1])
}

func FastestWayTD(a, t [][]int, e, x []int, n int) int {
	f := make([][]int, 2)
	for i := range f {
		f[i] = make([]int, n)
	}

	// Time taken to leave the first station.
	f[0][0] = e[0] + a[0][0]
	f[1][0] = e[1] + a[1][0]

	fastestWayUtilTD(f, a, t, n-1)
	return min(f[0][n-1]+x[0], f[1][n-1]+x[1])
}

func fastestWayUtilTD(f, a, t [][]int, i int) {
	if i == 0 {
		return
	}
	fastestWayUtilTD(f, a, t, i-1)
	// Fill the tables f1[] and f2[] using a top-down approach.
	f[0][i] = min(f[0][i-1]+a[0][i], f[1][i-1]+t[1][i-1]+a[0][i])
	f[1][i] = min(f[1][i-1]+a[1][i], f[0][i-1]+t[0][i-1]+a[1][i])
}

func main() {
	a := [][]int{{7, 9, 3, 4, 8, 4}, {8, 5, 6, 4, 5, 7}}
	t := [][]int{{2, 3, 1, 3, 4}, {2, 1, 2, 2, 1}}
	e := []int{2, 4}
	x := []int{3, 2}
	n := 6

	fmt.Println("FastestWayBU:", FastestWayBU(a, t, e, x, n))
	fmt.Println("FastestWayBU2:", FastestWayBU2(a, t, e, x, n))
	fmt.Println("FastestWayTD:", FastestWayTD(a, t, e, x, n))
}

/*
FastestWayBU: 38
FastestWayBU2: 38
FastestWayTD: 38
*/
