package bubble_sort

import "fmt"

// BEST CASE n and WORST CASE = n(n+1)/2

func bubbleSort(arr []int) []int {

	flag := true
	for flag {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {

				smaller := arr[i+1]
				greater := arr[i] // this is optional, for simplicity

				arr[i+1] = greater
				arr[i] = smaller

				flag = true
			}

		}
	}

	return arr
}

func BubbleSortTest() {
	fmt.Println("BUBBLE SORT:")
	bubbleBefore := []int{12, 4, 5, 7, 8, 234, 1, 8, 0}
	fmt.Println(bubbleBefore)
	bubbleAfter := bubbleSort(bubbleBefore)
	fmt.Println(bubbleAfter)
}
