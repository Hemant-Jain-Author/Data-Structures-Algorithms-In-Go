package main

import (
	"fmt"
	"sort"
)

type Activity struct {
	start, stop int
}

func NewActivity(a int, b int) *Activity {
	p := new(Activity)
	p.start = a
	p.stop = b
	return p
}

func MaxActivities(s []int, f []int, n int) {
	act := make([]*Activity, n)
	for i := 0; i < n; i++ {
		act[i] = NewActivity(s[i], f[i])
	}

	// sort according to finish time.
	sort.Slice(act, func(i, j int) bool {
		return act[i].stop < act[j].stop
	})
	i := 0 // The first activity at index 0 is always gets selected.
	fmt.Print("Activities are : (", act[i].start, ",", act[i].stop, ")")
	for j := 1; j < n; j++ {
		// Find next activity whose start time is greater than or equal
		// to the finish time of previous activity.
		if act[j].start >= act[i].stop {
			fmt.Print(", (", act[j].start, ",", act[j].stop, ")")
			i = j
		}
	}
}

func main() {
	s := []int{1, 5, 0, 3, 5, 6, 8}
	f := []int{2, 6, 5, 4, 9, 7, 9}
	n := len(s)
	MaxActivities(s, f, n)
}