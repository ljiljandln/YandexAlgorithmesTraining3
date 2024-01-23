package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, prev, ans int
	fmt.Fscan(in, &n, &prev)

	for ; n > 1; n-- {
		var curr int
		fmt.Fscan(in, &curr)

		if prev > curr {
			ans += curr
		} else {
			ans += prev
		}
		prev = curr
	}

	fmt.Print(ans)
}
