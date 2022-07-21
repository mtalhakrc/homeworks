package main

import "fmt"

type node struct {
	data int
	next *node
}
type stack struct {
	head *node
	len  int
}

func main() {
	var list = stack{}
	list.push(10)
	fmt.Println(list.head)
	/*
		list.push(20)
		list.push(30)
		fmt.Println(list.pop())
		fmt.Println(list.pop())
		list.push(50)
		fmt.Println(list.pop())
		fmt.Println(list.pop())
	*/
}

func (list *stack) push(data int) {
	newNode := &node{
		data: data,
		next: nil,
	}
	if list.len == 0 {
		list.head = newNode
		list.len++
		return
	}
	newNode.next = list.head
	list.head = newNode
	list.len++
}

func (list *stack) pop() int {
	if list.len == 0 {
		return -99999
	}
	head := list.head.data
	list.head = list.head.next
	list.len--
	return head
}
