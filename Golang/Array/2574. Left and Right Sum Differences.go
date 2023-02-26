package main

func leftRigthDifference(nums []int) []int {
	var leftsum, rightsum, res []int
	leftsum = append(leftsum, 0)
	rightsum = append(rightsum, 0)
	for i := 0; i < len(nums)-1; i++ {

		leftsum = append(leftsum, leftsum[len(leftsum)-1]+nums[i])
		rightsum = append([]int{rightsum[0] + nums[len(nums)-i-1]}, rightsum...)
	}
	for i := 0; i < len(nums); i++ {
		cur := leftsum[i] - rightsum[i]
		if cur < 0 {
			cur = -cur
		}
		res = append(res, cur)
	}
	return res
}

func main() {
	input := []int{1}
	output := leftRigthDifference(input)
	for i := 0; i < len(output); i++ {
		println(output[i])
	}

}
