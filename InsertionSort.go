package main

// BEST CASE n and WORST CASE = n(n+1)/2

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		tempIndex := i
		for tempIndex > 1 && arr[tempIndex-1] > arr[tempIndex] {

			smaller := arr[tempIndex]

			arr[tempIndex] = arr[tempIndex-1]
			arr[tempIndex-1] = smaller

			tempIndex = tempIndex - 1
		}
	}

	return arr

}
