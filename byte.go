package main

import "fmt"

func main() {
	// Declare a slice of bytes
	bytesSlice := []byte{'H', 'e', 'l', 'l', 'o'}

	// Print the slice
	fmt.Println(bytesSlice)

	// Access individual elements
	fmt.Println(bytesSlice[0]) // Output: 72 (ASCII value of 'H')
	fmt.Println(bytesSlice[1]) // Output: 101 (ASCII value of 'e')

	// Modify a byte in the slice
	bytesSlice[0] = 'G'
	fmt.Println(bytesSlice) // Output: [71 101 108 108 111]

	// Iterate over the slice
	for _, b := range bytesSlice {
		fmt.Printf("%c ", b)
	}

	fmt.Println("str ",string(bytesSlice))
	// Output: G e l l o


	bytes := []byte{72, 101, 108, 108, 111}  // 'H', 'e', 'l', 'l', 'o'
fmt.Println(len(bytes))                  // Output: 5
fmt.Println(string(bytes))                // Output: "Hello"
bytes[0] = 74                            // Modify first element to 'J'
fmt.Println(string(bytes))                // Output: "Jello"

}
