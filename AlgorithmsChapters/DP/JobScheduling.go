package main

import (
	"fmt"
	"sort"
)

type Job struct {
	start int
	stop  int
	value int
}

func NewJob(i1, i2, i3 int) *Job {
	j := new(Job)
	j.start = i1
	j.stop = i2
	j.value = i3
	return j
}

// Also known as Activity Selection Weighted.
func MaxValueJobs(s []int, f []int, v []int, n int) int {
	act := make([]*Job, n)
	for i := 0; i < n; i++ {
		act[i] = NewJob(s[i], f[i], v[i])
	}

	// sort according to finish time.
	sort.Slice(act, func(i, j int) bool {
		return act[i].stop < act[j].stop
	})

	return maxValueJobUtil(act, n)
}

func maxValueJobUtil(arr []*Job, n int) int {
	// Base case
	if n == 1 {
		return arr[0].value
	}

	// Find Value when current job is included
	incl := arr[n-1].value
	for j := n - 1; j >= 0; j-- {
		if arr[j].stop <= arr[n-1].start {
			incl += maxValueJobUtil(arr, j+1)
			break
		}
	}

	// Find Value when current job is excluded
	excl := maxValueJobUtil(arr, n-1)
	return max(incl, excl)
}

func MaxValueJobsTD(s []int, f []int, v []int, n int) int {
	act := make([]*Job, n)
	for i := 0; i < n; i++ {
		act[i] = NewJob(s[i], f[i], v[i])
	}

	// sort according to finish time.
	sort.Slice(act, func(i, j int) bool {
		return act[i].stop < act[j].stop
	}) 
	dp := make([]int, n)
	return maxValueJobUtilTD(dp, act, n)
}

func maxValueJobUtilTD(dp []int, arr []*Job, n int) int {
	if n == 0 {
		return 0
	}
	if dp[n-1] != 0 {
		return dp[n-1]
	}

	// Find Value when current job is included
	incl := arr[n-1].value
	for j := n - 2; j >= 0; j-- {
		if arr[j].stop <= arr[n-1].start {
			incl += maxValueJobUtilTD(dp, arr, j+1)
			break
		}
	}

	// Find Value when current job is excluded
	excl := maxValueJobUtilTD(dp, arr, n-1)
	dp[n-1] = max(incl, excl)
	return dp[n-1]
}

func MaxValueJobsBU(s []int, f []int, v []int, n int) int {
	act := make([]*Job, n)
	for i := 0; i < n; i++ {
		act[i] = NewJob(s[i], f[i], v[i])
	}

	// sort according to finish time.
	sort.Slice(act, func(i, j int) bool {
		return act[i].stop < act[j].stop
	})

	dp := make([]int, n)
	dp[0] = act[0].value
	for i := 1; i < n; i++ {
		incl := act[i].value
		for j := i - 1; j >= 0; j-- {
			if act[j].stop <= act[i].start {
				incl += dp[j]
				break
			}
		}
		dp[i] = max(incl, dp[i-1])
	}
	return dp[n-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	start := []int{1, 5, 0, 3, 5, 6, 8}
	finish := []int{2, 6, 5, 4, 9, 7, 9}
	value := []int{2, 2, 4, 3, 10, 2, 8}
	n := len(start)
	fmt.Println(MaxValueJobs(start, finish, value, n))
	fmt.Println(MaxValueJobsTD(start, finish, value, n))
	fmt.Println(MaxValueJobsBU(start, finish, value, n))
}

/*
17
17
17
*/