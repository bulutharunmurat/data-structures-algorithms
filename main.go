package main

import (
	"fmt"
)

func main() {

	var linkedList = LinkedList{
		head: Node{nextNode: nil},
		tail: Node{nextNode: nil}}

	linkedList.addWithoutPrev(5)
	linkedList.addBetween(7, 5)
	linkedList.addBetween(8, 7)
	fmt.Println(linkedList)

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
