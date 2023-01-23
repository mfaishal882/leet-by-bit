package main

import (
	"fmt"
)

func finalValueAfterOperations(operations []string) int {
	var result int
	for _, value := range operations {
		if value == "++X" || value == "X++" {
			result = result + 1
		} else {
			result = result - 1
		}
	}
	return result
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
