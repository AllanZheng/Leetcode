package main

import (
	"fmt"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	record := map[int]string{}
	for i := 0; i < len(words); i++ {
		current := []byte(words[i])
		order := int(current[len(current)-1]) - 48
		record[order] = string(current[:len(current)-1])
	}
	result := ""
	for i := 1; i <= len(words); i++ {
		result = result + record[i]
		if i != len(words) {
			result = result + " "
		}

	}
	return result
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))

	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}
