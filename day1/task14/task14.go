package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	items []byte
}

func (stack *Stack) size() int {
	return len(stack.items)
}

func (stack *Stack) push(number byte) {
	stack.items = append(stack.items, number)
}

func (stack *Stack) pop() byte {
	x := stack.items[stack.size()-1]
	stack.items = stack.items[:stack.size()-1]
	return x
}

func (stack *Stack) back() byte {
	return stack.items[stack.size()-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n byte
	fmt.Fscan(in, &n)

	arr := make([]byte, n)
	for i := byte(0); i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	var st Stack
	var count byte = 0
	for _, number := range arr {
		st.push(number)

		for st.size() != 0 && st.back() == count+1 {
			st.pop()
			count++
		}
	}

	if count == n {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
