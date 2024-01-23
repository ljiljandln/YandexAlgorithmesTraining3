package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkExtremum(number int, min, max *int) {
	if number < *min {
		*min = number
	} else if number > *max {
		*max = number
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var xMin, yMin int
	fmt.Fscan(in, &xMin, &yMin)
	xMax, yMax := xMin, yMin

	for ; n > 1; n-- {
		var x, y int
		fmt.Fscan(in, &x, &y)
		checkExtremum(x, &xMin, &xMax)
		checkExtremum(y, &yMin, &yMax)
	}

	fmt.Printf("%d %d %d %d", xMin, yMin, xMax, yMax)
}
