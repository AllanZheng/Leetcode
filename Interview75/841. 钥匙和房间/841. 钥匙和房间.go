package main

import (
	"container/list"
	"fmt"
)

func canVisitAllRooms(rooms [][]int) bool {
	list := list.New()
	list.Init()
	record := map[int]bool{}
	for j := 0; j < len(rooms[0]); j++ {
		list.PushBack(rooms[0][j])
	}
	record[0] = true
	for list.Len() > 0 {
		cur := list.Front().Value.(int)
		record[cur] = true
		for j := 0; j < len(rooms[cur]); j++ {
			if record[rooms[cur][j]] != true {
				list.PushBack(rooms[cur][j])
			}

		}
		list.Remove(list.Front())
	}
	if len(record) == len(rooms) {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(canVisitAllRooms([][]int{{}, {2}, {3}, {1}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 2, 3}, {2}, {3}, {}}))
}
