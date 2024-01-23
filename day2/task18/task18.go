package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	data []int
}

func (d *Deque) pushBack(num int) {
	d.data = append(d.data, num)
}

func (d *Deque) pushFront(num int) {
	newData := []int{num}
	d.data = append(newData, d.data...)
}

func (d *Deque) popFront() (int, bool) {
	if d.size() == 0 {
		return 0, false
	} else {
		el := d.data[0]
		d.data = append(d.data[:0], d.data[1:]...)
		return el, true
	}
}

func (d *Deque) popBack() (int, bool) {
	if d.size() == 0 {
		return 0, false
	}
	x := d.data[d.size()-1]
	d.data = d.data[:d.size()-1]
	return x, true
}

func (d *Deque) front() (int, bool) {
	if d.size() == 0 {
		return 0, false
	} else {
		return d.data[0], true
	}
}

func (d *Deque) back() (int, bool) {
	if d.size() == 0 {
		return 0, false
	}
	x := d.data[d.size()-1]
	return x, true
}

func (d *Deque) size() int {
	return len(d.data)
}

func (d *Deque) clear() {
	d.data = d.data[:0]
}

func printRes(x int, ok bool) {
	if ok {
		fmt.Println(x)
	} else {
		fmt.Println("error")
	}
}

func (d *Deque) toProcessRequest(line string) {
	if line == "front" {
		printRes(d.front())
	} else if line == "back" {
		printRes(d.back())
	} else if line == "pop_back" {
		printRes(d.popBack())
	} else if line == "pop_front" {
		printRes(d.popFront())
	} else if line == "clear" {
		d.clear()
		fmt.Println("ok")
	} else if line == "size" {
		fmt.Println(d.size())
	} else if line == "exit" {
		fmt.Print("bye")
		os.Exit(0)
	} else if string(line[5]) == "f" {
		s := strings.Split(line, " ")
		x, _ := strconv.Atoi(s[1])
		d.pushFront(x)
		fmt.Println("ok")
	} else {
		s := strings.Split(line, " ")
		x, _ := strconv.Atoi(s[1])
		d.pushBack(x)
		fmt.Println("ok")
	}
}

func main() {
	var d Deque

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		d.toProcessRequest(line)
	}
}
