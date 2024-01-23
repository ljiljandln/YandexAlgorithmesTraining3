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
	vCount     int
	vertexes   []Vertex
	visited    []bool
	components [][]int
}

func (graph *Graph) initGraph(v int) {
	graph.vertexes = make([]Vertex, v+1)
	graph.visited = make([]bool, v+1)
	graph.vCount = v
}

func (graph *Graph) addNeighbour(a, b int) {
	graph.vertexes[a].neighbors = append(graph.vertexes[a].neighbors, b)
	graph.vertexes[b].neighbors = append(graph.vertexes[b].neighbors, a)
}

func (graph *Graph) dfs(v, comp int) {
	graph.visited[v] = true
	graph.components[comp] = append(graph.components[comp], v)

	for _, to := range graph.vertexes[v].neighbors {
		if !graph.visited[to] {
			graph.dfs(to, comp)
		}
	}
}

func (graph *Graph) printRes() {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(len(graph.components)))
	sb.WriteString("\n")
	for i := 0; i < len(graph.components); i++ {
		sb.WriteString(strconv.Itoa(len(graph.components[i])))
		sb.WriteString("\n")
		for j := 0; j < len(graph.components[i]); j++ {
			sb.WriteString(strconv.Itoa(graph.components[i][j]))
			if j != len(graph.components[i])-1 {
				sb.WriteString(" ")
			}
		}
		if i != len(graph.components) {
			sb.WriteString("\n")
		}
	}
	fmt.Print(sb.String())
}

func (graph *Graph) searchComp() {
	comp := 0
	for i := 1; i <= graph.vCount; i++ {
		if !graph.visited[i] {
			var newComp []int
			graph.components = append(graph.components, newComp)
			graph.dfs(i, comp)
			comp++
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var v, e int
	fmt.Fscan(in, &v, &e)

	var graph Graph
	graph.initGraph(v)
	var a, b int

	for i := 0; i < e; i++ {
		fmt.Fscan(in, &a, &b)
		graph.addNeighbour(a, b)
	}

	graph.searchComp()
	graph.printRes()
}
