package main

import (
	"fmt"
)

func main() {
	accounts := []int{5, 1, 6}
	result := subsetXORSum(accounts)
	fmt.Println(result)
}

func subsetXORSum(nums []int) int {
	totalSum := 0
	for i := 0; i < 1<<len(nums); i++ {
		mySum := 0
		for j := 0; j < len(nums); j++ {
			if i&(1<<j) != 0 {
				mySum = nums[j] ^ mySum
			}
		}
		totalSum += mySum
	}

	return totalSum
}
