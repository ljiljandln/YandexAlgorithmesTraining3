package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	dp, inpData := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &inpData[i])
	}

	sort.Ints(inpData)

	dp[0], dp[1] = inpData[1]-inpData[0], inpData[1]-inpData[0]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + inpData[i] - inpData[i-1]
	}
	fmt.Println(dp[len(dp)-1])
}
