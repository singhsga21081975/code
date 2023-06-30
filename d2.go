package main

import "fmt"

func main() {
	// Create a 2D slice with 3 rows and 4 columns
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
	}

	// Assign values to the elements of the 2D slice
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[0][3] = 4
	matrix[1][0] = 5
	matrix[1][1] = 6
	matrix[1][2] = 7
	matrix[1][3] = 8
	matrix[2][0] = 9
	matrix[2][1] = 10
	matrix[2][2] = 11
	matrix[2][3] = 12

	// Display the elements of the 2D slice
	for _, row := range matrix {
		for _, element := range row {
			fmt.Print(element, " ")
		}
		fmt.Println()
	}

	row := 4
	col := 4
	m := make([][]int, row)
	for i := range m {
		m[i] = make([]int, col)


	}

}

