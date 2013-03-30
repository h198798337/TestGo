package stack

import (
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

//func main() {
//	var s stack
//	s.push(23)
//	s.push(33)
////	fmt.Printf("%d\n", s.pop())
////	fmt.Printf("%d\n", s.pop())
//	fmt.Printf("%s\n", s.String())
//}

func (s *Stack) Push(j int) {
	if s.i >= 9 {
		return
	}
	s.data[s.i] = j
	s.i++
}

func (s *Stack) Pop() (r int) {
	s.i--
	if s.i < 0 {
		return 0
	}
	return s.data[s.i]
}

func (s *Stack) String() string {
	rs := ""
	for i := 0; i < s.i; i++ {
		rs += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "] "
	}
	return rs
}
