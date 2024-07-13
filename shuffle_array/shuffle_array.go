package main

import "fmt"

func main() {
	arr := []int{2, 5, 1, 3, 4, 7}
	result := shuffle(arr, 2)
	fmt.Println(result)
}

func shuffle(nums []int, n int) []int {
	if n == 1 {
		return nums
	}
	y := nums[n:]
	result := []int{}
	i := 0
	for i = 0; i < (len(nums) - len(y)); i++ {
		result = append(result, nums[i])
		if len(y) > i {
			result = append(result, y[i])
		}
	}
	if i < len(y) {
		result = append(result, y[n:]...)
	}
	return result
}

func removeElement(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
