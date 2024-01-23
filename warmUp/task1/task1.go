package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func addStr(keys *[]int, symbolsMap *map[int]int, max *int, str string) {
	_map := *symbolsMap
	for i := range str {
		symbol := int(str[i])
		if symbol > 32 && symbol < 123 {
			if val, ok := _map[symbol]; ok {
				_map[symbol] = val + 1
			} else {
				_map[symbol] = 1
				*keys = append(*keys, symbol)
			}

			if _map[symbol] > *max {
				*max = _map[symbol]
			}
		}
	}
}

func printRes(keys []int, symbolsMap map[int]int, max int) {
	var sb strings.Builder
	for step := 0; step < max; step++ {
		for _, symbol := range keys {
			val := symbolsMap[symbol]
			if val >= max-step {
				sb.WriteString("#")
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())

	for _, symbol := range keys {
		fmt.Printf("%c", symbol)
	}
}

func main() {

	var keys []int
	var symbolsMap = map[int]int{}

	max := -1

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 100000)
	scanner.Buffer(buf, 100000)

	for {
		ok := scanner.Scan()
		if !ok {
			break
		}
		line := scanner.Text()
		addStr(&keys, &symbolsMap, &max, line)
	}
	sort.Ints(keys)

	printRes(keys, symbolsMap, max)
}
