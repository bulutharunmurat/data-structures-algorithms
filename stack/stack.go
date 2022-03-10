package stack

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
		panic("Stack is full, so push is not executed")
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

func (stack *MyStack) clear() {
	stack.top = -1
}

func (stack *MyStack) peek() int {
	return stack.array[stack.top]
}

func (stack MyStack) String() string {

	s := "[ "

	for i := 0; i <= stack.top; i++ {
		s = s + strconv.Itoa(stack.array[i]) + ", "
	}

	s = s + " ]"
	return s
}

func StackTest() {
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
}
