package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
   rec := make(map[int][]int,0)
   for _,i:= range pieces{
   	if len(i)==1{
   		rec[i[0]] = append(rec[i[0]],0)
	}else{
		for j:=1;j<len(i);j++{
			rec[i[0]] = append(rec[i[0]],i[j])
		}
	}
   }
   for i:=0;i<len(arr);i++{
   	if _,ok := rec[arr[i]];!ok{
		return false
	}
   	if rec[arr[i]][0]!=0{
   		order := rec[arr[i]]
   		for _,j:= range order{
   			i++
   			if j != arr[i]{
   				return false
			}
		}

	}
   }
   return  true
}

func main(){
	a := []int{91,4,64}
	b := [][]int{{64,4,91}}
	fmt.Println(canFormArray(a,b))
}