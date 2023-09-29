package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	max:=0
	for i:=0;i<len(sentences);i++{
		if len(sentences[i])>1{
			current := strings.Count(sentences[i]," ")

				if max < current+1 {
					max = current + 1
				}

		}else if sentences[i]!=" "{
			if max==0{
				max=1
			}
		}

	}

	return max
}

func main()  {
	fmt.Println(mostWordsFound([]string{" "}))
}