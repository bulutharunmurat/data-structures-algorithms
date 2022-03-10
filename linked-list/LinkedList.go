package linked_list

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	head *Node
	tail *Node
}

func (linkedList *LinkedList) add(data int, prevData int) {
	prevNode := find(linkedList, prevData)
	currentNode := Node{nextNode: nil, data: data}

	if linkedList.head == nil {
		addToEmptyList(linkedList, &currentNode)
	} else if prevNode == linkedList.tail {
		addToTail(linkedList, &currentNode)
	} else if prevNode == nil {
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

func (linkedList *LinkedList) remove(data int) {

	removeNode := find(linkedList, data)

	if removeNode == nil {
		fmt.Println("Node to remove cannot find in a list")
		return
	}

	prevNode := findPrev(linkedList, data)

	if removeNode == linkedList.head && removeNode == linkedList.tail {
		removeToEmpty(linkedList)
	} else if removeNode == linkedList.head {
		removeHead(linkedList)
	} else if removeNode == linkedList.tail {
		removeTail(linkedList, prevNode)
	} else {
		removeBetween(linkedList, removeNode, prevNode)
	}

}

func findPrev(linkedList *LinkedList, data int) *Node {

	p := linkedList.head

	for p != nil {
		if p.nextNode != nil && p.nextNode.data == data {
			return p
		}
		p = p.nextNode
	}
	return nil
}

func removeToEmpty(linkedlist *LinkedList) {
	linkedlist.head = nil
	linkedlist.tail = nil
}

func removeHead(linkedlist *LinkedList) {
	linkedlist.head = linkedlist.head.nextNode
}

func removeTail(linkedlist *LinkedList, prevNode *Node) {
	linkedlist.tail = prevNode
	prevNode.nextNode = nil
}

func removeBetween(linkedlist *LinkedList, removeNode *Node, prevNode *Node) {
	prevNode.nextNode = removeNode.nextNode
}

func find(linkedList *LinkedList, data int) *Node {

	p := linkedList.head

	for p != nil {
		if p.data == data {
			return p
		}
		p = p.nextNode
	}
	return nil
}

func (l LinkedList) String() string {

	if l.head == nil && l.tail == nil {
		return "[ ]"
	}

	p := l.head
	var s = "[ "

	for p.nextNode != nil {
		s = s + strconv.Itoa(p.data) + ", "
		p = p.nextNode
	}
	return s + strconv.Itoa(p.data) + " ]"
}

func LinkedListTest() {
	var linkedList = LinkedList{
		head: nil,
		tail: nil}

	linkedList.add(2, 0)    // add to empty linked list
	linkedList.add(4, 2)    // add tail
	linkedList.add(1, 1000) // add to head when prev is not found
	linkedList.add(3, 2)    // add to between two data
	linkedList.add(5, 4)    // add to between two data
	fmt.Println(linkedList)

	linkedList.remove(1)
	fmt.Println(linkedList)
	linkedList.remove(5)
	fmt.Println(linkedList)
	linkedList.remove(3)
	fmt.Println(linkedList)
	linkedList.remove(2)
	fmt.Println(linkedList)
	linkedList.remove(4)
	fmt.Println(linkedList)
}
