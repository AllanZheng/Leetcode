package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	if n == 0 {
		return true
	}
	if length == 1 {
		if flowerbed[0] == 0 && n == 1 {
			return true
		} else {
			return false
		}
	}
	if length == 2 {
		if n == 1 && (flowerbed[0] == 0 && flowerbed[1] == 0) {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < length; {
		if n == 0 {
			break
		}
		if flowerbed[i] == 1 {
			i = i + 2
		} else {
			if i == 0 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--

				}
				i = i + 1
				continue
			}
			if i+1 == length {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n--
					break
				}
			}
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
				i = i + 1
			} else {
				i = i + 1
			}
		}

	}
	if n == 0 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0, 1, 0, 0, 1}, 3))
}
