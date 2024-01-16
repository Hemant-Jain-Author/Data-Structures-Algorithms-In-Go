package main

import (
	"fmt"
	"sort"
)

type Job struct {
	id       rune
	deadline int
	profit   int
}

func NewJob(id rune, deadline int, profit int) *Job {
	p := new(Job)
	p.id = id
	p.deadline = deadline
	p.profit = profit
	return p
}

type JobSequencing struct {
	jobs  []*Job
	n     int
	maxDL int
}

func NewJobSequencing(ids []rune, deadlines []int,
	profits []int, n int) (jobSeq *JobSequencing) {
	jobSeq = &JobSequencing{}
	jobSeq.jobs = make([]*Job, n)
	jobSeq.n = n

	jobSeq.maxDL = deadlines[0]
	for i := 1; i < n; i++ {
		if deadlines[i] > jobSeq.maxDL {
			jobSeq.maxDL = deadlines[i]
		}
	}
	
	for i := 0; i < n; i++ {
		jobSeq.jobs[i] = NewJob(ids[i], deadlines[i], profits[i])
	}
	return
}

func (jobSeq *JobSequencing) Print() {
	sort.Slice(jobSeq.jobs,func(i, j int) bool {
		return jobSeq.jobs[i].profit > jobSeq.jobs[j].profit
	})
	result := make([]bool, jobSeq.maxDL)
	job := make([]rune, jobSeq.maxDL)
	profit := 0

	for i := 0; i < jobSeq.n; i++ {
		for j := jobSeq.jobs[i].deadline - 1; j >= 0; j-- {
			if result[j] == false {
				result[j] = true
				job[j] = jobSeq.jobs[i].id
				profit += jobSeq.jobs[i].profit
				break
			}
		}
	}

	fmt.Println("Profit is::", profit)
	fmt.Print("Jobs selected are :: ")
	for i := 0; i < jobSeq.maxDL; i++ {
		if job[i] != '0' {
			fmt.Print(string(job[i]), " ")
		}
	}
}

func main() {
	id := []rune{'a', 'b', 'c', 'd', 'e'}
	deadline := []int{3, 1, 2, 4, 4}
	profit := []int{50, 40, 27, 31, 30}
	js := NewJobSequencing(id, deadline, profit, 5)
	js.Print()
}

/*
Profit is:: 151
Jobs selected are :: b e a d 
*/