package main

import "fmt"

type Action struct {
	row int
	col int
}

func findSoldierID(N int, actions []Action, k int) int {
	soldiers := make([]int, N)
	for i := 0; i < N; i++ {
		soldiers[i] = i + 1
	}

	for _, action := range actions {
		row := action.row - 1
		col := action.col - 1
		for row < col {
			soldiers[row], soldiers[col] = soldiers[col], soldiers[row]
			row++
			col--
		}
		for i := 0; i < N; i++ {
			fmt.Print(soldiers[i])
		}
		fmt.Println()
	}

	return soldiers[k-1]
}
func main() {
	N := 10
	//count := 2
	actions := []Action{Action{1, 5}, Action{6, 10}}
	fmt.Println(findSoldierID(N, actions, 1))
	fmt.Println("helloHSBC!")
}
