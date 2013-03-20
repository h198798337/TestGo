package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 44, 5, 69, 10}
	//	fmt.Printf("Max:%d\n", getMax(a[:]))
	//	fmt.Printf("Min:%d\n", getMin(a[:]))
	fmt.Printf("Max:%d\n", max(a[:]))
}

func max(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v > m {
			m = v
		}
	}
	return m
}

/*func getMax(slice []int) int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] < slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice[0]
}

func getMin(slice []int) int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice[0]
}sb做法*/
