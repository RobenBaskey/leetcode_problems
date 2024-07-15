package main

import "fmt"

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	result := maximumWealth(accounts)
	fmt.Println(result)
}

func maximumWealth(accounts [][]int) int {
	maximum := 0
	for i := range accounts {
		myCount := 0
		for j := range accounts[i] {
			myCount += accounts[i][j]
		}
		maximum = max(maximum, myCount)
	}
	return maximum
}
