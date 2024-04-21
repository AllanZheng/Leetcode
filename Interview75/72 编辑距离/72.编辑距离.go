package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}
	record := make([][]int, len1)
	for i := 0; i < len1; i++ {
		record[i] = make([]int, len2)
		if i == 0 {
			if word1[i] != word2[0] {
				record[i][0] = 1
			} else {
				record[i][0] = 0
			}
		} else {
			if word1[i] == word2[0] && record[i-1][0] == i {
				record[i][0] = i
			} else {
				record[i][0] = record[i-1][0] + 1
			}
		}

	}
	for j := 1; j < len2; j++ {
		if word1[0] == word2[j] && record[0][j-1] == j {
			record[0][j] = j
		} else {
			record[0][j] = record[0][j-1] + 1
		}
	}
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			curmin := record[i-1][j-1]
			if record[i-1][j] < curmin {
				curmin = record[i-1][j]
			} else if record[i][j-1] < curmin {
				curmin = record[i][j-1]
			}
			if word1[i] == word2[j] {
				record[i][j] = record[i-1][j-1]
			} else {
				record[i][j] = curmin + 1
			}
			fmt.Print(record[i][j], " ")
		}
		fmt.Println()
	}
	return record[len1-1][len2-1]
}
func main() {
	fmt.Println(minDistance("zoologicoarchaeologist", "zoogeologist"))

}
