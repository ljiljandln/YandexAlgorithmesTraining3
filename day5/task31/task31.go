package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vertex struct {
	edges map[int]struct{}
}

type Graph struct {
	vertexes   []Vertex
	visited    []bool
	components []int
}

func (graph *Graph) initGraph(v int) {
	graph.vertexes = make([]Vertex, v+1)
	graph.visited = make([]bool, v+1)
}

func (graph *Graph) checkV(a int) {
	if graph.vertexes[a].edges == nil {
		graph.vertexes[a].edges = make(map[int]struct{})
	}
}

func (graph *Graph) addEdge(a, b int) {
	if _, exist := graph.vertexes[a].edges[b]; !exist && a != b {
		graph.checkV(a)
		graph.checkV(b)
		graph.vertexes[a].edges[b] = struct{}{}
		graph.vertexes[b].edges[a] = struct{}{}
	}
}

func (graph *Graph) dfs(v int) {
	graph.visited[v] = true
	graph.components = append(graph.components, v)

	for to := range graph.vertexes[v].edges {
		if !graph.visited[to] {
			graph.dfs(to)
		}
	}
}

func (graph *Graph) printRes() {
	fmt.Println(len(graph.components))
	var sb strings.Builder
	for i := 0; i < len(graph.components); i++ {
		sb.WriteString(strconv.Itoa(graph.components[i]))
		if i != len(graph.components)-1 {
			sb.WriteString(" ")
		}
	}
	fmt.Print(sb.String())
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
		graph.addEdge(a, b)
	}

	graph.dfs(1)
	sort.Ints(graph.components)
	graph.printRes()
}
