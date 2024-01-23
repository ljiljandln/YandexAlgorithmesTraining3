package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMatrix() (matrix [][]int, prev [][]string, n, m int) {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	matrix = make([][]int, n+1)
	matrix[0] = make([]int, m+1)
	prev = make([][]string, n+1)
	prev[0] = make([]string, m+1)

	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, m+1)
		prev[i] = make([]string, m+1)
		for j := 1; j <= m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}
	return
}

func main() {
	matrix, prev, n, m := getMatrix()

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			r, d := matrix[i][j-1], matrix[i-1][j]
			if r > d {
				prev[i][j] = "R"
				matrix[i][j] += r
			} else {
				prev[i][j] = "D"
				matrix[i][j] += d
			}
		}
	}

	fmt.Println(matrix[n][m])

	ans := make([]string, m+n-2)
	i, j := n, m
	for k := m + n - 3; k >= 0; k-- {
		ans[k] = prev[i][j]
		if prev[i][j] == "R" {
			j--
		} else {
			i--
		}
	}

	sb := strings.Builder{}
	for _, str := range ans {
		sb.WriteString(str)
		sb.WriteString(" ")
	}
	fmt.Print(sb.String())
}
