package main

import "fmt"

func main() {
	arr := []int32{1, 4, 4, 4, 5, 3}
	result := migratoryBirds(arr)
	fmt.Println("Result is ", result)
}

func migratoryBirds(arr []int32) int32 {

	var frequency int32 = 0
	var result int32 = 0

	for i := 1; i < 6; i++ {

		conteinResult := contains(arr, int32(i))
		fmt.Println(i, " Contains = ", conteinResult, " , Frequ = ", frequency)
		if conteinResult > 1 && conteinResult > frequency {
			frequency = conteinResult
			result = int32(i)
		}
	}
	return result
}

func contains(slice []int32, item int32) int32 {
	var result int32 = 0
	for _, element := range slice {
		if element == item {
			result++
		}
	}
	return result
}
