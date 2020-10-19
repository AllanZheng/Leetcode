package main

import "fmt"

func intervalIntersection(A [][]int, B [][]int) [][]int {
    signA := 0
    signB := 0
    res := [][]int{}
    for signA<len(A)&& signB<len(B){
    	curres,next := common(A[signA],B[signB])
    	if curres[0]!=-1{
    		res =append(res,curres)
		}
    	if next == "A"{
    		signA = signA+1
		}else{
			signB = signB + 1
		}
	}

    return res
}
func common(A []int, B []int) (res  []int,next string){
	left:=-1
	right:= -1
	if  A[0]<B[0]{
		left = B[0]
	}else{
		left = A[0]
	}
	if A[1]<B[1]{
		right = A[1]
		next  = "A"
	}else{
		right = B[1]
		next = "B"
	}
	if left>right {
		res = append(res, -1)
	}else{
		res = append(res,left,right)
	}

	return res,next
}

func main(){
	A := [][]int{{0,2},{5,10},{13,23},{24,25}}
	B :=  [][]int{{1,5},{8,12},{15,24},{25,26}}
	res := intervalIntersection(A,B)
    for _,cur := range(res){
    	fmt.Println(cur[0],cur[1])

	}

}