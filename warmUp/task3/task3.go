package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func binSearch(numbers *[]int, search int) int {
	_numbers := *numbers
	var res, left, right = 0, 0, len(_numbers) - 1

	for left <= right {
		middle := (left + right) / 2
		x := _numbers[middle]

		if x < search {
			res = middle + 1
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return res
}

func getRes(m, _min, _max int, in *bufio.Reader, numbers *[]int) {
	resMap := make(map[int]int)
	var sb strings.Builder
	for i := 0; i < m; i++ {
		var curr int
		fmt.Fscan(in, &curr)

		if curr <= _min {
			sb.WriteString("0")
		} else if curr > _max {
			sb.WriteString(strconv.Itoa(len(*numbers)))
		} else {
			value, exist := resMap[curr]
			if exist {
				sb.WriteString(strconv.Itoa(value))
			} else {
				sb.WriteString(strconv.Itoa(binSearch(numbers, curr)))
			}
		}
		if i != m-1 {
			sb.WriteString("\n")
		}
	}
	fmt.Print(sb.String())
}

func main() {
	var n, m, curr int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	nMap := make(map[int]struct{})
	var numbers []int

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &curr)
		if _, exist := nMap[curr]; !exist {
			nMap[curr] = struct{}{}
			numbers = append(numbers, curr)
		}
	}
	sort.Ints(numbers)
	var _min, _max = numbers[0], numbers[len(numbers)-1]

	fmt.Fscan(in, &m)

	getRes(m, _min, _max, in, &numbers)
}
