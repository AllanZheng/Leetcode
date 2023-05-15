package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	mid :=make(map[int][][]int,0)
	res,_:= backtrack(candidates,target,mid)
	return  res
}
func backtrack(candidates []int, target int,mid map[int][][]int)([][]int,map[int][][]int){
	var res [][]int
	if len(candidates)==1{
		if target>=candidates[0]{
			div := target/candidates[0]
			sub := []int{}
			for i:=1;i<=div;i++{
				sub =append(sub,candidates[0])
				if i*candidates[0]==target{
					res = append(res,sub)
				}else{
					mid[i*candidates[0]] =append(mid[i*candidates[0]],sub)
				}
			}
		}else{
			return res,mid
		}
	}else{
		res,mid =backtrack(candidates[1:],target,mid)
		if target>candidates[0]{
			div := target/candidates[0]

			for j,arr:=range mid{
				if target-j>=0{
					if (target-j)%candidates[0]==0{
						cycle:= (target-j)/candidates[0]
						addition :=[]int{}
						for k:=0;k<cycle;k++{
							addition= append(addition,candidates[0])
						}
						for a:= range arr{
							res=append(res,append(arr[a],addition...))
						}
					}else{
						sub2:=[]int{}
						for b:=candidates[0];b<target-j;b=b+b{
							sub2 = append(sub2,candidates[0])
							for a:= range arr{
								mid[b+j] =append(mid[b+j],append(arr[a],sub2...))
							}
						}
					}
				}

			}
			sub := []int{}
			for i:=1;i<=div;i++{
				sub =append(sub,candidates[0])
				if i*candidates[0]==target{
					res = append(res,sub)
				}else{
					//if mid[i*candidates[0]]==nil{
					//	mid[i*candidates[0]]=[][]int{}
					//}

					mid[i*candidates[0]] =append(mid[i*candidates[0]],sub)
				}
			}
		}else{
			return res,mid
		}
	}
	return  res,mid
}
func main(){
	result := combinationSum([]int{7,3,2},18)
	for i:=0;i<len(result);i++{
		for j:=0;j<len(result[i]);j++{
			fmt.Print(result[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
