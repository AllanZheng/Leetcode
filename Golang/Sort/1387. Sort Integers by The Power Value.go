package main

import (
	"fmt"
	"sort"
)

func getKth(lo int, hi int, k int) int {
     if lo == hi {
     	return lo
	 }

     var midarr [][]int
     for cur:= lo;cur<=hi; cur++{
         mid := cur
         rec  := 0
         for mid!=1{
         	if mid%2==0{
         		mid = mid / 2
			}else{
				mid = mid*3  + 1
			}
         	rec = rec + 1
		 }
         midarr = append(midarr,[]int{rec,cur})
		 }



sort.Slice(midarr, func(i, j int) bool{
	return midarr[i][0] < midarr[j][0] || midarr[i][0] == midarr[j][0] && midarr[i][1] < midarr[j][1]
})
     return midarr[k-1][1]
}


func main(){

		fmt.Println(getKth(1,1000,777))

	}


