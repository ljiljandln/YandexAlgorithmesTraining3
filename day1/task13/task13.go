package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func (stack *Stack) size() int {
	return len(stack.items)
}

func (stack *Stack) push(number int) {
	stack.items = append(stack.items, number)
}

func (stack *Stack) pop() int {
	x := stack.items[stack.size()-1]
	stack.items = stack.items[:stack.size()-1]
	return x
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1*1024*1024*1024)
	in.Buffer(buf, 1*1024*1024*1024)
	in.Scan()

	line := in.Text()
	leksems := strings.Fields(line)

	var st Stack
	for _, curr := range leksems {
		if curr == "+" || curr == "-" || curr == "*" {
			b := st.pop()
			a := st.pop()
			var c int
			if curr == "+" {
				c = a + b
			} else if curr == "-" {
				c = a - b
			} else {
				c = a * b
			}
			st.push(c)
		} else {
			x, _ := strconv.Atoi(curr)
			st.push(x)
		}
	}
	fmt.Print(st.pop())
}
