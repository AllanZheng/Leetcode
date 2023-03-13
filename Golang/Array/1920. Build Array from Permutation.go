package main

func buildArray(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[nums[i]])
	}
	return res
}

func main() {
	input := []int{0, 2, 1, 5, 3, 4}
	output := buildArray(input)
	for i := 0; i < len(output); i++ {
		println(output[i])
	}

}
