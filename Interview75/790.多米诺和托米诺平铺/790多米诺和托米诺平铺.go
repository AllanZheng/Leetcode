package main

import (
	"fmt"
	"math"
)

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 5
	}
	//var result int
	modnumber := int(math.Pow(10, 9) + 7)
	dp := []int{1, 2, 5}
	for i := 3; i < n; i++ {
		dp = append(dp, (2*dp[i-1]+dp[i-3])%modnumber)

	}

	// fmt.Println(dp[n-1])
	// if dp[n-1] > modnumber {
	// 	result = int(dp[n-1] % modnumber)
	// } else {
	// 	result = int(dp[n-1])
	// }
	return dp[n-1]
}

func main() {
	fmt.Println(numTilings(60))
	fmt.Println(numTilings(30))
}
