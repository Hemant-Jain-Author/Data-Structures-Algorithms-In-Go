package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x int
	y int
}

func NewPoint(a int, b int) *Point {
    p := new(Point)
    p.x = a
    p.y = b
    return p
}

func distance(a *Point, b *Point) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

func ClosestPairBF(arr [][]int) float64 {
	n := len(arr)
	var dMin float64
	dMin = 999999
	var d float64
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d = math.Sqrt(float64((arr[i][0]-arr[j][0])*(arr[i][0]-arr[j][0]) +
				(arr[i][1]-arr[j][1])*(arr[i][1]-arr[j][1])))
			if d < dMin {
				dMin = d
			}
		}
	}
	return dMin
}

func ClosestPairDC(arr [][]int) float64 {
	n := len(arr)
	p := make([]*Point, n)
	for i := 0; i < n; i++ {
		p[i] = NewPoint(arr[i][0], arr[i][1])
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].x < p[j].x
	})

	q := make([]*Point, n)
	copy(q, p)

	sort.Slice(p, func(i, j int) bool {
		return p[i].y < p[j].y
	})

	return closestPairUtil(p, 0, n-1, q, n)
}


func closestPairUtil(p []*Point, start int, stop int, q []*Point, n int) float64 {
	if stop-start < 1 {
		return 999999
	}
	if stop-start == 1 {
		return distance(p[start], p[stop])
	}
	mid := (start + stop) / 2
	dl := closestPairUtil(p, start, mid, q, n)
	dr := closestPairUtil(p, mid+1, stop, q, n)
	d := math.Min(dl, dr)
	strip := make([]*Point, n)
	j := 0
	for i := 0; i < n; i++ {
		if math.Abs(float64(q[i].x-p[mid].x)) < d {
			strip[j] = q[i]
			j++
		}
	}
	return math.Min(d, stripMin(strip, j, d))
}

func stripMin(q []*Point, n int, d float64) float64 {
	min := d
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && (float64)(q[j].y-q[i].y) < min; j++ {
			d = distance(q[i], q[j])
			if d < min {
				min = d
			}
		}
	}
	return min
}

func main() {
	arr := [][]int{
		{648, 896},
		{269, 879},
		{250, 922},
		{453, 347},
		{213, 17}}
	fmt.Println("Smallest distance is:", ClosestPairBF(arr))
	fmt.Println("Smallest distance is:", ClosestPairDC(arr))
}

/*
Smallest distance is: 47.01063709417264
Smallest distance is: 47.01063709417264
*/