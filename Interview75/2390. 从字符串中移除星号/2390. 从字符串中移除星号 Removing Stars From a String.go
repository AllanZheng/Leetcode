package main

import (
	"container/list"
	"fmt"
)

func removeStars(s string) string {
	list := list.New()
	current, res := []byte{}, []byte{}
	list.Init()
	current = []byte(s)
	for i := 0; i < len(current); i++ {
		if current[i] != '*' {
			list.PushBack(current[i])
		} else {
			list.Remove(list.Back())
		}

	}
	for list.Len() > 0 {
		res = append(res, list.Front().Value.(byte))
		list.Remove(list.Front())
	}
	return string(res)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
	fmt.Println(removeStars("erase*****x"))
}
