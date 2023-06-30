package main

import "fmt"

func processMap(m map[string]int) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func main() {
	myMap := map[string]int {}
	
	
		myMap["apple"]=1

	
	processMap(myMap)
}
