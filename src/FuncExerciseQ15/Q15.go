package main 

import (
	"fmt"
)

func main() {
	p:=plusTwo()
	fmt.Printf("%d\n",p(2))
	
	p2:=plusX(3)
	fmt.Printf("%d\n",p2(2))
}

func plusTwo() func(int) int {
	return func(v int) int {
		return v + 2
	}
}

func plusX(x int) func(int) int {
	return func(v int) int {
		return v + x
	}
}
