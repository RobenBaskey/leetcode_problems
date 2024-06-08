package main

import "fmt"

func main() {
	resultAr := getConcatenation([]int{1, 3, 2, 1})
	for _, data := range resultAr {
		fmt.Println(data)
	}

}

func getConcatenation(nums []int) []int {
	// var ans []int
	// numLen := len(nums)

	// for i := 0; i < numLen*2; i++ {
	// 	if i >= numLen {
	// 		ans = append(ans, nums[i-numLen])
	// 	} else {
	// 		ans = append(ans, nums[i])
	// 	}

	// }
	return append(nums, nums...)

}
