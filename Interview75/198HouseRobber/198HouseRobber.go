package main

func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	odd := 1
	even := 0
	oddsum, evensum := nums[odd], nums[even]
	for i := 0; i <= length/2; i++ {
		if odd+2 < length {
			odd = odd + 2
			oddsum += nums[odd]
		}
		if even+2 < length {
			even = even + 2
			evensum += nums[even]
		}
	}
	if oddsum > evensum {
		return oddsum
	} else {
		return evensum
	}
}
func main() {

}
