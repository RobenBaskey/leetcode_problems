package main

import "fmt"

func main() {
	result := generate(1)
	fmt.Println(result)
}

func generate(numRows int) [][]int {

	if numRows == 0 {
		return [][]int{}
	}

	if numRows == 1 {
		return [][]int{{1}}
	}

	result := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		row := []int{}

		for j := 0; j < i+1; j++ {
			data := 0
			if j == 0 {
				data = 0 + result[i-1][j]

			} else {
				if j >= len(result[i-1]) {
					data = 0 + result[i-1][j-1]

				} else {
					data = result[i-1][j] + result[i-1][j-1]
				}

			}
			row = append(row, data)
		}

		result = append(result, row)
	}

	return result
}
