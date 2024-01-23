package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	size   int16
	matrix [][]int16
	prev   []int16
	colors []byte
	sb     strings.Builder
}

func (gr *Graph) initGraph(matrix *[][]int16, n int16) {
	gr.size = n
	gr.matrix = *matrix
	gr.prev = make([]int16, n+1)
	gr.colors = make([]byte, n+1)
}

func readData() *Graph {
	in := bufio.NewReader(os.Stdin)

	var n int16
	fmt.Fscan(in, &n)

	matrix := make([][]int16, n+1)
	for i := int16(1); i <= n; i++ {
		matrix[i] = make([]int16, n+1)
		var j int16
		for j = 1; j <= n; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	var gr Graph
	gr.initGraph(&matrix, n)
	return &gr
}

func (gr *Graph) dfs(v int16) bool {
	gr.colors[v] = 1

	for i := int16(1); i <= gr.size; i++ {
		if i != gr.prev[v] && gr.matrix[v][i] != 0 {
			if gr.colors[i] == 0 {
				gr.prev[i] = v
				if gr.dfs(i) {
					return true
				}
			} else if gr.matrix[v][i] == 1 {
				gr.getRes(i, v)
				return true
			}
		}
	}
	gr.colors[v] = 2
	return false
}

func (gr *Graph) getRes(i, v int16) {
	gr.sb.WriteString("YES\n")

	var sb2 strings.Builder
	sb2.WriteString(strconv.Itoa(int(i)))
	sb2.WriteString(" ")

	count := 1
	for v != i {
		count++
		sb2.WriteString(strconv.Itoa(int(v)))
		sb2.WriteString(" ")
		v = gr.prev[v]
	}
	gr.sb.WriteString(strconv.Itoa(count))
	gr.sb.WriteString("\n")
	gr.sb.WriteString(sb2.String())
}

func main() {
	gr := readData()

	for i := int16(1); i <= gr.size; i++ {
		if gr.colors[i] == 0 && gr.dfs(i) {
			fmt.Print(gr.sb.String())
			return
		}
	}
	fmt.Print("NO")
}
