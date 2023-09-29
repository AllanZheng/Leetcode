package main

import "fmt"

func letterCombinations(digits string) []string {
	maps := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	digitbyte := []byte(digits)
	resbyte := make([][]byte, 0)
	for i := 0; i < len(digitbyte); i++ {

		length := len(maps[digitbyte[i]])
		curbyte := make([][]byte, 0)
		for j := 0; j < length; j++ {
			if i == 0 {
				curbyte = append(curbyte, []byte{maps[digitbyte[i]][j]})
			} else {
				for k := 0; k < len(resbyte); k++ {
					current := append(resbyte[k], maps[digitbyte[i]][j])
					x := make([]byte, len(current))
					copy(x, current)
					curbyte = append(curbyte, x)
				}
			}
		}
		resbyte = curbyte
	}

	result := []string{}
	for i := 0; i < len(resbyte); i++ {
		result = append(result, string(resbyte[i]))
	}
	return result
}
func main() {
	fmt.Println(letterCombinations("2345"))
}
