package main

import (
	"container/list"
	"fmt"
)

func longestOnes(nums []int, k int) int {
	if len(nums)==1{
	if nums[0]==1||k==1{
		return 1
	}else{
		return 0
	}
	}
       start:=0
       end:=0
       zeros:=list.New()
       zeros.Init()
       maximum:=0
       for end<len(nums){
		   if k == 0 {
			   if nums[end] == 0 {
				   start = end + 1
			   } else {
				   if end-start+1 > maximum {
					   maximum = end - start+1
				   }
			   }
		   }else {
				   if nums[end] == 0 {
					   if zeros.Len() >= k {
						   if end-start > maximum {
							   maximum = end - start
						   }
						   start = zeros.Front().Value.(int) + 1
						   zeros.Remove(zeros.Front())
					   }
					   zeros.PushBack(end)
				   }
				   if end-start+1 > maximum {
					   maximum = end - start+1
				   }
		   }
		end++
	   }

    return  maximum
}

func main()  {
	fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0},2))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1},3))
	fmt.Println(longestOnes([]int{1,1,1,1,1},0))
	fmt.Println(longestOnes([]int{0,0,0,0,1},5))
}