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

	tree := BinarySearchTree{root: nil}
	tree.add(20)
	tree.add(15)
	tree.add(25)
	tree.add(13)
	tree.add(17)
	tree.add(23)
	tree.add(10)
	tree.add(14)
	tree.add(19)
	tree.add(22)
	tree.add(24)
	tree.add(20)

	tree.printBF()
	tree.printInOrder()
	tree.printPreOrder()
	tree.printPostOrder()

	data := tree.search(15).data
	fmt.Println(data)

	fmt.Println(tree.searchParent(10).data)

	tree.remove(20)
	tree.remove(13)
	tree.remove(24)
	tree.remove(17)
	tree.remove(23)
	tree.printBF()

	maxHeap := Heap{}
	maxHeap.add(15)
	maxHeap.add(10)
	maxHeap.add(12)
	maxHeap.add(5)
	maxHeap.add(4)
	maxHeap.add(3)
	maxHeap.add(17)
	maxHeap.add(11)
	maxHeap.add(9)

	fmt.Println(maxHeap.array)
	maxHeap.remove(15)
	fmt.Println(maxHeap.array)
	maxHeap.remove(5)
	fmt.Println(maxHeap.array)
	maxHeap.remove(11)
	fmt.Println(maxHeap.array)
	maxHeap.remove(4)
	fmt.Println(maxHeap.array)
	maxHeap.remove(10)
	fmt.Println(maxHeap.array)
	maxHeap.remove(17)
	fmt.Println(maxHeap.array)
	maxHeap.remove(9)
	fmt.Println(maxHeap.array)
	maxHeap.remove(3)
	fmt.Println(maxHeap.array)
	maxHeap.remove(12)
	fmt.Println(maxHeap.array)

	a := 5
	fmt.Println(a)

	unSortedList := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(unSortedList)
	sortedList := selectionSort(unSortedList)
	fmt.Println(sortedList)

	fmt.Println("INSERTION SORT:")
	unSortedList2 := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(unSortedList2)
	sortedList2 := insertionSort(unSortedList)
	fmt.Println(sortedList2)

	fmt.Println("BUBBLE SORT:")
	bubbleBefore := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(bubbleBefore)
	bubbleAfter := bubbleSort(bubbleBefore)
	fmt.Println(bubbleAfter)
}
