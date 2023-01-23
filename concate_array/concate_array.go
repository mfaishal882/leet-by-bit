package main

import "fmt"

func getConcatenation(nums []int) []int {
	// var result []int = nums
	// for _, value := range nums {
	// 	result = append(result, value)
	// }
	// return result
	return append(nums, nums...)
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
}
