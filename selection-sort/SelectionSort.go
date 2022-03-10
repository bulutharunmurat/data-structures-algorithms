package selection_sort

import "fmt"

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// swap
		minValue := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = minValue

	}
	return arr
}

func SelectionSortTest() {
	unSortedList := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(unSortedList)
	sortedList := selectionSort(unSortedList)
	fmt.Println(sortedList)
}
