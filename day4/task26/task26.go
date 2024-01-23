package main

import (
	"bufio"
	"fmt"
	"os"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func getMatrix() (matrix [][]int, n, m int) {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	matrix = make([][]int, n+1)
	matrix[0] = make([]int, m+1)
	matrix[0][0], matrix[0][1] = 0, 0
	for i := 2; i <= m; i++ {
		matrix[0][i] = MaxInt
	}

	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, m+1)
		matrix[i][0] = MaxInt
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}
	return
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	matrix, n, m := getMatrix()

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			matrix[i][j] += min(matrix[i][j-1], matrix[i-1][j])
		}
	}
	fmt.Print(matrix[n][m])
}
