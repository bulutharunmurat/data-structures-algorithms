package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	var linkedList = LinkedList{head: Node{nextNode: nil}, tail: Node{nextNode: nil}}

	linkedList.addWithoutPrev(5)
	linkedList.addBetween(7, 5)
	linkedList.addBetween(8, 7)
	fmt.Println(linkedList)

}
