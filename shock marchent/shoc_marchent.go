package main

import (
	"fmt"
	"sort"
)

func main() {
	//data1 := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	data2 := []int32{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
	result := sockMerchant(15, data2)
	fmt.Println(result)
}

func sockMerchant(n int32, ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	var i, currentValueCount, pairCount int32
	currentValue := ar[0]

	for i = 0; i < n; i++ {
		if ar[i] == currentValue {
			currentValueCount++
		} else {
			currentValue = ar[i]
			pair := currentValueCount / 2
			pairCount = pairCount + pair
			currentValueCount = 1
		}

	}
	pair := currentValueCount / 2
	pairCount = pairCount + pair
	return pairCount
}
