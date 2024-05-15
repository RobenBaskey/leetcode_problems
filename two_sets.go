package main

import "fmt"

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func findMax(a []int32) int32 {
	maxNum := a[0]
	for _, j := range a {
		if j > maxNum {
			maxNum = j
		}
	}
	return maxNum
}

func findMin(b []int32) int32 {
	minNum := b[0]
	for _, j := range b {
		if j < minNum {
			minNum = j
		}
	}
	return minNum
}

func getTotalX(a []int32, b []int32) int32 {
	var result int32 = 0
	maxNum := findMax(a)
	fmt.Println("My max is ", maxNum)
	minNum := findMin(b)
	fmt.Println("My min is ", minNum)
	for i := maxNum; i < minNum+1; i++ {
		isFactorMultiple := true

		for _, data := range a {
			if i%data != 0 {
				isFactorMultiple = false
				break
			}
		}

		for _, data := range b {
			if data%i != 0 {
				isFactorMultiple = false
				break
			}
		}

		if isFactorMultiple == true {
			result++
		}
	}
	return result
}

func main() {
	a := []int32{2, 6}
	b := []int32{24, 36}
	result := getTotalX(a, b)
	fmt.Println("My Result is ", result)
}
