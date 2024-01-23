package main

import (
	"fmt"
	"os"
)

type Stack struct {
	items []rune
}

func (stack *Stack) size() int {
	return len(stack.items)
}

func (stack *Stack) push(symbol rune) {
	stack.items = append(stack.items, symbol)
}

func (stack *Stack) pop() (rune, bool) {
	if stack.size() == 0 {
		return 0, false
	}
	x := stack.items[stack.size()-1]
	stack.items = stack.items[:stack.size()-1]
	return x, true
}

func main() {
	var st Stack
	var line string
	fmt.Scanf("%s\n", &line)

	for _, sym := range line {
		if sym == '(' {
			st.push(')')
		} else if sym == '[' {
			st.push(']')
		} else if sym == '{' {
			st.push('}')
		} else {
			x, ok := st.pop()
			if !ok || x != sym {
				fmt.Print("no")
				os.Exit(0)
			}
		}
	}

	if st.size() == 0 {
		fmt.Print("yes")
	} else {
		fmt.Print("no")
	}
}
