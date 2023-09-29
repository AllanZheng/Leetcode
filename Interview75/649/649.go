package main

import "container/list"

func predictPartyVictory(senate string) string {
	bytes := []byte(senate)
	queue := list.New()
	var current byte
	for i := 0; i < len(senate); i++ {
		if queue.Len() == 0 {
			queue.PushBack(bytes[i])
		} else {
			front := queue.Front().Value.(byte)
			if front == bytes[i] {
				queue.PushBack(bytes[i])
			} else {
				current = front
				queue.Remove(queue.Front())

			}
		}
	}
	if queue.Len() > 0 {
		current = queue.Front().Value.(byte)
	}
	if current == 'R' {
		return "Radiant"
	} else {
		return "Dire"
	}

}
