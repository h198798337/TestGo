package main

import (
	//	"container/list"
	"fmt"
)

type e interface{}

type node struct {
	data e
	next *node
}

type cumlist struct {
	count int
	front *node
	rear  *node
}

func (cl *cumlist) push(item e) {
	n := new(node)
	n.data = item
	if cl.front == nil {
		n.next = nil
		cl.front = n
		cl.rear = n
		cl.count++
		return
	}
	n.next = cl.front
	cl.front = n
	cl.count++
}

func (cl *cumlist) get_front() *node {
	return cl.front
}

func main() {
	/*mylist := list.New()
	mylist.PushFront(1)
	mylist.PushFront(2)
	mylist.PushFront(4)
	for e := mylist.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		fmt.Printf("%v\n", e.Value)
	}*/
	mylist := new(cumlist)
	mylist.push(1)
	mylist.push(2)
	mylist.push(4)
	for e := mylist.get_front(); e != nil; e = e.next {
		// do something with e.Value
		fmt.Printf("%v\n", e.data)
	}
}
