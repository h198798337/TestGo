package main

import (
	"fmt"
)

type Xi []int

type Compare interface {
	Len() int
	Less(i int, j interface{})
	More(i int, j interface{})
}

func (x Xi) Len() int {
	return len(x)
}

func (x Xi) Less(i int, j interface{}) {
	point := j.(*int)
	if x[i] < *point {
		*point = x[i]
	}
}

func (x Xi) More(i int, j interface{}) {
	point := j.(*int)
	if x[i] > *point {
		*point = x[i]
	}
}

func Min(c Compare, j interface{}){
	for i := 0; i < c.Len(); i++ {
		c.Less(i, j)
	}
}

func main() {
	xi := Xi{10, 2, 10, 33}
	var i int = xi[0]
	Min(xi, &i)
	fmt.Printf("%d\n", i)
}
