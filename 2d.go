package main

import "fmt"

func main() {
	// Define a 2D array
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Access individual elements
	fmt.Println("Accessing individual elements:")
	fmt.Println("arr[0][0] =", arr[0][0]) // Accessing the element at row 0, column 0
	fmt.Println("arr[1][2] =", arr[1][2]) // Accessing the element at row 1, column 2
	fmt.Println("arr[2][1] =", arr[2][1]) // Accessing the element at row 2, column 1

	// Accessing rows and columns
	fmt.Println("\nAccessing rows and columns:")
	fmt.Println("Row 1:", arr[1])              // Accessing the entire row at index 1
	//fmt.Println("Column 2:", getColumn(arr, 2)) // Accessing the entire column at index 2
}

// Function to get a column from a 2D array
func getColumn(arr [][]int, columnIndex int) []int {
	column := make([]int, len(arr))
	for i := range arr {
		column[i] = arr[i][columnIndex]
	}
	return column
}
