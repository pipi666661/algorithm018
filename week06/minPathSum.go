package week07

import "fmt"

func main(){

	a := [][]int{
		{0, 1, 2, 3} ,   /*  第一行索引为 0 */
		{4, 5, 6, 7} ,   /*  第二行索引为 1 */
		{8, 9, 10, 11},   /* 第三行索引为 2 */
	}

	fmt.Println(minPathSum(a))

}

func minPathSum(grid [][]int) int {


	m := len(grid)
	n :=len(grid[0])
	if m==0 || n==0{
		return 0
	}


	dp := make([][]int, m)
	for i, arr := range grid {
		dp[i] = make([]int, len(arr))
	}



	dp[0][0]= grid[0][0]


	for k:=1; k<n; k++{
		dp[0][k] = grid[0][k]+dp[0][k-1]
	}



	for k:=1; k<m; k++{
		dp[k][0] = grid[k][0]+dp[k-1][0]
	}

	for i:=1; i<m; i++{
		for j:=1;j<n;j++{
			dp[i][j] = min(dp[i-1][j],dp[i][j-1])+ grid[i][j]

		}
	}
	return dp[m-1][n-1]

}

func min(x int, y int) int{
	if x < y {
		return x
	}

	return y
}
