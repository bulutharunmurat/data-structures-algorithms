package main

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

func swap(heap *Heap, parentIndex int, curIndex int) {
	parentValue := heap.array[parentIndex]
	curValue := heap.array[curIndex]

	heap.array[curIndex] = parentValue
	heap.array[parentIndex] = curValue
}

func findParentIndex(curIndex int) int {
	return (curIndex - 1) / 2
}