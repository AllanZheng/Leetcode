package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	byte1 := []byte(word1)
	byte2 := []byte(word2)
	index := 0
	res := []byte{}
	for {
		if index < len(byte1) {
			res = append(res, byte1[index])
		} else {
			res = append(res, byte2[index:]...)
			break
		}
		if index < len(byte2) {
			res = append(res, byte2[index])
		} else {
			res = append(res, byte1[index+1:]...)
			break
		}
		index++
	}
	result := string(res)
	return result
}

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcde", "pq"))
	fmt.Println(mergeAlternately("abc", "pqr"))
}
