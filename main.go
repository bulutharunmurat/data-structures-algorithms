package main

import (
	"fmt"
)

func main() {
	// 0xc000010280 address
	// to take an address for variable we use &var
	// pointers are that store addresses
	// var pointerForInt *int, pointerForInt = &minutes
	// *pointerForStr changing of value by pointers -> dereferencing or indirecting.

	var node = Node{}

	var linkedList = LinkedList{
		head: &node,
		tail: &node}

	linkedList.add(10, 0)    // add to empty linked list
	linkedList.add(5, 10)    // add tail
	linkedList.add(555, 44)  // add to head when prev is not found
	linkedList.add(888, 555) // add to between two data

	fmt.Println(linkedList)
	fmt.Println("----------STACK----------")
	var s = MyStack{top: -1, capacity: 5}

	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)
	s.push(9)
	peek := s.peek()
	fmt.Println("Peeked: ", peek)
	fmt.Println(s)

	fmt.Println("----------QUEUE----------")
	q := MyQueue{capacity: 5, front: 0, rear: -1, count: 0}
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	q.enqueue(7)
	q.enqueue(8)
	dequeue := q.dequeue()
	fmt.Println(dequeue)
	fmt.Println(q)

}
