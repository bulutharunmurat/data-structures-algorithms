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

	heap.moveDown(remIndex, lastIndex)
}

func (heap *Heap) moveDown(fromIndex int, lastIndex int) {
	for fromIndex < lastIndex && (fromIndex*2)+1 < lastIndex {
		leftChild := (fromIndex * 2) + 1
		rightChild := (fromIndex * 2) + 2
		len := len(heap.array)
		if (leftChild < len && heap.array[fromIndex] < heap.array[leftChild]) ||
			(rightChild < len && heap.array[fromIndex] < heap.array[rightChild]) {
			if rightChild < len && heap.array[leftChild] < heap.array[rightChild] {
				swap(heap, fromIndex, rightChild)
				fromIndex = rightChild
			} else {
				swap(heap, fromIndex, leftChild)
				fromIndex = leftChild
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

func (heap *Heap) heapify(arr []int) {

	// O(n) complexity
	for _, value := range arr {
		heap.add(value)
	}
}
