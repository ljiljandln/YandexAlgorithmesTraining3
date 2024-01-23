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

func (stack *Stack) pop() (int, bool) {
	if stack.size() == 0 {
		return 0, false
	}
	x := stack.items[stack.size()-1]
	stack.items = stack.items[:stack.size()-1]
	return x, true
}

func (stack *Stack) back() (int, bool) {
	if stack.size() == 0 {
		return 0, false
	}
	x := stack.items[stack.size()-1]
	return x, true
}

func (stack *Stack) clear() {
	stack.items = stack.items[:0]
}

func printXOrError(x int, ok bool) {
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("error")
	}
}

func main() {
	var st Stack

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "back" {
			printXOrError(st.back())
		} else if line == "pop" {
			printXOrError(st.pop())
		} else if line == "clear" {
			st.clear()
			fmt.Println("ok")
		} else if line == "size" {
			fmt.Println(st.size())
		} else if line == "exit" {
			fmt.Print("bye")
			os.Exit(0)
		} else {
			s := strings.Split(line, " ")
			x, _ := strconv.Atoi(s[1])
			st.push(x)
			fmt.Println("ok")
		}
	}
}
