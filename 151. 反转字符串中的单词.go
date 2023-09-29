package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s," ")
	words1:=[]string{}

	for i:=len(words)-1;i>=0;i--{
		current := strings.ReplaceAll(words[i]," ","",)

		if len(current)>0{
			words1 = append(words1,current)
		}

	}
	return strings.Join(words1," ")
}

func main()  {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world     app    sfdosdjo "))
}