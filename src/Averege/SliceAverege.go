package main 

import (
	"fmt"
)

func main() {
	a:=[...]float64{1.2,2.3,4.5,3.4,17.2}
	s:=a[:]
	ave:=averger(s)
	fmt.Printf("%f\n",ave)
}

func averger(array []float64)(ave float64) {
	sum:=0.0
	switch len(array) {
		case 0:
			ave=0.0
		default:
			for _,v:=range array {
				sum+=v
			}
			ave=sum/float64(len(array))
	}
	return
}

