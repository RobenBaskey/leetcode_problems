package main

import "fmt"

func main() {
	result := buildArray([]int{5, 0, 1, 2, 3, 4})
	fmt.Println(result)
}

func buildArray(nums []int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
