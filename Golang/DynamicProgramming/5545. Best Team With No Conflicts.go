package main

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {

	//rec  := make(map[int][]int,len(scores))
	rec:= make([][]int,0)
	for index,_ := range scores{

		rec = append(rec,[]int{ages[index],scores[index]})

	}

	sort.Slice(rec, func(i, j int) bool {
		return rec[i][0]<rec[j][0]|| (rec[i][0]==rec[j][0]&&rec[i][1]<rec[j][1])
	})
	//min := 0
	dp:= make([]int,len(scores))
	for i:=0;i<len(scores);i++{
		dp[i] =rec[i][1]
	}
	max :=0
	for i,_:=range scores{
		for j:=0;j<i;j++{
			if rec[j][1]<=rec[i][1]{
				dp[i]= Maximum(dp[i],dp[j]+rec[i][1])
			}

		}
		if dp[i]>max{
			max = dp[i]
		}

	}


	return max
}

func Maximum(a,b int )int{
	if a>b{
		return a
	}else{
		return b
	}

}

func main(){

	fmt.Println(bestTeamScore([]int{1,2,3,3,4,1},[]int{1,3,5,5,10,15}))

}
