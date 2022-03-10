package heap

import (
	"errors"
	"fmt"
	"log"
)

type Heap struct {
	Array []int
}

func (heap *Heap) add(data int) {

	heap.Array = append(heap.Array, data)

	curIndex := len(heap.Array) - 1
	parentIndex := findParentIndex(curIndex)

	for heap.Array[parentIndex] < heap.Array[curIndex] {
		Swap(heap, parentIndex, curIndex)
		curIndex = parentIndex
		parentIndex = findParentIndex(curIndex)
	}
}

func Swap(heap *Heap, firstIndex int, secondIndex int) {
	parentValue := heap.Array[firstIndex]
	curValue := heap.Array[secondIndex]

	heap.Array[secondIndex] = parentValue
	heap.Array[firstIndex] = curValue
}

func (heap *Heap) remove(data int) {
	remIndex, err := findIndex(heap, data)
	if err != nil {
		log.Fatal(err, data)
	}
	lastIndex := len(heap.Array) - 1

	Swap(heap, remIndex, lastIndex)
	heap.Array = heap.Array[:lastIndex]

	heap.MoveDown(remIndex, lastIndex)
}

func (heap *Heap) MoveDown(fromIndex int, lastIndex int) {
	for fromIndex < lastIndex && (fromIndex*2)+1 < lastIndex {
		leftChild := (fromIndex * 2) + 1
		rightChild := (fromIndex * 2) + 2
		len := len(heap.Array)
		if (leftChild < len && heap.Array[fromIndex] < heap.Array[leftChild]) ||
			(rightChild < len && heap.Array[fromIndex] < heap.Array[rightChild]) {
			if rightChild < len && heap.Array[leftChild] < heap.Array[rightChild] {
				Swap(heap, fromIndex, rightChild)
				fromIndex = rightChild
			} else {
				Swap(heap, fromIndex, leftChild)
				fromIndex = leftChild
			}
		}
	}
}

func findIndex(heap *Heap, data int) (int, error) {
	for i, v := range heap.Array {
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

func (heap *Heap) Heapify(arr []int) {

	// O(n) complexity
	for _, value := range arr {
		heap.add(value)
	}
}

func MaxHeapTest() {

	fmt.Println("--------------- MAX HEAP ---------------")
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

	fmt.Println(maxHeap.Array)
	maxHeap.remove(15)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(5)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(11)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(4)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(10)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(17)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(9)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(3)
	fmt.Println(maxHeap.Array)
	maxHeap.remove(12)
	fmt.Println(maxHeap.Array)
	fmt.Println("--------------- MAX HEAP ---------------")

}
