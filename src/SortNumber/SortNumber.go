package main

import (
	"fmt"
)

func main() {
	i, j := 8, 3
	fmt.Printf("%d,%d\n", i, j)
	k, o := sort(i, j)
	fmt.Printf("%d,%d", k, o)
}

func sort(n1, n2 int) (r1, r2 int) {
	if n1 >= n2 {
		r1 = n2
		r2 = n1
	} else {
		r1 = n1
		r2 = n2
	}
	return
}
