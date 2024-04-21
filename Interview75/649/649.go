package main

import (
	"container/list"
	"strings"
)

func predictPartyVictory(senate string) string {
	bytes := []byte(senate)
	rq, dq := list.New(), list.New()

	length := len(senate)
	for i := 0; i < length; i++ {
		if bytes[i] == 'R' {
			rq.PushBack(i)
		} else {
			dq.PushBack(i)
		}
	}
	for rq.Len() > 0 && dq.Len() > 0 {
		currentR := rq.Front()
		currentD := dq.Front()
		if currentD.Value.(int) < currentR.Value.(int) {

			dq.PushBack(currentD.Value.(int) + length)
		} else {
			rq.PushBack(currentR.Value.(int) + length)
		}
		rq.Remove(rq.Front())
		dq.Remove(dq.Front())
	}
	if rq.Len() > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
	(s string, substr string)
}
