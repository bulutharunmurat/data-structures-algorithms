package merge_sort

import "fmt"

// BEST CASE (n)log(n) and WORST CASE = (n)log(n)
func mergeSort(arr []int) []int {
	var right, left, result []int

	if len(arr) <= 1 {
		return arr
	} else {
		mid := len(arr) / 2
		left = arr[:mid]
		right = arr[mid:]

		left = mergeSort(left)
		right = mergeSort(right)

		result = merge(left, right)
		return result
	}
}

func merge(left []int, right []int) []int {
	var result []int

	curLeft := 0
	curRight := 0
	for curLeft < len(left) && curRight < len(right) {
		if left[curLeft] < right[curRight] {
			result = append(result, left[curLeft])
			curLeft++
		} else {
			result = append(result, right[curRight])
			curRight++
		}
	}

	for curLeft < len(left) {
		result = append(result, left[curLeft])
		curLeft++
	}

	for curRight < len(right) {
		result = append(result, right[curRight])
		curRight++
	}

	return result
}

func MergeSortTest() {
	fmt.Println("MERGE SORT:")
	list := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(list)
	list = mergeSort(list)
	fmt.Println(list)
}
