package main

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	New1 := RecentCounter{[]int{}}

	return New1
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)
	result := 0
	for i := 0; i < len(this.requests); i++ {
		if t-this.requests[i] <= 3000 {
			result++
		}

	}
	return result
}
