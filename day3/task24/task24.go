package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	max     = int(^uint(0) >> 1)
	virtual = 3
)

type cost struct {
	a, b, c int
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	dp := make([]int, n+virtual)
	costs := make([]cost, n+virtual)

	for i := 0; i < virtual; i++ {
		dp[i] = 0
		costs[i] = cost{max, max, max}
	}

	for i := virtual; i < len(dp); i++ {
		var c cost
		fmt.Fscan(in, &c.a, &c.b, &c.c)
		costs[i] = c

		dp[i] = min(dp[i-1]+costs[i].a, dp[i-2]+costs[i-1].b, dp[i-3]+costs[i-2].c)
	}

	fmt.Println(dp[len(dp)-1])
}
