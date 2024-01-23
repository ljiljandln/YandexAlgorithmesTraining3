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
	for pos > 0 && h.data[pos] > h.data[(pos-1)/2] {
		h.data[pos], h.data[(pos-1)/2] = h.data[(pos-1)/2], h.data[pos]
		pos = (pos - 1) / 2
	}
}

func (h *Heap) pop() int {
	ans := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	pos := 0

	for pos*2+2 < len(h.data) {
		maxIndex := 2*pos + 1
		if h.data[maxIndex+1] > h.data[maxIndex] {
			maxIndex++
		}

		if h.data[pos] < h.data[maxIndex] {
			h.data[pos], h.data[maxIndex] = h.data[maxIndex], h.data[pos]
			pos = maxIndex
		} else {
			break
		}
	}
	h.data = h.data[:len(h.data)-1]

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var h Heap
	var sb strings.Builder

	for ; n > 0; n-- {
		var currType byte
		fmt.Fscan(in, &currType)

		var number int
		if currType == 0 {
			fmt.Fscan(in, &number)
			h.push(number)
		} else {
			number = h.pop()
			sb.WriteString(strconv.Itoa(number))
			if n != 1 {
				sb.WriteString("\n")
			}
		}
	}
	fmt.Print(sb.String())
}
