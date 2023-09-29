package main

import (
	"container/list"
	"fmt"
)

func asteroidCollision(asteroids []int) []int {
	stack := list.New()
	for i := 0; i < len(asteroids); i++ {
		if stack.Len() == 0 {
			stack.PushFront(asteroids[i])
		} else {
			frontv := stack.Front().Value.(int)
			collision := false
			for frontv > 0 && asteroids[i] < 0 {
				absv := -asteroids[i]
				if frontv < absv {
					stack.Remove(stack.Front())
					if stack.Len() > 0 {
						frontv = stack.Front().Value.(int)
					} else {
						break
					}

				} else {
					if frontv == absv {
						stack.Remove(stack.Front())

					}
					collision = true
					break

				}

			}

			if !collision {
				stack.PushFront(asteroids[i])
			}

		}
	}
	Result := []int{}

	for stack.Len() > 0 {

		Result = append(Result, stack.Back().Value.(int))
		stack.Remove(stack.Back())
	}
	return Result

}
func main() {
	res := asteroidCollision([]int{10, 2, -8, -9})
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])

	}
	fmt.Println()
	res1 := asteroidCollision([]int{5, 10, -5})
	for i := 0; i < len(res1); i++ {
		fmt.Print(res1[i])

	}
	fmt.Println()
	res2 := asteroidCollision([]int{-2, -1, 1, 2})
	for i := 0; i < len(res2); i++ {
		fmt.Print(res2[i])

	}
	fmt.Println()
}
