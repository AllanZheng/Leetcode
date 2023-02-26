package main

import "fmt"

func countWine( m int ,  n int ) int {
	ping := 0
	gai := 0
	res := 0

	for ping >=2 ||gai>=4||n>=m{
		gai += n/m
		ping +=n/m
		res+=n/m
		n = n%m + gai/4*m+ping/2*m
		gai =gai%4
		ping = ping%2

	}
	return res

	// write code here
}
func main(){
	fmt.Println(countWine(2,3))
}