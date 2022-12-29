package main

func bubbleSort1(arr []int) []int {
	var leng int = len(arr)

	for i := 0; i < leng; i++ {
		for j := 0; j < (leng-1)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	var swapped, leng = true, len(arr)

	for swapped {
		swapped = false
		for i := 0; i < leng-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
