package main

import (
	"fmt"
)

func main() {
	slice := [...]float64{4.5, 4.8, 8.9}
	sum := 0.0
	ave := 0.0
	switch len(slice) {
	case 0:
		ave = 0
	default:
		for _, v := range slice {
			sum += v
		}
		ave = sum / float64(len(slice))
	}
	fmt.Printf("%f\n", ave)

}
