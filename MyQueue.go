package main

import "strconv"

type MyQueue struct {
	front    int
	rear     int
	capacity int
	array    []int
	count    int
}

func (queue *MyQueue) isEmpty() bool {
	return queue.count == 0
}

func (queue *MyQueue) isFull() bool {
	return queue.count == queue.capacity
}

func (queue *MyQueue) enqueue(data int) {
	if !queue.isFull() {
		queue.count = queue.count + 1
		queue.rear = (queue.rear + 1) % queue.capacity
		queue.array = append(queue.array, data)
	} else {
		panic("Queue is full!!")
	}
}
func (queue *MyQueue) dequeue() int {
	if !queue.isEmpty() {
		queue.count = queue.count - 1
		i := queue.array[queue.front]
		queue.front = (queue.front + 1) % queue.capacity
		return i
	} else {
		panic("Queue is empty!!")
	}
}
func (queue *MyQueue) clear() {
	queue.front = 0
	queue.rear = -1
}

func (queue MyQueue) String() string {

	s := "[ "

	for i := 0; i < queue.count; i++ {
		s = s + strconv.Itoa(queue.array[(queue.front+i)%queue.capacity]) + ", "
	}

	s = s + " ]"
	return s
}
