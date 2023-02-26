package main
func kthSmallestPath(destination []int, k int) string {
	res :=""
	cur := []int{0,0}
		for k -destination[0]>0{
			res = res+"V"
			k = k-destination[0]
			cur[1]+=1
		}

}
