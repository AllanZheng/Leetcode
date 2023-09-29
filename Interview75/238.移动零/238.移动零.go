package main

import "fmt"

func moveZeroes(nums []int) {
	zero := 0
	nonzero := 0
	for zero < len(nums) && nonzero < len(nums) {
		if nums[zero] == 0 {
			if nums[nonzero] != 0 {
				nums[zero], nums[nonzero] = nums[nonzero], nums[zero]

			} else {
				nonzero++
			}
		} else {
			zero++
			nonzero++
		}
	}

}

func main() {
	result := []int{1, 0}
	moveZeroes(result)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
