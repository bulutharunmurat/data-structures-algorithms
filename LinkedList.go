package main

import (
	"strconv"
)

type LinkedList struct {
	head *Node
	tail *Node
}

func (linkedList *LinkedList) add(data int, prevData int) {
	prevNode := find(linkedList, prevData)
	currentNode := Node{nextNode: nil, data: data}
	nilNode := Node{}

	if *linkedList.head == nilNode {
		addToEmptyList(linkedList, &currentNode)
	} else if prevNode == linkedList.tail { // prevNode == &linkedList.tail why that is false?
		addToTail(linkedList, &currentNode)
	} else if *prevNode == nilNode {
		addToHead(linkedList, &currentNode)
	} else {
		addBetween(linkedList, data, prevData)
	}
}

func addBetween(linkedList *LinkedList, data int, prevData int) {
	prevNode := find(linkedList, prevData)
	var currentNode = Node{nextNode: nil, data: data}
	currentNode.nextNode = prevNode.nextNode
	prevNode.nextNode = &currentNode
}

// TODO check if there is method overloading
func (linkedList *LinkedList) addWithoutPrev(givenData int) {

	var currentNode = Node{nextNode: nil, data: givenData}
	addToEmptyList(linkedList, &currentNode)

}

func addToEmptyList(linkedList *LinkedList, node *Node) {
	linkedList.head = node
	linkedList.tail = node
}

func addToTail(linkedList *LinkedList, node *Node) {
	linkedList.tail.nextNode = node
	linkedList.tail = node
}

func addToHead(linkedList *LinkedList, node *Node) {
	node.nextNode = linkedList.head
	linkedList.head = node
}

func find(linkedList *LinkedList, data int) *Node {

	p := linkedList.head
	nilNode := Node{}

	for p != nil {
		if p.data == data {
			return p
		}
		p = p.nextNode
	}

	return &nilNode

}

func (l LinkedList) String() string {

	p := l.head
	var s = "[ "

	for p.nextNode != nil {
		s = s + strconv.Itoa(p.data) + ", "
		p = p.nextNode
	}
	return s + strconv.Itoa(p.data) + " ]"
}
