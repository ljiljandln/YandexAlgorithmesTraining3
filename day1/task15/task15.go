package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	cost, index int
}

type Stack struct {
	items []Pair
}

func (stack *Stack) size() int {
	return len(stack.items)
}

func (stack *Stack) push(cost, index int) {
	stack.items = append(stack.items, Pair{cost, index})
}

func (stack *Stack) pop() int {
	x := stack.items[stack.size()-1].index
	stack.items = stack.items[:stack.size()-1]
	return x
}

func (stack *Stack) back() int {
	return stack.items[stack.size()-1].cost
}

func printRes(ans *[]int) {
	var sb strings.Builder
	for _, index := range *ans {
		if index == 0 {
			sb.WriteString(strconv.Itoa(-1))
		} else {
			sb.WriteString(strconv.Itoa(index))
		}
		sb.WriteString(" ")
	}
	fmt.Print(sb.String())
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	ans := make([]int, n)

	var st Stack
	for index := 0; index < n; index++ {
		var cost int
		fmt.Fscan(in, &cost)

		for st.size() != 0 && st.back() > cost {
			index1 := st.pop()
			ans[index1] = index
		}
		st.push(cost, index)
	}

	printRes(&ans)
}
