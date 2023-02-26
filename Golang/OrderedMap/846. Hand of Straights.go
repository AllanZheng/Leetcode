package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W!=0{
		return false
	}
	sort.Ints(hand)
    rec := make(map[int]int,0)
    for _,i:=range hand{
    	if rec[i]==0{
    		rec[i]=1
		}else{
			rec[i]+=1
		}
	}
	var keys []int
	for k := range rec {
		keys = append(keys, k)
	}
	sort.Ints(keys)
    for i:=0;i<len(keys);i++{
    	if rec[keys[i]]<0{
    		return false
		}else if rec[keys[i]]>0{
			for j:=keys[i]+1;j<keys[i]+W;j++ {
				if rec[j] > 0 {
					rec[j] -= rec[keys[i]]
				} else {
					return false
				}
			}
			rec[keys[i]] = 0
		}

	}

    return  true
}

func main(){
	fmt.Println(isNStraightHand([]int{1,2,3,4,5},4))
}