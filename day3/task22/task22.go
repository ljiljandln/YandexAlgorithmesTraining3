package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := i - 1; j > 0 && j >= i-k; j-- {
			dp[i] += dp[j]
		}
	}

	fmt.Print(dp[n])
}
