package main

import "fmt"

func isSubsequence(s string, t string) bool {
	sbyte := []byte(s)
	tbyte := []byte(t)
	index := 0
	for i := 0; i < len(sbyte); i++ {
		for index < len(tbyte) && sbyte[i] != tbyte[index] {
			index++
		}
		if index >= len(tbyte) {
			return false
		} else {
			index++
		}
	}
	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("aaaa", "bbaaa"))
}
