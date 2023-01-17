package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 3
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Println("Target found at index:", index)
	} else {
		fmt.Println("Target not found in array")
	}
}
