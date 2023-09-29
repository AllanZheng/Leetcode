package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {

	digits := []int{}
	for num != 0 {
		cur := num % 10
		digits = append(digits, cur)

		num = num / 10
	}
	sort.Ints(digits)

	return (digits[0]+digits[1])*10 + digits[2] + digits[3]

}

func main() {
	fmt.Println(minimumSum(4091))
}
