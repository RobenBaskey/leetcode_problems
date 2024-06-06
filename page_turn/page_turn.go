package main

import "fmt"

func main() {
	result := pageCount(6, 5)
	fmt.Println(result)
}

func pageCount(n int32, p int32) int32 {

	if n < p {
		return 0
	}

	if (n / 2) >= p {
		fmt.Println("worj 1")
		return p / 2
	} else {

		fmt.Println("worj 2")
		if n%2 == 0 {
			reversTotal := (n - p) + 1
			return reversTotal / 2
		}
		reversTotal := (n - p)
		if reversTotal%2 == 0 {
			return reversTotal / 2
		}

		return (reversTotal - 1) / 2

	}

}
