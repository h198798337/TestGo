package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func main() {
	var s stack
	s.push(23)
	s.push(33)
//	fmt.Printf("%d\n", s.pop())
//	fmt.Printf("%d\n", s.pop())
	fmt.Printf("%s\n", s.String())
}

func (s *stack) push(j int) {
	if s.i >= 9 {
		return
	}
	s.data[s.i] = j
	s.i++
}

func (s *stack) pop() (r int) {
	s.i--
	if s.i < 0 {
		return 0
	}
	return s.data[s.i]
}

func (s *stack) String() string {
	rs := ""
	for i := 0; i < s.i; i++ {
		rs += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "] "
	}
	return rs
}
