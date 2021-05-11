package main

import "fmt"

/*
unselfsolved
 */
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	res := 0
    for row:=0;row<len(matrix);row++{
    	for column := 1;column<len(matrix[0]);column++{
    		matrix[row][column] += matrix[row][column-1]

		}
	}
		for column := 0;column<len(matrix[0]);column++{
			for curcol:= column;curcol<len(matrix[0]);curcol++{
				counter := map[int]int{}
				counter[0] = 1
				cur := 0
				for k := 0; k<len(matrix);k++{
					sub :=0
					if column>0{
						sub = matrix[k][column-1]
					}

					cur += matrix[k][curcol]- sub

					res +=  counter[cur-target]
					fmt.Println(k,curcol,sub,cur,cur-target,counter[cur-target])
					//if counter[cur]!=0{
					//	counter[cur]+=1
					//}else{
					//	counter[cur]=1
					//}
                    counter[cur]++

				}
			}
		}

    return res
}



func main() {
	input := [][]int{{1,-1},{-1,1}}
	fmt.Print(numSubmatrixSumTarget(input,0))
}