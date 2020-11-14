package main

import "fmt"

func main(){

	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))

}


func maxProfit(prices []int) int {

	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	// dp[0][0] = 0
	for i:=1; i < n; i++{

		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])

	}
	return dp[n-1][0]

}



func max(x int, y int) int{
	if x > y {
		return x
	}
	return y

}