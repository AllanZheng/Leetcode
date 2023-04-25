package main

import (
	"fmt"
	"strings"
)
func finalValueAfterOperations(operations []string) int {
	 res :=0
     for i:=0;i<len(operations);i++{
     	if strings.Contains(operations[i],"+"){

			res++
		} else if strings.Contains(operations[i],"-"){
     		res--
		 }

	 }
	 return res
}

func main(){
	operations := []string{"++X","++X","X++"}
	fmt.Print(finalValueAfterOperations(operations))
}
