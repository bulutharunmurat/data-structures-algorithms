package heap_sort

import (
	"data-structures-algorithms/heap"
	"fmt"
)

// BEST CASE (n)log(n) and WORST CASE = (n)log(n)
func heapSort(arr []int) []int {

	maxHeap := heap.Heap{}
	maxHeap.Heapify(arr)

	end := len(maxHeap.Array) - 1
	for end > 0 {
		heap.Swap(&maxHeap, 0, end)
		maxHeap.MoveDown(0, end-1)
		end = end - 1
	}

	return maxHeap.Array
}

func HeapSortTest() {
	fmt.Println("HEAP SORT:")
	heapBefore := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(heapBefore)
	heapAfter := heapSort(heapBefore)
	fmt.Println(heapAfter)
}
