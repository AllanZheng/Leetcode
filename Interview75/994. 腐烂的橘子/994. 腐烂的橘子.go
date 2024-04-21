package main

func orangesRotting(grid [][]int) int {
	width := len(grid[0])
	height := len(grid)
	bad := [][]int{}
	count := 0
	result := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] == 1 {
				if len(neighbour(i, j, width, height, grid)) == 0 {
					return -1
				} else {
					count++
				}
			}
			if grid[i][j] == 2 {
				bad = append(bad, []int{i, j})
			}
		}
	}
	if count == 0 {
		return 0
	}
	curbad := [][]int{}
	for count > 0 && len(bad) > 0 {
		for i := 0; i < len(bad); i++ {

			rec := neighbour(bad[i][0], bad[i][1], width, height, grid)

			for j := 0; j < len(rec); j++ {
				if grid[rec[j][0]][rec[j][1]] == 1 {
					grid[rec[j][0]][rec[j][1]] = 2
					count--
					curbad = append(curbad, rec[j])
				}
			}
		}
		result++
		bad = curbad
		curbad = [][]int{}
	}
	if count > 0 {
		return -1
	}
	return result
}

func neighbour(i int, j int, width int, height int, grid [][]int) [][]int {
	result := [][]int{}
	if i-1 >= 0 && grid[i-1][j] >= 1 {
		result = append(result, []int{i - 1, j})
	}
	if i+1 < height && grid[i+1][j] >= 1 {
		result = append(result, []int{i + 1, j})
	}
	if j+1 < width && grid[i][j+1] >= 1 {
		result = append(result, []int{i, j + 1})
	}
	if j-1 >= 0 && grid[i][j-1] >= 1 {
		result = append(result, []int{i, j - 1})
	}
	return result
}
