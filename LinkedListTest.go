package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	var linkedList = LinkedList{}

	linkedList.addWithoutPrev(5)

	fmt.Println(linkedList.head.data)

}
