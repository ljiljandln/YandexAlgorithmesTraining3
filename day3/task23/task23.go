package main

import (
	"fmt"
)

func setPrev(prev *[]int, i, div int) {
	_prev := *prev
	step := 0
	for ; step < i%div; step++ {
		_prev[i-step] = i - step - 1
	}
	_prev[i-step] = i / div
}

func minOfThree(a, b, c, i int, prev *[]int) int {
	if a <= b && a <= c {
		_prev := *prev
		_prev[i] = i - 1
		return a
	} else if b <= a && b <= c {
		setPrev(prev, i, 2)
		return b
	}
	setPrev(prev, i, 3)
	return c
}

func main() {
	var n, size int
	fmt.Scan(&n)

	if n > 3 {
		size = n
	} else {
		size = 3
	}

	dp := make([]int, size+1)
	prev := make([]int, size+1)
	dp[2], prev[2], dp[3], prev[3] = 1, 1, 1, 1

	for i := 4; i <= n; i++ {
		a := dp[i-1] + 1
		b := dp[i/2] + i%2 + 1
		c := dp[i/3] + i%3 + 1

		dp[i] = minOfThree(a, b, c, i, &prev)
	}

	fmt.Println(dp[n])

	ans := make([]int, dp[n]+1)
	step, curr := dp[n], n
	for ; curr > 0; step-- {
		ans[step] = curr
		curr = prev[curr]
	}

	for _, val := range ans {
		fmt.Printf("%d ", val)
	}
}
