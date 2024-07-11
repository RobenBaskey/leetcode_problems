package main

import "fmt"

func main() {
	arr := []string{"--X", "X++", "X++"}
	result := finalValueAfterOperations(arr)
	fmt.Println(result)
}

func finalValueAfterOperations(operations []string) int {
	var x = 0
	for _, value := range operations {
		if value == "--X" || value == "X--" {
			x = x - 1
		} else {
			x = x + 1
		}
	}
	return x
}
