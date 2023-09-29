package main

import "fmt"

func maxVowels(s string, k int) int {
	index := 0
	bytes := []byte(s)
	length := len(bytes)
	vowels := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	current, max := 0, 0
	for i := 0; i < k; i++ {
		if vowels[bytes[i]] == 1 {
			current++
		}
	}
	max = current
	for index+k < length {
		if max == k {
			return k
		}
		if vowels[bytes[index]] == 1 {
			current--
		}
		if vowels[bytes[index+k]] == 1 {
			current++
		}
		if current > max {
			max = current
		}
		index++
	}
	return max
}
func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("xdsou", 2))
	fmt.Println(maxVowels("leetcode", 1))
	fmt.Println(maxVowels("bcd", 1))
}
