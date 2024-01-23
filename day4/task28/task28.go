package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	dp := make([][]int, n+1)
	dp[0], dp[1] = make([]int, m+1), make([]int, m+1)
	dp[1][1] = 1

	left, right := 2, 3
	for i := 2; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := left; j <= min(m, right); j++ {
			dp1, dp2 := dp[i-2][j-1], dp[i-1][j-2]
			if dp1 != 0 || dp2 != 0 {
				dp[i][j] = dp1 + dp2
			}
		}
		right += 2
		if i%2 == 1 {
			left++
		}
	}

	fmt.Print(dp[n][m])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
