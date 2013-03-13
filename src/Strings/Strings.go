package main 

import (
	"fmt"
	"unicode/utf8"
)

func main() {
//	goStrings()
//	RunCount()
//	replace()
	reverse()
}

func goStrings() {
	var str string;
	str = "A"
	for i:=0;i<100;i++ {
		fmt.Println(str)
		str += "A"	
	}
}

func RunCount() {
	str:="asSASA ddd dsjkdsjs dk"
	fmt.Printf("String %s\nLenght: %d, Runes: %d\n",str,len([]byte(str)),utf8.RuneCount([]byte(str)))
}

func replace() {
	str:="asSASA ddd dsjkdsjs dk"
	b:=[]byte(str)
	b[4] = 'a'
	b[5] = 'b'
	b[6] = 'c'
	fmt.Printf("%s",string(b));
}

func reverse() {
	str:="foobar";
	a:=[]byte(str)
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1 {
		a[i],a[j]=a[j],a[i]
	}
	fmt.Printf("%s\n",string(a));
}


