package main
type SubrectangleQueries struct {
   Rectangle [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
	newrectangle:= SubrectangleQueries{}
	newrectangle.Rectangle = rectangle
	return newrectangle
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	for currow:= row1;currow<=row2;currow++{
		for curcol:=col1;curcol<=col2;curcol++{
			this.Rectangle[currow][curcol] = newValue
		}
	}
    return
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.Rectangle[row][col]

}

func main(){

}
/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */