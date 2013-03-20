package main

import (
	"fmt"
)

func main() {
	a := []int{1, 23, 45, 5, 67}
	bubbling(a[:])
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
}

func bubbling(a []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(a)-1; i++ {
			if a[i-1] > a[i] {
				temp := a[i-1]
				a[i] = a[i-1]
				a[i-1] = temp
				swapped = true
			}
		}
	}
}
