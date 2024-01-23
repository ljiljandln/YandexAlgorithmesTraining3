package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	data []uint8
}

func (q *Queue) push(num uint8) {
	q.data = append(q.data, num)
}

func (q *Queue) size() int {
	return len(q.data)
}

func (q *Queue) pop() uint8 {
	el := q.data[0]
	q.data = append(q.data[:0], q.data[1:]...)
	return el
}

type Graph struct {
	n      uint8
	dist   []uint8
	prev   []uint8
	matrix [][]uint8
	q      Queue
}

func getGraph(in *bufio.Reader) *Graph {
	var gr Graph
	fmt.Fscan(in, &gr.n)

	gr.matrix = make([][]uint8, gr.n+1)
	for i := uint8(1); i <= gr.n; i++ {
		gr.matrix[i] = make([]uint8, gr.n+1)
		for j := uint8(1); j <= gr.n; j++ {
			fmt.Fscan(in, &gr.matrix[i][j])
		}
	}
	gr.dist = make([]uint8, gr.n+1)
	gr.prev = make([]uint8, gr.n+1)

	return &gr
}

func (gr *Graph) bfs(v uint8) {
	for i := uint8(1); i <= gr.n; i++ {
		if gr.matrix[v][i] != 0 && gr.dist[i] == 0 {
			gr.dist[i] = gr.dist[v] + 1
			gr.prev[i] = v
			gr.q.push(i)
		}
	}
}

func (gr *Graph) search(start, end uint8) {
	v := start
	gr.bfs(v)
	for gr.q.size() > 0 {
		v = gr.q.pop()
		if v == end {
			return
		}
		gr.bfs(v)
	}
}

func (gr *Graph) getWay(start, end uint8) string {
	var str string
	for end != start {
		str = fmt.Sprintf(" %d%s", end, str)
		end = gr.prev[end]
	}
	return fmt.Sprintf("%d%s", start, str)
}

func (gr *Graph) printRes(start, end uint8) {
	if start != end {
		gr.search(start, end)
		if gr.dist[end] != 0 {
			fmt.Println(gr.dist[end])
			fmt.Print(gr.getWay(start, end))
		} else {
			fmt.Print(-1)
		}
	} else {
		fmt.Print(0)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	gr := getGraph(in)

	var start, end uint8
	fmt.Fscan(in, &start, &end)

	gr.printRes(start, end)
}
