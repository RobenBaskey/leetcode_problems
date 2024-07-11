package main

import "fmt"

func main() {
	arr := []int{3, 6, 9}
	result := minimumOperations(arr)
	fmt.Println(result)
}

func minimumOperations(nums []int) int {
	result := 0
	for _, j := range nums {
		if (j+1)%3 == 0 {
			result++
		} else if (j-1)%3 == 0 {
			result++
		}
	}
	return result
}
