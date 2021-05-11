package main

import (
	"fmt"

)

func minimumOneBitOperations(n int) int {
   res:=0
   for n>0{
   	res ^= n
   	n = n/2
   }
   return res
}
func main(){

	fmt.Println(minimumOneBitOperations(6))

}
