package main

import "strconv"

func equalPairs(grid [][]int) int {
	record := map[string]int{}
	row := len(grid)
	column := len(grid[0])
	result := 0
	for i := 0; i < row; i++ {
		current := ""
		for j := 0; j < column; j++ {
			current = current + strconv.Itoa(grid[i][j]) + "-"
		}
		record[current]++
	}
	for i := 0; i < column; i++ {
		current := ""
		for j := 0; j < row; j++ {
			current = current + strconv.Itoa(grid[j][i]) + "-"
		}
		if record[current] >= 1 {
			result = result + record[current]
		}
	}
	return result
}
func main() {

}
