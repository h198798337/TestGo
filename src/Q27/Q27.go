package main

import (
	"fmt"
)

func dup3(in chan int) (chan int, chan int, chan int) {
	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 1
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
}

/*func main() {
	ch:=make(chan int)

	go Fibonacci(10,ch)
	var more bool = true
    var val int
    for more {
        select{
        //channel 会返回两个值，一个是内容，一个是还有没有内容
        case val, more = <- ch:
            if more {
                fmt.Println (val)
            }else{
                fmt.Println ("channel closed!")
            }
        }
    }
}

func Fibonacci(val int, ch chan int) {
	x := make([]int, val)
	x[0], x[1] = 1, 1
	ch<-x[0]
	ch<-x[1]
	for i := 2; i < val; i++ {
		x[i] = x[i-1] + x[i-2]
		ch<-x[i]
	}
	close(ch)
}个人做法*/
