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

	fmt.Println("deneme")
	s.push(5)
	s.push(6)
	s.push(7)
	s.push(8)
	s.push(9)
	s.push(10)
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()
	s.pop()

	fmt.Println(s)

}
