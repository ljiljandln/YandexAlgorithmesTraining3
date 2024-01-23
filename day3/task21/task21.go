package main

import "fmt"

func main() {
	var n, size int
	fmt.Scan(&n)

	if n > 3 {
		size = n
	} else {
		size = 3
	}

	dp := make([]int, size+1)
	dp[1] = 2
	dp[2] = 4
	dp[3] = 7

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	fmt.Print(dp[n])
}
