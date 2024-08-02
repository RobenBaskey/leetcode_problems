package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1, 1, 3}
	result := goodPairHashTable(arr)
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

func goodPairHashTable(nums []int) int {
	var result int

	m := make(map[int]int)
	for _, num := range nums {
		result += m[num]
		fmt.Println(result)
		m[num]++
	}

	return result
}
