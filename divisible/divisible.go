package main

import "fmt"

func main() {

	var n int32 = 6
	var k int32 = 3

	arr := []int32{1, 3, 2, 6, 1, 2}

	divisibleSumPairs(n, k, arr)

	checkIfElse("5")

}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	if n == 0 {
		return 0
	}

	var result int32 = 0
	for i := 1; i < int(n); i++ {
		sublist := ar[i:n]
		for _, data := range sublist {
			if (ar[i-1]+data)%k == 0 {
				result++
			}
		}
	}

	return result
}

func checkIfElse(data string) {
	if data != "" && data != "1" && data != "2" && data != "3" {
		fmt.Println("Invalid data")
	} else {
		fmt.Println("Valid")
	}
}
