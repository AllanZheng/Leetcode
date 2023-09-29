package main

import "fmt"

func reverseVowels(s string) string {
	bytes := []byte(s)
	pre := 0
	suf := len(bytes) - 1
	vowels := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	for pre < suf {
		if vowels[bytes[pre]] != 1 {
			pre++
		} else {
			if vowels[bytes[suf]] != 1 {
				suf--
			} else {
				mid := bytes[suf]
				bytes[suf] = bytes[pre]
				bytes[pre] = mid
				pre++
				suf--
			}
		}
	}
	return string(bytes)
}

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("abcdE"))
}
