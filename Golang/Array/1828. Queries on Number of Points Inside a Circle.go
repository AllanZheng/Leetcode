package main

func countPoints(points [][]int, queries [][]int) []int {
	var res []int
	for i := 0; i < len(queries); i++ {
		curRes := 0
		for j := 0; j < len(points); j++ {
			row := (queries[i][0] - points[j][0]) * (queries[i][0] - points[j][0])
			column := (queries[i][1] - points[j][1]) * (queries[i][1] - points[j][1])
			if row+column <= queries[i][2]*queries[i][2] {
				curRes = curRes + 1
			}
		}
		res = append(res, curRes)
	}
	return res
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	queries := [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}
	output := countPoints(points, queries)
	for i := 0; i < len(output); i++ {
		println(output[i])
	}
}
