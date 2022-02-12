package main

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	head Node
	tail Node
}

func (linkedList *LinkedList) addBetween(data int, prev int) {
	prevNode := findPrev(linkedList, prev)
	var currentNode = Node{nextNode: nil, data: data}
	prevNode.nextNode = &currentNode

}

// TODO check if there is method overloading
func (linkedList *LinkedList) addWithoutPrev(givenData int) {

	var currentNode = Node{nextNode: nil, data: givenData}
	addToEmptyList(linkedList, currentNode)

}

func addToEmptyList(linkedList *LinkedList, node Node) {
	linkedList.head = node
	linkedList.tail = node
}

func findPrev(linkedList *LinkedList, data int) *Node {

	p := &(linkedList.head)
	nilNode := Node{}
	fmt.Println(p)
	fmt.Println(*p)
	for *p != nilNode {
		if p.data == data {
			return p
		}
		p = p.nextNode
	}

	return &nilNode

}

func (l LinkedList) String() string {

	p := l.head
	var s = " "

	for p.nextNode != nil {
		s = s + strconv.Itoa(p.data) + ", "
		p = *(p.nextNode)
	}
	return s + strconv.Itoa(p.data)
}

/*
func addToTail(linkedList *LinkedList, node Node) {
	*linkedList.tail.nextNode = node
	linkedList.tail = node
}

*/
