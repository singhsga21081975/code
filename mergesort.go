package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)

	i := 0
	j := 0
	for i < len(left) && j < len(right)  {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i+=1
		} else {
			result = append(result, right[j])
			j+=1
		}
	}

	// Append any remaining elements
	for i < len(left) {
		
		result = append(result, left[i])
		i+=1
	} 
	for j < len(right) {
		result = append(result, right[j])
		j+=1
	}

	return result
}

func main() {
	arr := []int{9, 2, 5, 1, 7, 6, 8, 3, 4}

	fmt.Println("Original array:", arr)

	sortedArr := mergeSort(arr)

	fmt.Println("Sorted array:", sortedArr)
}
