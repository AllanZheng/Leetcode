package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}

	num := 0
	result := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			nu, _ := strconv.Atoi(string(char))
			num = num*10 + nu
		} else if char == '[' {
			numStack = append(numStack, num)
			num = 0
			strStack = append(strStack, result)
			result = ""

		} else if char == ']' {
			popstr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			popnum := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			result = popstr + strings.Repeat(result, popnum)
			fmt.Println(popstr, popnum)
		} else {
			result = result + string(char)
		}
	}
	return result

}
func main() {
	fmt.Println(decodeString("2[d2[e2[a]]]2[bc]"))
}
