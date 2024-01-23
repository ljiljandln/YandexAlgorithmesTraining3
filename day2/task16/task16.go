package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	data []int
}

func (q *Queue) size() int {
	return len(q.data)
}

func (q *Queue) pop() (int, bool) {
	if q.size() == 0 {
		return 0, false
	} else {
		el := q.data[0]
		q.data = append(q.data[:0], q.data[1:]...)
		return el, true
	}
}

func (q *Queue) front() (int, bool) {
	if q.size() == 0 {
		return 0, false
	} else {
		return q.data[0], true
	}
}

func (q *Queue) push(num int) {
	q.data = append(q.data, num)
}

func (q *Queue) clear() {
	q.data = q.data[:0]
}

func printXOrError(x int, ok bool) {
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("error")
	}
}

func main() {
	var q Queue

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "front" {
			printXOrError(q.front())
		} else if line == "pop" {
			printXOrError(q.pop())
		} else if line == "clear" {
			q.clear()
			fmt.Println("ok")
		} else if line == "size" {
			fmt.Println(q.size())
		} else if line == "exit" {
			fmt.Print("bye")
			os.Exit(0)
		} else {
			s := strings.Split(line, " ")
			x, _ := strconv.Atoi(s[1])
			q.push(x)
			fmt.Println("ok")
		}
	}
}
