package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vertex struct {
	neighbors []int
}

type Graph struct {
	vertexes []Vertex
	colors   []byte
}

func (gr *Graph) initGraph(v int) {
	gr.vertexes = make([]Vertex, v+1)
	gr.colors = make([]byte, v+1)
}

func (gr *Graph) addNeighbour(a, b int) {
	gr.vertexes[a].neighbors = append(gr.vertexes[a].neighbors, b)
	gr.vertexes[b].neighbors = append(gr.vertexes[b].neighbors, a)
}

func (gr *Graph) dfs(v int, color byte) bool {
	gr.colors[v] = color

	for _, n := range gr.vertexes[v].neighbors {
		if gr.colors[n] == color ||
			(gr.colors[n] == 0 && !gr.dfs(n, 3-color)) {
			return false
		}
	}
	return true
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
		if gr.colors[i] == 0 && !gr.dfs(i, 1) {
			fmt.Print("NO")
			return
		} else {
			for _, n := range gr.vertexes[i].neighbors {
				if gr.colors[n] == gr.colors[i] {
					fmt.Print("NO")
					return
				}
			}
		}
	}

	fmt.Print("YES")
}
