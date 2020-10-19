package main
//1380. Lucky Numbers in a Matrix
func luckyNumbers (matrix [][]int) []int {
    rowmin := []int{}
    colmax := []int{}
    for index,i := range matrix{
    	currowmin := matrix[index][0]
    	for _,j := range i{
    		if j < currowmin {
    			currowmin = j
			}
		}

		rowmin = append(rowmin, currowmin)
	}
	for i := 0; i < len(matrix[0]); i ++ {
		curcolmax := 0
		for j := 0;j < len(matrix);j++{
			if matrix[j][i] > curcolmax {
				curcolmax =  matrix[j][i]
			}
		}
		colmax = append(colmax, curcolmax)
	}
	res := []int{}
	for _,i := range rowmin{
		for _,j := range colmax{
			if i == j{
				res = append(res, i)
				print(i)
			}
			
		}
	}
	return res
}

func main() {
	input := [][]int{{7,8},{1,2}}
	print(luckyNumbers(input))
}