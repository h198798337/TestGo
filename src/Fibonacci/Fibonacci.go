package main

import (
	"fmt"
)

func main() {
	a := Fibonacci(10)
	for _, v := range a {
		fmt.Printf("%v ", v)
	}
}

func Fibonacci(val int) []int {
	x := make([]int, val)
	x[0], x[1] = 1, 1
	for i := 2; i < val; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}
