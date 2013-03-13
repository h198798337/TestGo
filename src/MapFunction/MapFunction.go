package main

import (
	"fmt"
)

func main() {
	f := func(v int) int {
		return v * v
	}
	a := [...]int{2, 5, 6, 9}
	mapFunction(f, a[:])
	fmt.Printf("%d ", a)
}

func mapFunction(f func(int) int, a []int) []int{
	for p, v := range a {
		a[p] = f(v)
	}
	return a
}
