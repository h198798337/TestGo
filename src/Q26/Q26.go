package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go slow(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func slow(ch chan int) {
	for {
		i := <-ch
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
