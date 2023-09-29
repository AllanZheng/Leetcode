package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	sort.SearchInts()
	length := len(potions)
	for i := 0; i < len(spells); i++ {
		pre := 0
		suf := length - 1
		mid := 0

		for pre <= suf {
			mid = (pre + suf) / 2
			if mid == 0 {
				if int64(potions[mid])*int64(spells[i]) < success {
					if length == 2 && int64(potions[1])*int64(spells[i]) >= success {
						spells[i] = 1
					} else {
						spells[i] = 0
					}
				} else {
					spells[i] = length
				}
				break
			}
			if mid == length-1 {
				if int64(potions[mid])*int64(spells[i]) < success {
					spells[i] = 0
				} else {
					spells[i] = 1
				}
				break
			}
			if int64(potions[mid-1])*int64(spells[i]) < success && int64(potions[mid])*int64(spells[i]) >= success {
				spells[i] = length - mid
				break
			}
			if int64(potions[mid-1])*int64(spells[i]) >= success {
				suf = mid
			}
			if int64(potions[mid])*int64(spells[i]) < success {
				pre = mid + 1
			}
		}
	}

	return spells
}

func main() {
	result := successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7)
	for i := 0; i < 3; i++ {
		fmt.Print(result[i])
	}
	fmt.Println()
	result1 := successfulPairs([]int{3, 1, 2}, []int{1}, 16)
	for i := 0; i < 3; i++ {
		fmt.Print(result1[i])
	}
	fmt.Println()
	result2 := successfulPairs([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, 25)
	for i := 0; i < 7; i++ {
		fmt.Print(result2[i])
	}
}
