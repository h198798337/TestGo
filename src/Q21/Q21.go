package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n')
		if e != nil {
			fmt.Printf("%s\n", "退出")
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.OpenFile(flag.Arg(i), os.O_RDONLY, 0)
		if e != nil {
			fmt.Fprint(os.Stderr, "%s:error reading from %s:%s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
		defer f.Close()
	}
}
