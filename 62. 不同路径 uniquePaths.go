package main

import "fmt"

func uniquePaths(m int, n int) int {
   record:=make([][]int,m)
   for i:=0;i<m;i++{
   	record[i]=make([]int,n)
   	record[i][0]=1
   }
   for j:=0;j<n;j++{
   	record[0][j]=1
   }
	for i:=1;i<m;i++  {
		for j:=1;j<n;j++{
			record[i][j]=record[i-1][j]+record[i][j-1]
		}
	}
	return record[m-1][n-1]
}

func main()  {
	fmt.Println(uniquePaths(3,7))
	fmt.Println(uniquePaths(3,2))
	fmt.Println(uniquePaths(2,2))
	fmt.Println(uniquePaths(1,9))
}