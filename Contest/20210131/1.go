package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	 rec := map[int]int{}
   for i := lowLimit;i<=highLimit;i++{
   	x := getNum(i)
   	  if rec[x]==0{
   	  	rec[x] =1
	  }else{
	  	rec[x]=rec[x]+1
	  }
   }
   max :=0
   for i:=range rec{
   	if rec[i]>max{
   		max = rec[i]
	}
   }
   return max
}

func getNum(i int)(ret int){
	for i!=0{
		ret+=i%10
		i=i/10
	}
	return ret
}

func main(){
	fmt.Println(countBalls(1,10))
	fmt.Println(countBalls(5,15))
	fmt.Println(countBalls(1,280000))
}