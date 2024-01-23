package main

import (
	"bufio"
	"fmt"
	"os"
)

type Segment struct {
	left, right int
}

func (s *Segment) hasInter(s1 *Segment) bool {
	return s.left <= s1.right && s1.left <= s.right
}

type OS struct {
	count    int
	systems  []Segment
	dontWork []bool
}

func (os *OS) addSegment(seg *Segment, index int) {
	for i, currSeg := range os.systems {
		if !os.dontWork[i] && currSeg.hasInter(seg) {
			os.dontWork[i] = true
			os.count--
		}
	}
	os.systems[index] = *seg
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(in, &m, &n)

	os := OS{count: n, systems: make([]Segment, n), dontWork: make([]bool, n)}

	for i := 0; i < n; i++ {
		var seg Segment
		fmt.Fscan(in, &seg.left, &seg.right)
		os.addSegment(&seg, i)
	}

	fmt.Print(os.count)
}
