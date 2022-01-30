package main

type LinkedList struct {
	head Node
	tail Node
}

func (linkedList *LinkedList) add(data int, prev int) {
	var currentNode = Node{}

	if &linkedList.head == nil {
		addToEmptyList(linkedList, currentNode)
	}
}

// TODO check if there is method overloading
func (linkedList *LinkedList) addWithoutPrev(givenData int) {
	var currentNode = Node{data: givenData}

	if &linkedList.head == nil {
		addToEmptyList(linkedList, currentNode)
	}
	//addToTail(linkedList, currentNode)
}

func addToEmptyList(linkedList *LinkedList, node Node) {
	linkedList.head = node
	linkedList.tail = node
}

/*
func addToTail(linkedList *LinkedList, node Node) {
	*linkedList.tail.nextNode = node
	linkedList.tail = node
}

*/
