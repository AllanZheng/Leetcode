package main
//1480. Running Sum of 1d Array
func runningSum(nums []int) []int {
   res := []int{}
   sum := 0
   for _,cur := range nums{
   	   sum = sum + cur
   	   res = append(res,sum)

   }
   return res
}

func main() {
	input := []int{3,1,2,10,1}
	res := runningSum(input)
	for _ ,cur := range res{
		println(cur)
	}

}