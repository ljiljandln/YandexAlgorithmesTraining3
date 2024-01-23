package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, q int
	fmt.Fscan(in, &n, &m, &q)

	matrix := make([][]int, n+1)
	matrix[0] = make([]int, m+1)
	for i := 1; i <= n; i++ {
		matrix[i] = make([]int, m+1)
		for j := 1; j <= m; j++ {
			var number int
			fmt.Fscan(in, &number)
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1] + number
		}
	}

	var sb strings.Builder
	for ; q > 0; q-- {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		ans := matrix[x2][y2] - matrix[x2][y1-1] - matrix[x1-1][y2] + matrix[x1-1][y1-1]
		sb.WriteString(strconv.Itoa(ans))
		if q != 1 {
			sb.WriteString("\n")
		}
	}

	fmt.Print(sb.String())
}
