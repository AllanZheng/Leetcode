package main
func longestCommonSubsequence(text1 string, text2 string) int {
     bytes1 :=[]byte(text1)
	 bytes2:= []byte(text2)
	 m,n:=len(bytes1)+1,len(bytes2)+1
	 dp :=make([][]int,m)
	 for i:= range dp{
		dp[i]= make([]int,n)
		dp[i][0]=0
		dp[0][i]=0
	 }
	 
	 for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if bytes1[i]==bytes2[j]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=max(dp[i-1][j],dp[i][j-1])
			}
		}
	 }
	 return dp[m-1][n-1]
	 
}
func max(a int,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func main(){

}