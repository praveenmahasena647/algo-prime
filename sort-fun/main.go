package main

import "log"

func main() {
	var arr []int = []int{4, 2, 6, 8}

	var newArr []int = bubbleSort(arr)
	log.Println(newArr)
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
