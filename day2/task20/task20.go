package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

func (h *Heap) push(number int) {
	h.data = append(h.data, number)
	pos := len(h.data) - 1
	for pos > 0 && h.data[pos] < h.data[(pos-1)/2] {
		h.data[pos], h.data[(pos-1)/2] = h.data[(pos-1)/2], h.data[pos]
		pos = (pos - 1) / 2
	}
}

func (h *Heap) pop() int {
	ans := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	pos := 0

	for pos*2+2 < len(h.data) {
		minIndex := 2*pos + 1
		if h.data[minIndex+1] < h.data[minIndex] {
			minIndex++
		}

		if h.data[pos] > h.data[minIndex] {
			h.data[pos], h.data[minIndex] = h.data[minIndex], h.data[pos]
			pos = minIndex
		} else {
			break
		}
	}
	h.data = h.data[:len(h.data)-1]

	return ans
}

func (h *Heap) printHeap() {
	n := len(h.data)
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(h.pop()))
		if i != n-1 {
			sb.WriteString(" ")
		}
	}
	fmt.Print(sb.String())
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var h Heap
	for i := 0; i < n; i++ {
		var number int
		fmt.Fscan(in, &number)

		h.push(number)
	}

	h.printHeap()
}
