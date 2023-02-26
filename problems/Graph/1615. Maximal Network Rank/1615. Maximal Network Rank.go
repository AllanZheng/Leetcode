package main

import "fmt"

func maximalNetworkRank(n int, roads [][]int) int {
	rec := make(map[int][]int,n)
	for _,i:= range roads{
		rec[i[0]]= append(rec[i[0]],i[1])
		rec[i[1]]=append(rec[i[1]],i[0])
	}
	max :=0

	for a:=range rec{
		for b := range rec {
			if a!=b {
				curmax := len(rec[b])
				for _, i := range rec[b] {
					if i == a {
						curmax -= 1
						break
					}
				}
				if len(rec[a])+curmax>max{
					max =len(rec[a])+curmax
				}
			}
		}
	}

	return max
}


func main(){

	fmt.Println(maximalNetworkRank(8,[][]int{{0,1},{2,3},{1,2},{2,4},{5,6},{5,7}}))

}