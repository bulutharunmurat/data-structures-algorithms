package main

import (
	"errors"
	"log"
)

type Heap struct {
	array []int
}

func (heap *Heap) add(data int) {

	heap.array = append(heap.array, data)

	curIndex := len(heap.array) - 1
	parentIndex := findParentIndex(curIndex)

	for heap.array[parentIndex] < heap.array[curIndex] {
		swap(heap, parentIndex, curIndex)
		curIndex = parentIndex
		parentIndex = findParentIndex(curIndex)
	}
}

func swap(heap *Heap, firstIndex int, secondIndex int) {
	parentValue := heap.array[firstIndex]
	curValue := heap.array[secondIndex]

	heap.array[secondIndex] = parentValue
	heap.array[firstIndex] = curValue
}

func (heap *Heap) remove(data int) {
	remIndex, err := findIndex(heap, data)
	if err != nil {
		log.Fatal(err, data)
	}
	lastIndex := len(heap.array) - 1

	swap(heap, remIndex, lastIndex)
	heap.array = heap.array[:lastIndex]

	for remIndex < lastIndex && (remIndex*2)+1 < lastIndex {
		leftChild := (remIndex * 2) + 1
		rightChild := (remIndex * 2) + 2
		len := len(heap.array)
		if (leftChild < len && heap.array[remIndex] < heap.array[leftChild]) ||
			(rightChild < len && heap.array[remIndex] < heap.array[rightChild]) {
			if rightChild < len && heap.array[leftChild] < heap.array[rightChild] {
				swap(heap, remIndex, rightChild)
				remIndex = rightChild
			} else {
				swap(heap, remIndex, leftChild)
				remIndex = leftChild
			}
		}
	}
}

func findIndex(heap *Heap, data int) (int, error) {
	for i, v := range heap.array {
		if v == data {
			return i, nil
		}
	}
	err := errors.New("data not found in array")
	return 0, err
}

func findParentIndex(curIndex int) int {
	return (curIndex - 1) / 2
}
