package main

import "fmt"

func eatenApples(apples []int, days []int) int {
	rec := [][]int{}
	for i:=0;i<len(apples);i++{
		if apples[i]>0{
			if days[i]>apples[i] {
				rec = append(rec,[]int{i,i+apples[i]})
			}else{
				rec = append(rec,[]int{i,i+days[i]})
			}
		}
	}
    front:=rec[0][0]
    end := rec[0][1]
    res :=0
    for _,i:= range rec{
        fmt.Println(i[0],i[1])
    	if i[0]>end{
    		res=end-front
    		front=i[0]
    		end=i[1]
		}else{
			if i[1]>end{
				end = i[1]
			}
		}
	}
    res += end-front
	return  res
}
func main(){
	fmt.Println(eatenApples([]int{1},[]int{2}))
	fmt.Println(eatenApples([]int{3,0,0,0,0,2},[]int{3,0,0,0,0,2}))
	fmt.Println(eatenApples([]int{1,2,3,5,2},[]int{3,2,1,4,2}))
	fmt.Println(eatenApples([]int{9,10,1,7,0,2,1,4,1,7,0,11,0,11,0,0,9,11,11,2,0,5,5},[]int{3,19,1,14,0,4,1,8,2,7,0,13,0,13,0,0,2,2,13,1,0,3,7}))
}