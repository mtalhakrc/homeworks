package main

import "fmt"

type node struct {
	data interface{}
	next *node
}
type queue struct {
	head *node
	last *node
	len  int
}

func main() {
	var list = queue{}
	list.enqueue("at")
	list.enqueue("kafası")
	list.enqueue("abdulsamet")
	list.enqueue(10)

	list.print()

	fmt.Printf("1- %v\n", list.dequeue())
	fmt.Printf("2- %v\n", list.dequeue())
	list.enqueue(10)
	fmt.Printf("3- %v\n", list.dequeue())
	fmt.Printf("4- %v\n", list.dequeue())
	fmt.Printf("5- %v\n", list.dequeue())
}

func (list *queue) enqueue(data interface{}) {
	newNode := &node{
		data: data,
		next: nil,
	}
	if list.len == 0 {
		list.head = newNode
		list.last = newNode
		list.len++
		return
	}

	list.last.next = newNode
	list.last = newNode
	list.len++
}

func (list *queue) dequeue() interface{} {
	if list.len == 0 {
		return nil
	}
	head := list.head
	list.head = list.head.next
	list.len--

	return head.data
}

func (list queue) print() {
	if list.head == nil {
		return
	}
	for ; list.head != nil; list.head = list.head.next {
		fmt.Println(list.head.data)
	}
}

//func (list queue) last() *node {
//	for ; list.head != nil && list.head.next != nil; list.head = list.head.next {
//	}
//	return list.head
//}
