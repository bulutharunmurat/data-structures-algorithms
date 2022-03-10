package insertion_sort

import "fmt"

// BEST CASE n and WORST CASE = n(n+1)/2

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		tempIndex := i
		for tempIndex >= 1 && arr[tempIndex-1] > arr[tempIndex] {

			smaller := arr[tempIndex]

			arr[tempIndex] = arr[tempIndex-1]
			arr[tempIndex-1] = smaller

			tempIndex = tempIndex - 1
		}
	}

	return arr

}
func InsertionSortTest() {
	fmt.Println("INSERTION SORT:")
	unSortedList2 := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(unSortedList2)
	sortedList2 := insertionSort(unSortedList2)
	fmt.Println(sortedList2)
}
