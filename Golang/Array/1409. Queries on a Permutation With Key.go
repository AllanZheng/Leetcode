package main
func processQueries(queries []int, m int) []int {
	per   := map[int]int{}
	res := []int{}
	for  cur:= 1; cur <=m; cur = cur + 1 {
		per[cur] = cur-1
	}

	for _,cur := range queries{
		res = append(res,per[cur])

        for i:=1;i<=m;i++{
        	if per[i]<per[cur]{
        		per[i] = per[i]+1
			}
		}
		per[cur] = 0
	}
	return res
}

/*
我的是暴力模拟仿真解法，每次都更新数字对应的位置。ONlogn解法是树状数组
 */
func main() {
	input := []int{7,5,5,8,3}
	m := 8
	res := processQueries(input,m)
	for _ ,cur := range res{
		println(cur)
	}

}

