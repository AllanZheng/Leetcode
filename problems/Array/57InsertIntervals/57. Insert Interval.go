package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	ins := false
	for i := 0; i < len(intervals); i++ {

		if intervals[i][1] < newInterval[0] {
			result = append(result, intervals[i])
			continue
		}
		if intervals[i][0] > newInterval[1] {
			if !ins {
				result = append(result, newInterval)
				ins = true
			}
			result = append(result, intervals[i])
			continue
		}
		if intervals[i][0] >= newInterval[0] && intervals[i][0] <= newInterval[1] && newInterval[1] < intervals[i][1] {
			newInterval[1] = intervals[i][1]
		}
		if intervals[i][1] >= newInterval[0] && intervals[i][1] <= newInterval[1] && newInterval[0] > intervals[i][0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][0] < newInterval[0] && intervals[i][1] > newInterval[1] {
			newInterval[0] = intervals[i][0]
			newInterval[1] = intervals[i][1]
		}
	}
	if len(result) == 0 || !ins {
		result = append(result, newInterval)
	}

	return result
}
func main() {
	res := insert([][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}, []int{4, 8})
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i][0], res[i][1])
	}
}
