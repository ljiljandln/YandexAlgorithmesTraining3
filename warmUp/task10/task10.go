package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1*1024*1024*1024)
	in.Buffer(buf, 1*1024*1024*1024)
	in.Scan()

	line := in.Text()

	var symbolsMap = map[int]int{}
	var keys []int

	for i := range line {
		symbol := int(line[i])
		ans := (i + 1) * (len(line) - i)
		if val, ok := symbolsMap[symbol]; ok {
			symbolsMap[symbol] = val + ans
		} else {
			symbolsMap[symbol] = ans
			keys = append(keys, symbol)
		}
	}

	sort.Ints(keys)

	for _, symbol := range keys {
		fmt.Printf("%c: %d\n", symbol, symbolsMap[symbol])
	}
}
