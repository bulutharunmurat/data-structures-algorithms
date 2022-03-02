package main

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
