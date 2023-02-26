package main

func majorityElement(nums []int) int {
 //
 rec := map[int]int{}
	for _,cur := range nums {
		var curnum int
		curnum = int(cur)
		if _,ok := rec[curnum];ok {
			rec[curnum] = rec[curnum] + 1

		}else {
			rec[curnum] = 1
		}

	}
	for key,value := range rec{

		if value >len(nums)/2{
			return key
		}

	}

  return -1
}

func main(){
	input := []int{2,2,1,0,3,2,2}
	print(majorityElement(input))
}