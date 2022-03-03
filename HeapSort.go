package main

// BEST CASE (n)log(n) and WORST CASE = (n)log(n)
func heapSort(arr []int) []int {

	maxHeap := Heap{}
	maxHeap.heapify(arr)

	end := len(maxHeap.array) - 1
	for end > 0 {
		swap(&maxHeap, 0, end)
		maxHeap.moveDown(0, end-1)
		end = end - 1
	}

	return maxHeap.array
}
