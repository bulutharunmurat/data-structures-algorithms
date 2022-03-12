package main

import (
	binary_search_tree "data-structures-algorithms/binary-search-tree"
	bubble_sort "data-structures-algorithms/bubble-sort"
	"data-structures-algorithms/heap"
	"data-structures-algorithms/heap-sort"
	insertion_sort "data-structures-algorithms/insertion-sort"
	linked_list "data-structures-algorithms/linked-list"
	merge_sort "data-structures-algorithms/merge-sort"
	"data-structures-algorithms/queue"
	selection_sort "data-structures-algorithms/selection-sort"
	"data-structures-algorithms/stack"
)

func main() {
	// 0xc000010280 address
	// to take an address for variable we use &var
	// pointers are that store addresses
	// var pointerForInt *int, pointerForInt = &minutes
	// *pointerForStr changing of value by pointers -> dereferencing or indirecting.
	linked_list.LinkedListTest()
	stack.StackTest()
	queue.QueueTest()
	binary_search_tree.BinarySearchTreeTest()
	heap.MaxHeapTest()
	selection_sort.SelectionSortTest()
	bubble_sort.BubbleSortTest()
	insertion_sort.InsertionSortTest()
	heap_sort.HeapSortTest()
	merge_sort.MergeSortTest()
}
