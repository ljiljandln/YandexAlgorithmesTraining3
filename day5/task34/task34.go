package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	neighbors []int
}

type Graph struct {
	vertexes []Vertex
	colors   []byte
	order    []int
}

func (gr *Graph) initGraph(v int) {
	gr.vertexes = make([]Vertex, v+1)
	gr.colors = make([]byte, v+1)
}

func (gr *Graph) addNeighbour(a, b int) {
	gr.vertexes[a].neighbors = append(gr.vertexes[a].neighbors, b)
}

func (gr *Graph) dfs(v int) bool {
	gr.colors[v] = 1

	for _, n := range gr.vertexes[v].neighbors {
		if gr.colors[n] == 1 {
			return false
		} else if gr.colors[n] == 0 {
			ans := gr.dfs(n)
			if !ans {
				return false
			}
		}
	}
	gr.colors[v] = 2
	gr.order = append(gr.order, v)
	return true
}

func (gr *Graph) printRes(v int) {
	var sb strings.Builder
	for i := v; i > 0; i-- {
		sb.WriteString(strconv.Itoa(gr.order[i-1]))
		if i != 1 {
			sb.WriteString(" ")
		}
	}
	fmt.Print(sb.String())
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var v, e int
	fmt.Fscan(in, &v, &e)

	var gr Graph
	gr.initGraph(v)

	var a, b int
	for ; e > 0; e-- {
		fmt.Fscan(in, &a, &b)
		gr.addNeighbour(a, b)
	}

	for i := 1; i <= v; i++ {
		if gr.colors[i] == 0 {
			ans := gr.dfs(i)
			if !ans {
				fmt.Print(-1)
				return
			}
		}
	}

	gr.printRes(v)
}
