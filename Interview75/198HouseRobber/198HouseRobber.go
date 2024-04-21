package main

func rob(nums []int) int {
   if len(nums)==1{
	return nums[0]
   }
   if len(nums)==2{
	return max(nums[0],nums[1])
   }
   dp :=[]int{nums[0], max(nums[0],nums[1])}
   for i:=2;i<len(nums);i++{
	 current:=max(dp[i-2]+nums[i],dp[i-1])
	 dp = append(dp, current)
   }
   return dp[len(nums)-1]
}
func max(a int,b int)int{
	if a>b{
		return a
	}else{
		return b 
	}
}
func main() {

}
