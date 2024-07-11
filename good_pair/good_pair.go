package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 1}
	result := numIdenticalPairs2(arr)
	fmt.Println(result)
}

func numIdenticalPairs(nums []int) int {
	ans := 0
	for i, value := range nums {
		for _, data := range nums[i+1:] {
			if value == data {
				ans++
			}
		}
	}

	return ans
}

func numIdenticalPairs2(nums []int) int {

	myMap := make(map[int]int)
	for _, value := range nums {
		myMap[value]++
	}

	ans := 0
	for _, value := range myMap {
		ans = ans + (value * (value - 1) / 2)
	}
	return ans
}
