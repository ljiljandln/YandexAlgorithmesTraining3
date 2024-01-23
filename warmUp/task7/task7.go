package main

import (
	"bufio"
	"fmt"
	"os"
)

func getSec(in *bufio.Reader) int {
	var h, m, s int
	fmt.Fscanf(in, "%d:%d:%d\n", &h, &m, &s)
	m += h * 60
	s += m * 60
	return s
}

func getDiff(s1, s3 int) int {
	if s3 > s1 {
		return (s3 - s1 + 1) / 2
	} else {
		return (s3 + (86400 - s1) + 1) / 2
	}
}

func printAns(ansSec int) {
	h := ansSec / 3600
	mAndSec := ansSec - h*3600
	m := mAndSec / 60
	s := mAndSec - m*60

	fmt.Printf("%02d:%02d:%02d", h, m, s)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	s1, s2, s3 := getSec(in), getSec(in), getSec(in)

	diff := getDiff(s1, s3)

	printAns((s2 + diff) % 86400)
}
