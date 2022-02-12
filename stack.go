package main

import (
	"fmt"
	"strconv"
)

type MyStack struct {
	top      int
	capacity int
	array    []int
}

func (stack *MyStack) push(data int) {

	if !stack.isFull() {
		stack.top = stack.top + 1
		stack.array = append(stack.array, data)
	} else {
		fmt.Println("Stack is full, so push is not executed for data: ", data)
	}

}

func (stack *MyStack) pop() int {
	if !stack.isEmpty() {
		pop := stack.array[stack.top]
		stack.top = stack.top - 1
		return pop
	}
	panic("Stack is empty so pop can not executed")
}

func (stack *MyStack) isFull() bool {
	return stack.top == stack.capacity-1
}

func (stack *MyStack) isEmpty() bool {
	return stack.top == -1
}

func (stack MyStack) String() string {

	s := " "

	for i := 0; i <= stack.top; i++ {
		s = s + strconv.Itoa(stack.array[i]) + ", "
	}
	return s
}
