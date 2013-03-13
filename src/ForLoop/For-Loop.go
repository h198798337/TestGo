package main 

import (
	"fmt"
)

func main() {
	array:=[...]int{0,1,2,3,4,5,6,7,8,9}
	for i:=0;i<len(array);i++{
		if i>2 {
			goto Exit
		}
		fmt.Println(i)
	}
	Exit:
		fmt.Println("Exit!")
}
