package main

import (
	"fmt"
)

func main() {
	arr := []int{1}
	result := maxProfit(arr)
	fmt.Println(result)
}

func maxProfit(prices []int) int {
	min := prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > result {
			result = prices[i] - min
		}
	}

	//i=0 => d=7 > m=7 => min=7 => result= 7 - 7 = 0 min no cng
	//i=1 => d=1>m=7 => m=1 => r=d-m=-6 min cng , result err
	//i=2 => d=5>m=1 => m=1 => r=d-m=4 min no cng, r=4
	//i=3 => d=3>m=1 => m=1 => r=d-m=3-1=2 min no cng r=2

	// if data < min { // data = 5 , min = 1 == true
	// 	min = data // min = 1
	// } else {
	// 	if data-min > result { // data = 1 - 1 = 0 < 0
	// 		result = data - min
	// 	}
	// }

	return result
}
